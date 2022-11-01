package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vitalis-virtus/go-movies-gallery/internal/entity"
	"github.com/vitalis-virtus/go-movies-gallery/internal/service"
	"github.com/vitalis-virtus/go-movies-gallery/pkg/utils"
)

type MovieHandler struct {
	movieService service.IMovieService
}

func NewMovieHandler(movieService service.IMovieService) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}

func (h *MovieHandler) GetAllMovie(w http.ResponseWriter, r *http.Request) {
	newMovies := h.movieService.GetAllMovie()
	res, _ := json.Marshal(newMovies)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	newMovie := &entity.Movie{}

	// we parse movie data from request in understandable form for go
	utils.ParseBody(r, newMovie)

	m, err := h.movieService.CreateMovie(newMovie)

	if err != nil {
		log.Print(err)
		return
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	movieDetails, _ := h.movieService.GetMovieById(ID)

	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	deletedMovie, err := h.movieService.DeleteMovie(ID)

	if err != nil {
		log.Print(err)
		return
	}

	res, _ := json.Marshal(deletedMovie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var updateMovie = &entity.Movie{}
	utils.ParseBody(r, updateMovie)
	vars := mux.Vars(r)
	movieId := vars["movieId"]
	ID, err := strconv.ParseInt(movieId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}

	movieDetails, db := h.movieService.GetMovieById(ID)

	if err != nil {
		log.Print(err)
		return
	}

	// we are checking for updated fields from user and replace this fields in moviesDetails
	if updateMovie.Name != "" {
		movieDetails.Name = updateMovie.Name
	}
	if updateMovie.Director != "" {
		movieDetails.Director = updateMovie.Director
	}
	if updateMovie.Rating != "" {
		movieDetails.Rating = updateMovie.Rating
	}

	db.Save(&movieDetails)
	res, _ := json.Marshal(movieDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
