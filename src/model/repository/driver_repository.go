package repository

import (
	"context"
	"errors"

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

	// Converte o ID para ObjectID se for necessário
	var objectID primitive.ObjectID
	if driver.Id != "" {
		var err error
		objectID, err = primitive.ObjectIDFromHex(driver.Id)
		if err != nil {
			return "", errors.New("invalid ID format")
		}
	} else {
		objectID = primitive.NewObjectID()
	}

	// Documento para salvar no banco
	entity := struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Status       int8               `bson:"status"`
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

// Busca um motorista pelo ID
func (r *DriverRepository) FindDriverByID(id string) (request.DriverRequest, error) {
	collection := r.DB.Collection("drivers")

	// Converte o ID para ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return request.DriverRequest{}, errors.New("invalid ID format")
	}

	// Busca o documento no banco
	var entity struct {
		ID           primitive.ObjectID `bson:"_id,omitempty"`
		Status       int8               `bson:"status"`
		Name         string             `bson:"name"`
		Rg           string             `bson:"rg"`
		Registration string             `bson:"registration"`
		Salary       float64            `bson:"salary"`
	}

	err = collection.FindOne(context.Background(), struct {
		ID primitive.ObjectID `bson:"_id"`
	}{ID: objectID}).Decode(&entity)
	if err != nil {
		return request.DriverRequest{}, err
	}

	// Retorna o request preenchido
	return request.DriverRequest{
		Id:           entity.ID.Hex(),
		Status:       entity.Status,
		Name:         entity.Name,
		Rg:           entity.Rg,
		Registration: entity.Registration,
		Salary:       entity.Salary,
	}, nil
}
