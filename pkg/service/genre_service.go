package service

import (
	"github.com/zhashkevych/todo-app/model"
	"github.com/zhashkevych/todo-app/pkg/repository"
)

type GenreService struct {
	repo repository.Genre
}

func NewGenreService(repo repository.Genre) *GenreService {
	return &GenreService{repo: repo}
}

func (r *GenreService) Create(genre model.Genre) (id int, err error){
	return r.repo.Create(genre)
}
func (r *GenreService) GetGenre() (genrelist [] model.Genre, err error)  {
	return r.repo.GetGenre()
}
func (r *GenreService) GetGenreById(genreID int) (genre model.Genre, err error)	{
	return r.repo.GetGenreById(genreID)
}
func(r *GenreService) DeleteGenre(genreID int) error{
	return r.repo.DeleteGenre(genreID)
}
func (r *GenreService) UpdateGenre(genreID int, input model.UpdateGenreInput) error {
	return r.repo.UpdateGenre(int(genreID),input)
}
