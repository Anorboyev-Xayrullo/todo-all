package service

import (
	"github.com/zhashkevych/todo-app/model"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (repo *BookService) CreateBook(book model.Books) error {
	return repo.repo.CreateBook(book)
}
func (repo *BookService) 	GetBookById(bookId int) (book model.Books, err error){
	return repo.repo.GetBookById(bookId)
}
func(repo *BookService) GetBookList() (bookList []model.Books, err error){
	return repo.repo.GetBookList()
}
func (repo *BookService) GetBooksByGenre(genre string) (book model.UpdateGenreOutput, err error) {
	return repo.repo.GetBooksByGenre(genre)
}
func (repo *BookService) UpdateBook(bookId int, input model.UpdateBook) error {
	return repo.repo.UpdateBook(bookId,input)
}
func (repo *BookService) DeleteBook(bookId int) error {
	return repo.repo.DeleteBook(bookId)
}