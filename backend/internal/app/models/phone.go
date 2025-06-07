package models

import "time"

type Phone struct {
	ID               string    `json:"id" bson:"_id,omitempty"`
	Serial           string    `json:"serial" bson:"serial"`
	Manufacturer     string    `json:"manufacturer" bson:"manufacturer"`
	Model            string    `json:"model" bson:"model"`
	AndroidVersion   string    `json:"androidVersion" bson:"androidVersion"`
	Cpuabi           string    `json:"cpuabi" bson:"cpuabi"`
	MarketName       string    `json:"marketName" bson:"marketName"`
	MarketNameSymbol string    `json:"marketNameSymbol" bson:"marketNameSymbol"`
	CreateTime       time.Time `json:"createTime" bson:"createTime"`
	UpdateTime       time.Time `json:"updateTime" bson:"updateTime"`
}
