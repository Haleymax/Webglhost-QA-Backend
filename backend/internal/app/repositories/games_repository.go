package repositories

import (
	"context"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var game_collection = "games"

type GamesRepository interface {
	GetCollection() *mongo.Collection
	Insert(game models.Game) error
	Delete(game models.Game) error
	Update(game models.Game) error
	FindAll(filter bson.M) ([]models.Game, error)
	FindById(id string) (models.Game, error)
	FindByName(name string) (models.Game, error)
	FindAllWXGames() ([]models.Game, error)
	FindAllRPK() ([]models.Game, error)
}

type GamesRepositoryImpl struct {
	mongo *mongo.Client
}

func NewGamesRepository(mongo *mongo.Client) *GamesRepositoryImpl {
	return &GamesRepositoryImpl{
		mongo: mongo,
	}
}

func (r *GamesRepositoryImpl) GetCollection() *mongo.Collection {
	return r.mongo.Database(db_name).Collection(game_collection)
}

func (r *GamesRepositoryImpl) Insert(game models.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := r.GetCollection().InsertOne(ctx, game)
	if err != nil {
		return err
	}
	return nil
}

func (r *GamesRepositoryImpl) Delete(game models.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(game.ID)
	filter := bson.M{"_id": objID}
	_, err := r.GetCollection().DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (r *GamesRepositoryImpl) Update(game models.Game) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objID, _ := primitive.ObjectIDFromHex(game.ID)
	filter := bson.M{"_id": objID}

	newData := map[string]interface{}{
		"game_id":     game.GameId,
		"package":     game.Package,
		"type":        game.Type,
		"case_type":   game.CaseType,
		"game_engine": game.GameEngine,
		"game_url":    game.GameUrl,
		"game_name":   game.GameName,
		"game_type":   game.GameType,
		"status":      game.Status,
	}
	update := bson.M{"$set": newData}
	_, err := r.GetCollection().UpdateOne(ctx, filter, update)
	return err
}

func (r *GamesRepositoryImpl) FindAll(filter bson.M) ([]models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var games []models.Game
	cursor, err := r.GetCollection().Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var game models.Game
		if err := cursor.Decode(&game); err != nil {
			return nil, err
		}
		games = append(games, game)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return games, nil
}

func (r *GamesRepositoryImpl) FindById(id string) (models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var game models.Game
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objID}
	cursor, err := r.GetCollection().Find(ctx, filter)
	if err != nil {
		return models.Game{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		if err := cursor.Decode(&game); err != nil {
			return models.Game{}, err
		}
	}
	if err := cursor.Err(); err != nil {
		return models.Game{}, err
	}
	return game, nil
}

func (r *GamesRepositoryImpl) FindByName(name string) (models.Game, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var game models.Game
	filter := bson.M{"name": game}
	cursor, err := r.GetCollection().Find(ctx, filter)
	if err != nil {
		return models.Game{}, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		if err := cursor.Decode(&game); err != nil {
			return models.Game{}, err
		}
	}
	if err := cursor.Err(); err != nil {
		return models.Game{}, err
	}
	return game, nil
}
