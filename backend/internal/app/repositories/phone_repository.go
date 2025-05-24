package repositories

import (
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type PhoneRepository interface {
	InsertOne(phone *models.Phone) error
}

type PhoneRepositoryImpl struct {
	mongo *mongo.Client
}

func NewPhoneRepository(mongo *mongo.Client) *PhoneRepositoryImpl {
	return &PhoneRepositoryImpl{
		mongo: mongo,
	}
}
