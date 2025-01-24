package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository struct {
	DB *mongo.Database
}

func (r *DriverRepository) UpdateDriverStatus(driverID string, newStatus int8) error {
	collection := r.DB.Collection("drivers")

	// Cria o filtro para buscar o motorista pelo ID
	filter := bson.M{"_id": driverID}

	// Define o update para alterar o status
	update := bson.M{"$set": bson.M{"status": newStatus}}

	// Atualiza o motorista no banco
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return nil
}
