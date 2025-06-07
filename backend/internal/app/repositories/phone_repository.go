package repositories

import (
	"context"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var device_collection = "devices"

type PhoneRepository interface {
	FindAll() ([]models.Phone, error)
	FindOne(serial string) (models.Phone, error)
	InsertOne(phone *models.Phone) error
	Update(phone *models.Phone) error
	Delete(phone models.Phone) error
}

type PhoneRepositoryImpl struct {
	mongo *mongo.Client
}

func NewPhoneRepository(mongo *mongo.Client) *PhoneRepositoryImpl {
	return &PhoneRepositoryImpl{
		mongo: mongo,
	}
}

func (r *PhoneRepositoryImpl) GetCollection() *mongo.Collection {
	return r.mongo.Database(db_name).Collection(device_collection)
}

func (r *PhoneRepositoryImpl) InsertOne(phone *models.Phone) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.GetCollection().InsertOne(ctx, phone)
	return err
}

func (r *PhoneRepositoryImpl) Update(phone *models.Phone) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(phone.ID)
	filter := bson.M{"_id": objID}

	newData := map[string]string{
		"androidVersion":   phone.AndroidVersion,
		"marketName":       phone.MarketName,
		"marketNameSymbol": phone.MarketNameSymbol,
		"updateTime":       phone.UpdateTime,
	}
	updata := bson.M{"$set": newData}
	_, err := r.GetCollection().UpdateOne(ctx, filter, updata)
	return err
}

func (r *PhoneRepositoryImpl) FindAll() ([]models.Phone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var phones []models.Phone
	cursor, err := r.GetCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var phone models.Phone
		if err := cursor.Decode(&phone); err != nil {
			return nil, err
		}
		phones = append(phones, phone)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return phones, nil
}

func (r *PhoneRepositoryImpl) FindOne(serial string) (models.Phone, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var phone models.Phone
	filter := bson.M{"serial": serial}
	err := r.GetCollection().FindOne(ctx, filter).Decode(&phone)
	if err != nil {
		return models.Phone{}, err
	}
	return phone, nil
}

func (r *PhoneRepositoryImpl) Delete(phone models.Phone) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(phone.ID)
	filter := bson.M{"_id": objID}
	_, err := r.GetCollection().DeleteOne(ctx, filter)
	return err
}
