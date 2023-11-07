package request

type AddFilm struct {
	Judul     string `json:"judul"`
	Desk      string `json:"deskripsi"`
	Genre     string `json:"genre"`
	Imdb      string `json:"imdb"`
	Poster    string `json:"poster"`
	Durations string `json:"durations"`
	Views     string `json:"views"`
	Pg        string `json:"p-g"`
	Price     string `json:"price"`
	ShowDate  string `json:"show_date"`
	ShowTimes string `json:"show_times"`
}

type BookingFilm struct {
	IdPemesan uint `json:"id_pemesan"`
	IdFilm    uint `json:"id_film"`
	IdSeat    uint `json:"id_seat"`
	IdBloc    uint `json:"id_bloc"`
}

type AddActor struct {
	Name      string `json:"name"`
	Profesion string `json:"profesion"`
	Photo     string `json:"photo"`
}

type ConnectActor struct {
	FilmID  uint `json:"film_id"`
	ActorID uint `json:"actor_id"`
}
