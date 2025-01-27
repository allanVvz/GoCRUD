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

func (r *DriverRepository) FindDriverByID(id string) (request.DriverRequest, *rest_err.RestErr) {
	collection := r.DB.Collection("drivers")

	// Converte o ID para ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		logger.Error("Invalid ID format in FindDriverByID", err, zap.String("driverId", id))
		return request.DriverRequest{}, rest_err.NewBadRequestError("Invalid ID format")
	}

	// Define o filtro para buscar o documento pelo _id
	filter := bson.M{"_id": objectID}

	// Define uma estrutura para armazenar todos os dados do motorista
	var entity struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Status       bool               `bson:"status"`
		Name         string             `bson:"name"`
		Rg           string             `bson:"rg"`
		Registration string             `bson:"registration"`
		Salary       float64            `bson:"salary"`
	}

	// Executa a busca no MongoDB
	err = collection.FindOne(context.Background(), filter).Decode(&entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			logger.Info("Driver not found in FindDriverByID", zap.String("driverId", id))
			return request.DriverRequest{}, rest_err.NewNotFoundError("Driver not found")
		}

		logger.Error("Error finding driver in FindDriverByID", err, zap.String("driverId", id))
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

	logger.Info("Driver found successfully in FindDriverByID", zap.String("driverId", driver.Id))
	return driver, nil
}
