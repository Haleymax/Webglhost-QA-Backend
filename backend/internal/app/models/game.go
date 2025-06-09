package models

type Game struct {
	ID         string   `json:"id" bson:"_id,omitempty"`
	Package    string   `json:"package" bson:"package"`
	Type       string   `json:"type" bson:"type"`
	CaseType   []string `json:"case_type" bson:"case_type"`
	GameEngine string   `json:"game_engine" bson:"game_engine"`
	GameUrl    string   `json:"game_url" bson:"game_url"`
	GameName   string   `json:"game_name" bson:"game_name"`
	GameType   string   `json:"game_type" bson:"game_type"`
	GameId     int      `json:"game_id" bson:"game_id"`
	Status     bool     `json:"status" bson:"status"`
}

type GameRequest struct {
	GameType string `json:"game_type" bson:"game_type"`
	CaseType string `json:"case_type" bson:"case_type"`
	GameName string `json:"game_name" bson:"game_name"`
}
