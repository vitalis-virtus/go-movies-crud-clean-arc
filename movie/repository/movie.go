package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/models"
)

type MovieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetAllMovie() []models.Movie {
	var Movies []models.Movie
	r.db.Find(&Movies)
	return Movies
}

func (r *MovieRepository) CreateMovie(m *models.Movie) (*models.Movie, error) {
	r.db.NewRecord(m)

	err := r.db.Create(&m).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *MovieRepository) GetMovieById(id int64) (*models.Movie, *gorm.DB) {
	var getMovie models.Movie
	db := r.db.Where("ID=?", id).Find((&getMovie))

	return &getMovie, db
}

func (r *MovieRepository) DeleteMovie(id int64) (*models.Movie, error) {
	var movie models.Movie
	err := r.db.Where("ID=?", id).Delete(movie).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}
