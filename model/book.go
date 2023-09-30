package model

type Genre struct {
	Id    int    `json:"id" db:"id"`
	Genre string `json:"genre" db:"genre"`
}

type Books struct {
	Id          uint   `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Type        bool   `json:"type" db:"type"`
	Number      uint   `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	Language    string `json:"language" db:"language"`
	Year        string `json:"year" db:"year"`
	GenerId     int    `json:"generid" db:"generid"`
}

type UpdateBook struct {
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Type        bool   `json:"type" db:"type"`
	Number      uint   `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	Language    string `json:"language" db:"language"`
	Year        string `json:"year" db:"year"`
}

type SuccessData struct {
	Data interface{} `json:"data"`
}
type UpdateGenreInput struct {
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Type        bool   `json:"type" db:"type"`
	Number      uint   `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	Language    string `json:"language" db:"language"`
	Year        string `json:"year" db:"year"`
	Genre string `json:"genre" db:"genre"`
}

type UpdateGenreOutput struct {
	Name        string `json:"name" db:"name"`
	Author      string `json:"author" db:"author"`
	Type        bool   `json:"type" db:"type"`
	Number      uint   `json:"number" db:"number"`
	Description string `json:"description" db:"description"`
	Language    string `json:"language" db:"language"`
	GenreId     int    `json:"genreId"`
	Year        string `json:"year" db:"year"`
}
