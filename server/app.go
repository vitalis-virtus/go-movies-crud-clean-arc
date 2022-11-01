package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/vitalis-virtus/go-movies-gallery/movie"
	"github.com/vitalis-virtus/go-movies-gallery/movie/delivery/handler"
	"github.com/vitalis-virtus/go-movies-gallery/movie/repository"
	"github.com/vitalis-virtus/go-movies-gallery/movie/usecase"
)

type App struct {
	httpServer *http.Server

	movieUC movie.MovieUseCase
}

func NewApp(db *gorm.DB) *App {

	movieRepo := repository.NewMovieRepository(db)

	return &App{
		movieUC: usecase.NewMovieUseCase(*movieRepo),
	}
}

func (a *App) Run(port string) error {
	router := mux.NewRouter()

	handler.RegisterHTTPEndpoints(router, a.movieUC)

	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)

}
