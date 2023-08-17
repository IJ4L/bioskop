package model

type FilmActor struct {
	FilmID  uint `gorm:"primaryKey"`
	ActorID uint `gorm:"primaryKey"`
}
