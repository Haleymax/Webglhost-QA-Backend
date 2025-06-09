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
	return gs.AddGame(game)
}

func (gs *GameServiceImpl) DeleteGame(game models.Game) error {
	exists, _ := gs.FindGameById(game.ID)
	if reflect.DeepEqual(exists, models.Game{}) {
		return errors.New("game not exist")
	}
	return gs.DeleteGame(game)
}

func (gs *GameServiceImpl) UpdateGame(game models.Game) error {
	return gs.UpdateGame(game)
}

func (gs *GameServiceImpl) FindAllGame() ([]models.Game, error) {
	return gs.FindAllGame()
}

func (gs *GameServiceImpl) FindGameById(gameId string) (models.Game, error) {
	return gs.FindGameById(gameId)
}

func (gs *GameServiceImpl) FindGameByName(name string) (models.Game, error) {
	return gs.FindGameByName(name)
}
