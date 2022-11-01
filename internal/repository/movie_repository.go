package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/internal/entity"
)

type MovieRepository struct {
	db *gorm.DB
}

type IMovieRepository interface {
	GetAllMovie() []entity.Movie
	GetMovieById(int64) (*entity.Movie, *gorm.DB)
	CreateMovie(*entity.Movie) (*entity.Movie, error)
	DeleteMovie(int64) (*entity.Movie, error)
}

func NewMovieRepository(db *gorm.DB) *MovieRepository {
	return &MovieRepository{
		db: db,
	}
}

func (r *MovieRepository) GetAllMovie() []entity.Movie {
	var Movies []entity.Movie
	r.db.Find(&Movies)
	return Movies
}

func (r *MovieRepository) CreateMovie(m *entity.Movie) (*entity.Movie, error) {
	r.db.NewRecord(m)

	err := r.db.Create(&m).Error

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *MovieRepository) GetMovieById(id int64) (*entity.Movie, *gorm.DB) {
	var getMovie entity.Movie
	db := r.db.Where("ID=?", id).Find((&getMovie))

	return &getMovie, db
}

func (r *MovieRepository) DeleteMovie(id int64) (*entity.Movie, error) {
	var movie entity.Movie
	err := r.db.Where("ID=?", id).Delete(movie).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}
