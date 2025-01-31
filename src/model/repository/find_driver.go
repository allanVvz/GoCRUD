package repository

import (
	"context"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// findDriver busca um motorista no banco de dados com um filtro específico.
func (r *DriverRepository) findDriver(filter bson.M) (request.DriverRequest, *rest_err.RestErr) {
	collection := r.DB.Collection("drivers")

	// Estrutura para armazenar os dados do motorista
	var entity struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Status       bool               `bson:"status"`
		Name         string             `bson:"name"`
		Rg           string             `bson:"rg"`
		Registration string             `bson:"registration"`
		Salary       float64            `bson:"salary"`
	}

	// Executa a busca no MongoDB
	err := collection.FindOne(context.Background(), filter).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Info("Driver not found", zap.Any("filter", filter))
			return request.DriverRequest{}, rest_err.NewNotFoundError("Driver not found")
		}

		logger.Error("Error finding driver", err, zap.Any("filter", filter))
		return request.DriverRequest{}, rest_err.NewInternalServerError("Error finding driver")
	}

	// Preenche a struct request.DriverRequest com os dados encontrados
	driver := request.DriverRequest{
		Id:           entity.ID.Hex(),
		Status:       entity.Status,
		Name:         entity.Name,
		Rg:           entity.Rg,
		Registration: entity.Registration,
		Salary:       entity.Salary,
	}

	logger.Info("Driver found successfully", zap.Any("filter", filter))
	return driver, nil
}

// FindDriverByID busca um motorista pelo ID.
func (r *DriverRepository) FindDriverByID(id string) (request.DriverRequest, *rest_err.RestErr) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Invalid ID format", err, zap.String("driverId", id))
		return request.DriverRequest{}, rest_err.NewBadRequestError("Invalid ID format")
	}

	return r.findDriver(bson.M{"_id": objectID})
}

// FindDriverByReg busca um motorista pelo registro.
func (r *DriverRepository) FindDriverByReg(registration string) (request.DriverRequest, *rest_err.RestErr) {
	return r.findDriver(bson.M{"registration": registration})
}

// FindDriverByRg busca um motorista pelo RG.
func (r *DriverRepository) FindDriverByRg(rg string) (request.DriverRequest, *rest_err.RestErr) {
	return r.findDriver(bson.M{"rg": rg})
}
