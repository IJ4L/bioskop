package model

type Film struct {
	Id        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Judul     string  `json:"judul" gorm:"varchar(50);not null"`
	Desk      string  `json:"deskripsi" gorm:"text"`
	Genre     string  `json:"genre" gorm:"varchar(20);not null"`
	Imdb      string  `json:"imdb" gorm:"varchar(20);not null"`
	Poster    string  `json:"poster" gorm:"varchar(50);not null"`
	Durations string  `json:"durations" gorm:"varchar(20);not null"`
	Views     string  `json:"views" gorm:"varchar(20);not null"`
	Pg        string  `json:"p-g" gorm:"varchar(20);not null"`
	Price     string  `json:"price" gorm:"varchar(20);not null"`
	Actors    []Actor `json:"actors" gorm:"many2many:film_actors;"`
}
