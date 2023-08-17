package model

type Seat struct {
	IdSeat     int `json:"id" gorm:"primaryKey;autoIncrement"`
	IdBloc     int `json:"idBloc" gorm:"foreignKey:idBloc"`
	SeatNumber int `json:"seat_number" gorm:"not null"`
}
