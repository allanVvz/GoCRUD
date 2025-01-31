package repository

import (
	"context"

	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *DriverRepository) FireDriver(driverID string, newStatus bool) *rest_err.RestErr {
	collection := r.DB.Collection("drivers")

	// Convert driverID to ObjectID
	objectID, err := primitive.ObjectIDFromHex(driverID)
	if err != nil {
		return rest_err.NewBadRequestError("Invalid driver ID format")
	}

	// Define the filter and update
	filter := bson.M{"_id": objectID}
	update := bson.M{"$set": bson.M{"status": newStatus}}

	// Execute the update
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return rest_err.NewInternalServerError("Failed to update driver status")
	}

	return nil
}
