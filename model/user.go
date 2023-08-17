package model

type User struct {
	Id           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name" gorm:"varchar(20);not null"`
	Email        string `json:"email" gorm:"varchar(30);not null"`
	Password     string `json:"password" gorm:"varchar(20);not null"`
	PhotoProfile string `json:"photo_profile" gorm:"varchar(255);not null"`
	CreatedAt    string `json:"created" gorm:"varchar(20);not null"`
}
