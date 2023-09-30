package repository

import (
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app/model"
)

type GenreRepo struct {
	db *sqlx.DB
}

func NewGenreRepo(db *sqlx.DB) *GenreRepo {
	return &GenreRepo{db: db}
}

func (repo *GenreRepo) Create(genre model.Genre) (id int, err error) {
	db := repo.db
	row, err := db.Query(CreateGenreQuery, genre.Id, genre.Genre)
	if err != nil {
		return 0, err
	}

	for row.Next() {
		err = row.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}

func (repo * GenreRepo) GetGenre() (genrelist [] model.Genre, err error)  {
	db := repo.db
	err = db.Select(&genrelist, GetGenreQuery)
	if err != nil {
		return genrelist ,err
	}
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return genrelist, err
		}
		return genrelist, err
	}
	return genrelist, err
}

func (repo *GenreRepo) GetGenreById(genreID int) (genre model.Genre, err error) {
	db := repo.db
	err = db.Get(&genre, GetByIdGenreQuery, genreID)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return genre, err
		}
		return genre, err
	}
	return genre, err
}
func (repo *GenreRepo) DeleteGenre(genreID int) error {
	db := repo.db
	row, err := db.Exec(DeleteGenreQuery, genreID)
	if err != nil {
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("not working")
	}
	return nil
}
func (repo *GenreRepo) UpdateGenre(genreID int, input model.UpdateGenreInput) error {
	db := repo.db
	row, err := db.Exec(UpdateGenreQuery, input.Genre, genreID)
	if err != nil {
		return err
	}
	rowAffected, err := row.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected == 0 {
		return errors.New("not working")
	}
	return nil
}
