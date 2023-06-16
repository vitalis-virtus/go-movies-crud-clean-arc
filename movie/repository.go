package movie

import (
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/models"
)

type MovieRepository interface {
	GetAllMovie() []models.Movie
	GetMovieById(int64) (*models.Movie, *gorm.DB)
	CreateMovie(*models.Movie) (*models.Movie, error)
	DeleteMovie(int64) (*models.Movie, error)
}
