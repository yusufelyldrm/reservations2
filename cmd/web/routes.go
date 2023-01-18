package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yusufelyldrm/reservaiton2/internal/config"
	"github.com/yusufelyldrm/reservaiton2/internal/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	/*mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))*/

	//mux is a router
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Use(SessionLoad)
	mux.Use(NoSurf)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/generals-quarter", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)

	//static file server
	fileServer := http.FileServer(http.Dir("./static"))

	//mux.Handle is a chi method that takes a path and a handler
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux

}
