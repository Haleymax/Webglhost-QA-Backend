package repositories

import (
	"context"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var collection = "watcher"
var db_name = "webglhost"

type WatcherRepository interface {
	GetCollection() *mongo.Collection
	Insert(watcher *models.Watcher) error
	Delete(watcher *models.Watcher) error
	Update(watcher *models.Watcher) error
	FindAll() ([]*models.Watcher, error)
	FindOne(filter interface{}) (*models.Watcher, error)
}

type WatcherRepositoryImpl struct {
	mongo *mongo.Client
}

func NewWatcherRepository(mongo *mongo.Client) *WatcherRepositoryImpl {
	return &WatcherRepositoryImpl{
		mongo: mongo,
	}
}

func (r *WatcherRepositoryImpl) GetCollection() *mongo.Collection {
	return r.mongo.Database(db_name).Collection(collection)
}

func (r *WatcherRepositoryImpl) Insert(watcher *models.Watcher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.GetCollection().InsertOne(ctx, watcher)
	return err
}

func (r *WatcherRepositoryImpl) Delete(watcher *models.Watcher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": watcher.ID}

	_, err := r.GetCollection().DeleteOne(ctx, filter)
	return err
}

func (r *WatcherRepositoryImpl) Update(watcher *models.Watcher) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": watcher.ID}
	update := bson.M{"$set": watcher}
	_, err := r.GetCollection().UpdateOne(ctx, filter, update)
	return err
}

func (r *WatcherRepositoryImpl) FindAll() ([]*models.Watcher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var watchers []*models.Watcher
	cursor, err := r.GetCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var watcher models.Watcher
		if err := cursor.Decode(&watcher); err != nil {
			return nil, err
		}
		watchers = append(watchers, &watcher)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return watchers, nil
}

func (r *WatcherRepositoryImpl) FindOne(filter interface{}) (*models.Watcher, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var watcher models.Watcher
	cursor, err := r.GetCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		if err := cursor.Decode(&watcher); err != nil {
			return nil, err
		}
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return &watcher, nil
}
