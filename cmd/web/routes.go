package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yusufelyldrm/bookings/pkg/config"
	"github.com/yusufelyldrm/bookings/pkg/handlers"
	"net/http"
)

func Routes(app *config.AppConfig) http.Handler {
	/*mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))*/

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	// mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	fileServer := http.FileServer(http.Dir("./static"))

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
