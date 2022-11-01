package api

// package api

import (
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/internal/handler"
	"github.com/vitalis-virtus/go-movies-gallery/internal/repository"
	"github.com/vitalis-virtus/go-movies-gallery/internal/service"
)

// SetupRouter create router and configurate all routes
func SetupRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	movieRepo := repository.NewMovieRepository(db)
	movieService := service.NewMovieService(*movieRepo)
	movieHandler := handler.NewMovieHandler(movieService)

	router.HandleFunc("/movie/", movieHandler.GetAllMovie).Methods("GET")
	router.HandleFunc("/movie/", movieHandler.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{movieId}", movieHandler.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", movieHandler.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", movieHandler.DeleteMovie).Methods("DELETE")

	return router
}
