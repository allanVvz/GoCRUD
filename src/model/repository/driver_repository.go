package repository

import (
	"context"

	"github.com/allanVvz/GoCRUD/src/controller/model/request"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DriverRepository struct {
	DB *mongo.Database
}

// Novo repositório
func NewDriverRepository(db *mongo.Database) *DriverRepository {
	return &DriverRepository{DB: db}
}

// Salva um motorista diretamente no banco usando o request
func (r *DriverRepository) SaveDriver(driver request.DriverRequest) (string, error) {
	collection := r.DB.Collection("drivers")

	// Gera ou converte o ObjectID
	objectID := primitive.NewObjectID()
	if driver.Id != "" {
		parsedID, err := primitive.ObjectIDFromHex(driver.Id)
		if err != nil {
			return "", err // Retorna erro se o ID for inválido
		}
		objectID = parsedID
	}

	// Documento para salvar no banco
	entity := struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Status       bool               `bson:"status"`
		Name         string             `bson:"name"`
		Rg           string             `bson:"rg"`
		Registration string             `bson:"registration"`
		Salary       float64            `bson:"salary"`
	}{
		ID:           objectID,
		Status:       driver.Status,
		Name:         driver.Name,
		Rg:           driver.Rg,
		Registration: driver.Registration,
		Salary:       driver.Salary,
	}

	// Insere no banco
	_, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		return "", err
	}

	// Retorna o ID gerado
	return objectID.Hex(), nil
}
