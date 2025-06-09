package services

import (
	"errors"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"reflect"
)

type GameService interface {
	AddGame(game models.Game) error
	DeleteGame(game models.Game) error
	UpdateGame(game models.Game) error
	FindAllGame() ([]models.Game, error)
	FindGameById(gameId string) (models.Game, error)
	FindGameByName(name string) (models.Game, error)
}

type GameServiceImpl struct {
	gameRepo repositories.GamesRepository
}

func NewGameService(gameRepo repositories.GamesRepository) *GameServiceImpl {
	return &GameServiceImpl{gameRepo: gameRepo}
}

func (gs *GameServiceImpl) AddGame(game models.Game) error {
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
	return gs.gameRepo.Update(game)
}

func (gs *GameServiceImpl) FindAllGame() ([]models.Game, error) {
	return gs.gameRepo.FindAll()
}

func (gs *GameServiceImpl) FindGameById(gameId string) (models.Game, error) {
	return gs.gameRepo.FindById(gameId)
}

func (gs *GameServiceImpl) FindGameByName(name string) (models.Game, error) {
	return gs.FindGameByName(name)
}
