package movie

import (
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/models"
)

type MovieUseCase interface {
	CreateMovie(*models.Movie) (*models.Movie, error)
	GetMovieById(int64) (*models.Movie, *gorm.DB)
	GetAllMovie() []models.Movie
	DeleteMovie(int64) (*models.Movie, error)
}
