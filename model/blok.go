package model

type Bloc struct {
	IdBloc int `json:"id" gorm:"primaryKey;autoIncrement"`
	NameBloc   string `json:"name" gorm:"varchar(20);not null"`
}