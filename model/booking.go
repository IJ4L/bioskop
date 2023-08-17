package model

type Booking struct {
	Id           int    `json:"id" gorm:"primaryKey;autoIncrement"`
	IdPemesan    int    `json:"id_pemesan" `
	IdFilm       int    `json:"id_film" `
	IdSeat       int    `json:"id_seat" gorm:"foreignKey:idSeat"`
	IdBloc       int    `json:"id_bloc" gorm:"foreignKey:idBloc"`
	BookingDate  string `json:"booking_date" gorm:"varchar(20);not null"`
	BookingSatus string `json:"booking_status" gorm:"varchar(20);not null"`
}
