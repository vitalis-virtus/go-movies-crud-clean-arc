package handler

import (
	"github.com/vitalis-virtus/go-movies-gallery/movie"

	"github.com/gorilla/mux"
)

func RegisterHTTPEndpoints(router *mux.Router, uc movie.MovieUseCase) {
	h := NewMovieHandler(uc)

	router.HandleFunc("/movie/", h.GetAllMovie).Methods("GET")
	router.HandleFunc("/movie/", h.CreateMovie).Methods("POST")
	router.HandleFunc("/movie/{movieId}", h.GetMovieById).Methods("GET")
	router.HandleFunc("/movie/{movieId}", h.UpdateMovie).Methods("PUT")
	router.HandleFunc("/movie/{movieId}", h.DeleteMovie).Methods("DELETE")
}
