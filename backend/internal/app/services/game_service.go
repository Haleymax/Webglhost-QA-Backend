package services

import (
	"errors"
	"github.com/Webglhost-QA-Backend/backend/config"
	"github.com/Webglhost-QA-Backend/backend/internal/app/models"
	"github.com/Webglhost-QA-Backend/backend/internal/app/repositories"
	"github.com/Webglhost-QA-Backend/backend/pkg/feishu"
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
	UpdateMongoByFeishu(messageChan *chan string, config config.FeishuConfig)
	FindByQuery(query models.GameRequest) (models.Game, error)
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

func (gs *GameServiceImpl) FindByQuery(query models.GameRequest) (models.Game, error) {
	filter := bson.M{
		"game_name": query.GameName,
		"case_type": query.CaseType,
		"game_type": query.GameType,
	}
	result, err := gs.gameRepo.FindAll(filter)
	if err != nil {
		return models.Game{}, err
	}
	log.Println(result)
	if len(result) == 1 {
		return result[0], nil
	} else {
		return models.Game{}, errors.New("game not exist")
	}
	return models.Game{}, errors.New("Game data information is err")
}

func (gs *GameServiceImpl) UpdateMongoByFeishu(messageChan *chan string, config config.FeishuConfig) {
	start_msg := "start run script"
	var error_msg string
	var process_msg string
	*messageChan <- start_msg
	feishu_client, err := feishu.NewFeishuClient(config)
	if err != nil {
		error_msg = err.Error()
		*messageChan <- error_msg
		return
	}
	err = feishu_client.GetUserToken()
	if err != nil {
		error_msg = err.Error()
		*messageChan <- error_msg
		return
	}
	err = feishu_client.GetFeishuSheetData()
	if err != nil {
		error_msg = err.Error()
		*messageChan <- error_msg
		return
	}
	for _, game := range feishu_client.GameList {
		QuerGame := models.GameRequest{
			GameName: game.GameName,
			GameType: game.GameType,
			CaseType: game.CaseType[0],
		}
		exist, _ := gs.FindByQuery(QuerGame)
		if reflect.DeepEqual(exist, models.Game{}) {
			exist.GameUrl = game.GameUrl
			NewType := make([]string, 2)
			NewType = append(NewType, "daily")
			NewType = append(NewType, game.CaseType[0])
			exist.CaseType = NewType
			if err := gs.UpdateGame(exist); err != nil {
				process_msg = "game " + game.GameName + "exist, update game faild" + err.Error()
			} else {
				process_msg = "game " + game.GameName + "exist, update game success"
			}
			*messageChan <- process_msg
		} else {
			game.Status = true
			err := gs.AddGame(game)
			if err != nil {
				process_msg = "game " + game.GameName + "not exist, insert game faild" + err.Error()
			} else {
				process_msg = "game " + game.GameName + " not exist, insert game success"
			}
			*messageChan <- process_msg
		}

	}
	end_msg := "end run script"
	*messageChan <- end_msg
}
