package model

type FilmActor struct {
	FilmID  uint `json:"film_id" gorm:"primaryKey"`
	ActorID uint `json:"actor_id" gorm:"primaryKey"`
}
