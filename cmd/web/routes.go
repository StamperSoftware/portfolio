package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"time"
)

func (app *application) routes() http.Handler {

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	bootstrapFileServer := http.FileServer(http.Dir("./node_modules/bootstrap/"))
	mux.Handle("/node_modules/bootstrap/*", http.StripPrefix("/node_modules/bootstrap", bootstrapFileServer))

	popperFileServer := http.FileServer(http.Dir("./node_modules/@popperjs/"))
	mux.Handle("/node_modules/@popperjs/*", http.StripPrefix("/node_modules/@popperjs", popperFileServer))

	mux.Get("/", app.Home)
	mux.Get("/{page}", app.Page)
	return mux
}
