package model

type College struct {
	Nim    string `gorm:"primary_key" json:"nim"` 
	Name   string `json:"name"`
	Campus string `json:"campus"`
}

func (u *College) TableName() string {
	return "college"
}