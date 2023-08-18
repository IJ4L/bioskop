package request

type AddFilm struct {
	Judul     string  `json:"judul"`
	Desk      string  `json:"deskripsi"`
	Genre     string  `json:"genre"`
	Imdb      string  `json:"imdb"`
	Poster    string  `json:"poster"`
	Durations string  `json:"durations"`
	Views     string  `json:"views"`
	Pg        string  `json:"p-g"`
	Price     string  `json:"price"`
	ShowDate  string  `json:"show_date"`
	ShowTimes string  `json:"show_times"`
}

type UpdateFilm struct {
	Judul     string  `json:"judul"`
	Desk      string  `json:"deskripsi"`
	Genre     string  `json:"genre"`
	Imdb      string  `json:"imdb"`
	Poster    string  `json:"poster"`
	Durations string  `json:"durations"`
	Views     string  `json:"views"`
	Pg        string  `json:"p-g"`
	Price     string  `json:"price"`
}