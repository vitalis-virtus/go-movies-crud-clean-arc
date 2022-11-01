package service

import (
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/internal/entity"
	"github.com/vitalis-virtus/go-movies-gallery/internal/repository"
)

type MovieService struct {
	movieRepo repository.MovieRepository
}

type IMovieService interface {
	CreateMovie(*entity.Movie) (*entity.Movie, error)
	GetMovieById(int64) (*entity.Movie, *gorm.DB)
	GetAllMovie() []entity.Movie
	DeleteMovie(int64) (*entity.Movie, error)
}

func NewMovieService(MovieRepo repository.MovieRepository) *MovieService {
	return &MovieService{
		movieRepo: MovieRepo,
	}
}

func (s *MovieService) GetAllMovie() []entity.Movie {
	result := s.movieRepo.GetAllMovie()
	return result
}

func (s *MovieService) CreateMovie(m *entity.Movie) (*entity.Movie, error) {
	m, err := s.movieRepo.CreateMovie(m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *MovieService) GetMovieById(id int64) (*entity.Movie, *gorm.DB) {
	m, db := s.movieRepo.GetMovieById(id)

	return m, db
}

func (s *MovieService) DeleteMovie(id int64) (*entity.Movie, error) {
	m, err := s.movieRepo.DeleteMovie(id)

	if err != nil {
		return nil, err
	}

	return m, nil
}
