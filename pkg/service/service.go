package service

import (
	"github.com/zhashkevych/todo-app"
	"github.com/zhashkevych/todo-app/model"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listID int) (todo.TodoList, error)
	Delete(userId, listID int) error
	Update(userId, listID int, input todo.UpdateListInput) error
}
type TodoItem interface {
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todo.UpdateItemInput) error
}

type Book interface {
	CreateBook(book model.Books) error
	GetBookList() (bookList []model.Books, err error)
	GetBookById(bookId int) (book model.Books, err error)
	UpdateBook(bookId int, input model.UpdateBook) error
	DeleteBook(bookId int) error
}
type Genre interface {
	Create(genre model.Genre) (id int, err error)
	GetGenre() (genrelist [] model.Genre, err error)
	GetGenreById(genreID int) (genre model.Genre, err error)
	DeleteGenre(genreID int) error
	UpdateGenre(genreID int, input model.UpdateGenreInput) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
	Book
	Genre
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
		Book:          NewBookService(repos.Book),
		Genre:         NewGenreService(repos.Genre),
	}
}
