package models

type Node struct {
	BaseModel
	Host     string `gorm:"size:255;unique;not null" json:"host"`
	Name     string `gorm:"size:255;not null" json:"name"`
	User     string `gorm:"size:255;not null" json:"user"`
	Password string `gorm:"size:255;not null" json:"password"`
}

func (Node) TableName() string {
	return "nodes"
}

type NodeAndPhone struct {
	Host   string `json:"host"`
	Serial string `json:"serial"`
}
