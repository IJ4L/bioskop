package model

type Actor struct {
	Id        int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string `json:"name" gorm:"varchar(20);not null"`
	Profesion string `json:"profesion" gorm:"varchar(20);not null"`
}
