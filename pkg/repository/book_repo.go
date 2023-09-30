package repository

import (
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/zhashkevych/todo-app/model"
)

type BookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (repo *BookRepo) CreateBook(book model.Books) error {

	db := repo.db
	row, err := db.Exec(CreateBookQuery, book.Name, book.Author , book.Type , book.Number , book.Description , book.Language, book.Year)
	if err != nil {
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		return errors.New("not update")
	}
	return nil
}

func (repo * BookRepo) GetBookList() (bookList []model.Books, err error)  {
	db := repo.db
	err = db.Select(&bookList, GetBookListQuery)
	if err != nil {
		return bookList ,err
	}
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return bookList, err
		}
		return bookList, err
	}
	return bookList, err
}

func (repo *BookRepo) GetBookById(bookId int) (book model.Books, err error){
	db := repo.db
	err = db.Get(&book, GetByIdGenreQuery, bookId)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return book, err
		}
		return book, err
	}
	return book, err
}
func (repo *BookRepo) GetBooksByGenre(genre string) (book model.UpdateGenreOutput, err error) {
	db := repo.db
	err = db.Get(&book, GetBooksByGenreQuery, genre)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return book, err
		}
		return book, err
	}
	return book, err
}

func (repo *BookRepo) UpdateBook(bookId int, input model.UpdateBook) error {
	db := repo.db
	row, err := db.Exec(UpdateBookQuery, input.Name, input.Number, input.Author, input.Year, input.Type, input.Language, input.Description, bookId)
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

func (repo *BookRepo) DeleteBook(bookId int) error {
	db := repo.db
	row, err := db.Exec(DeleteBookQuery, bookId)
	if err != nil {
		return err
	}

	rowAffected, err := row.RowsAffected()
	if err != nil {
		logrus.Error(err)
		return err
	}
	if rowAffected == 0 {
		return errors.New("not update")
	}
	return nil
}
