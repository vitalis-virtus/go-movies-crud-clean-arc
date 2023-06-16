package usecase

import (
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/models"
	"github.com/vitalis-virtus/go-movies-gallery/movie/repository"
)

type MovieUseCase struct {
	movieRepo repository.MovieRepository
}

func NewMovieUseCase(MovieRepo repository.MovieRepository) *MovieUseCase {
	return &MovieUseCase{
		movieRepo: MovieRepo,
	}
}

func (s *MovieUseCase) GetAllMovie() []models.Movie {
	result := s.movieRepo.GetAllMovie()
	return result
}

func (s *MovieUseCase) CreateMovie(m *models.Movie) (*models.Movie, error) {
	m, err := s.movieRepo.CreateMovie(m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (s *MovieUseCase) GetMovieById(id int64) (*models.Movie, *gorm.DB) {
	m, db := s.movieRepo.GetMovieById(id)

	return m, db
}

func (s *MovieUseCase) DeleteMovie(id int64) (*models.Movie, error) {
	m, err := s.movieRepo.DeleteMovie(id)

	if err != nil {
		return nil, err
	}

	return m, nil
}
