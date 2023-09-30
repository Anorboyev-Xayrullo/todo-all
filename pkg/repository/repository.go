package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/model"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}
type TodoList interface {
	Create(userID int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listID int) (todo.TodoList, error)
	Delete(userId, listID int) error
	Update(userId, listID int, input todo.UpdateListInput) error
}
type TodoItem interface {
	Create(listID int, item todo.TodoItem) (int, error)
	GetAll(userId, listID int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Book interface {
	CreateBook(book model.Books) error
	GetBookList() (bookList []model.Books, err error)
	GetBookById(bookId int) (book model.Books, err error)
	GetBooksByGenre(genre string) (book model.UpdateGenreOutput, err error)
	UpdateBook(bookId int, input model.UpdateBook) error
	DeleteBook(bookId int) error
}

type Genre interface {
	Create(genre model.Genre) (id int, err error)
	GetGenreById(genreID int) (genre model.Genre, err error)
	GetGenre() (genrelist [] model.Genre, err error)
	DeleteGenre(genreID int) error
	UpdateGenre(genreID int, input model.UpdateGenreInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
	Book
	Genre
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
		Book:          NewBookRepo(db),
		Genre:         NewGenreRepo(db),
	}
}
