package repository

import (
	"context"
	"os"

	"github.com/allanVvz/GoCRUD/src/config/logger"
	"github.com/allanVvz/GoCRUD/src/config/rest_err"
	"github.com/allanVvz/GoCRUD/src/model"
	"github.com/allanVvz/GoCRUD/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

// DriverRepository representa o repositório para o recurso Driver.
type DriverRepository struct {
	databaseConnection *mongo.Database
}

// NewDriverRepository cria uma nova instância de DriverRepository.
func NewDriverRepository(dbConnection *mongo.Database) *DriverRepository {
	return &DriverRepository{
		databaseConnection: dbConnection,
	}
}

// CreateDriver insere um novo motorista no banco de dados.
func (dr *DriverRepository) CreateDriver(
	driverDomain model.Driver,
) (model.Driver, *rest_err.RestErr) {
	logger.Info("Init CreateDriver repository",
		zap.String("journey", "createDriver"))

	// Nome da coleção
	collectionName := os.Getenv("MONGODB_DRIVER_DB")
	if collectionName == "" {
		collectionName = "drivers" // Valor padrão caso a variável de ambiente não esteja configurada
	}

	collection := dr.databaseConnection.Collection(collectionName)

	// Converte o domínio para a entidade MongoDB
	entity := converter.ConvertDomainToEntity(driverDomain)

	// Insere o documento na coleção
	result, err := collection.InsertOne(context.Background(), entity)
	if err != nil {
		logger.Error("Error trying to create driver", err,
			zap.String("journey", "createDriver"))
		return model.Driver{}, rest_err.NewInternalServerError(err.Error())
	}

	// Atualiza o ID da entidade com o valor gerado pelo MongoDB
	entity.ID = result.InsertedID.(primitive.ObjectID)

	logger.Info("CreateDriver repository executed successfully",
		zap.String("driverId", entity.ID.Hex()),
		zap.String("journey", "createDriver"))

	// Converte a entidade de volta para o domínio e retorna
	return converter.ConvertEntityToDomain(*entity), nil
}
