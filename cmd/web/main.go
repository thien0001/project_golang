package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	// "github.com/alexedwards/scs/v2"

	"github.com/alexedwards/scs/v2"
	"github.com/project/golang/jebthien/pkg/config"
	"github.com/project/golang/jebthien/pkg/handlers"
	"github.com/project/golang/jebthien/pkg/render"
)

const PortNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

// main is the main application funtion
func main() {

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	n, err := fmt.Fprintf(w, "Hello World!")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	// })

	// change this to true when in production

	// https://github.com/alexedwards/scs

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", PortNumber))
	// _ = http.ListenAndServe(PortNumber, nil)

	srv := &http.Server{
		Addr:    PortNumber,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
