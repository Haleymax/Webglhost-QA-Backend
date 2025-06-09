package services

import (
	"errors"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"reflect"
)

type GameService interface {
	AddGame(game models.Game) error
	DeleteGame(game models.Game) error
	UpdateGame(game models.Game) error
	FindAllGame() ([]models.Game, error)
	FindGameById(gameId string) (models.Game, error)
	FindGameByName(name string) (models.Game, error)
	FindAllWxGame() ([]models.Game, error)
	FindAllRPK() ([]models.Game, error)
	FindAllByType(request models.GameRequest) ([]models.Game, error)
}

type GameServiceImpl struct {
	gameRepo repositories.GamesRepository
}

func NewGameService(gameRepo repositories.GamesRepository) *GameServiceImpl {
	return &GameServiceImpl{gameRepo: gameRepo}
}

func (gs *GameServiceImpl) AddGame(game models.Game) error {
	filter := bson.M{
		"game_name": game.GameName,
		"game_type": game.GameType,
		"case_type": game.CaseType,
	}
	exists, _ := gs.gameRepo.FindByName(filter)
	log.Println(reflect.DeepEqual(exists, models.Game{}))
	if !reflect.DeepEqual(exists, models.Game{}) {
		return errors.New("game exist")
	}
	return gs.gameRepo.Insert(game)
}

func (gs *GameServiceImpl) DeleteGame(game models.Game) error {
	exists, _ := gs.FindGameById(game.ID)
	if reflect.DeepEqual(exists, models.Game{}) {
		return errors.New("game not exist")
	}
	return gs.gameRepo.Delete(game)
}

func (gs *GameServiceImpl) UpdateGame(game models.Game) error {
	exists, _ := gs.FindGameById(game.ID)
	log.Println(exists)
	log.Println(reflect.DeepEqual(exists, models.Game{}))
	if reflect.DeepEqual(exists, models.Game{}) {
		return errors.New("game not exist")
	}
	return gs.gameRepo.Update(game)
}

func (gs *GameServiceImpl) FindAllGame() ([]models.Game, error) {
	filter := bson.M{}
	return gs.gameRepo.FindAll(filter)
}

func (gs *GameServiceImpl) FindAllWxGame() ([]models.Game, error) {
	filter := bson.M{"game_type": "weixin_minigame"}
	return gs.gameRepo.FindAll(filter)
}

func (gs *GameServiceImpl) FindAllRPK() ([]models.Game, error) {
	filter := bson.M{"game_type": "quick_minigame"}
	return gs.gameRepo.FindAll(filter)
}

func (gs *GameServiceImpl) FindAllByType(request models.GameRequest) ([]models.Game, error) {
	filter := bson.M{}
	if request.GameName == "" {
		filter = bson.M{"game_type": request.GameType, "case_type": request.CaseType}
	} else {
		filter = bson.M{"game_type": request.GameType, "case_type": request.CaseType, "game_name": request.GameName}
	}
	return gs.gameRepo.FindAll(filter)
}

func (gs *GameServiceImpl) FindGameById(gameId string) (models.Game, error) {
	return gs.gameRepo.FindById(gameId)
}

func (gs *GameServiceImpl) FindGameByName(name string) (models.Game, error) {
	return gs.FindGameByName(name)
}
