package repository

const (
	CreateBookQuery      = `insert into Book (name, author, type, number, description, language, year) values ($1, $2, $3, $4 , $5 , $6, $7 )`
	UpdateBookQuery      = `update Book set name = $1 , number = $2 , author = $3 , year = $4 , type = $5 , language = $6 , description = $7  where id = $8`
	DeleteBookQuery      = `delete from Book where id = $1`
	GetBooksByGenreQuery = `select name, author, year, language, type, description, genres.genre, number from Book inner join genres on Book.genre_id = genres.id where genres.genre = $1`
	GetBookListQuery     = `select * from Book`
	CreateGenreQuery     = `insert into genres(id, genre) values ($1, $2)`
	GetGenreQuery        = `select * from genres`
	GetAllGenresQuery    = `select * from genres limit $1 offset $2`
	GetByIdGenreQuery    = `select * from genres where id = $1`
	DeleteGenreQuery     = `delete from genres where id = $1`
	UpdateGenreQuery     = `update genres set genre = $1 where id = $2`
)
