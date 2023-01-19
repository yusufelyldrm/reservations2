package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/yusufelyldrm/reservations2/internal/config"
	"github.com/yusufelyldrm/reservations2/internal/handlers"
	"github.com/yusufelyldrm/reservations2/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// create a template cache
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// assign the template cache to the app config
	app.TemplateCache = tc

	// assign the session to the app config
	app.UseCache = false

	// create a new repo
	repo := handlers.NewRepo(&app)
	// create a new handlers
	handlers.NewHandlers(repo)

	// create a new template set
	render.NewTemplates(&app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	log.Println(" Started port")

	//_ = http.ListenAndServe(portNumber, nil)

	// create a new server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	// start the server
	err = srv.ListenAndServe()
	log.Fatal(err)
}
