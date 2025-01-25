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

	// Define uma estrutura para armazenar os dados do documento
	var entity struct {
		ID     primitive.ObjectID `bson:"_id,omitempty"`
		Status bool               `bson:"status"`
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

	// Converte o documento encontrado em um objeto request.DriverRequest
	driver := request.DriverRequest{
		Id:     entity.ID.Hex(),
		Status: entity.Status,
	}

	logger.Info("Driver found successfully in FindDriverByID", zap.String("driverId", driver.Id))
	return driver, nil
}
