package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"gitlab.com/idoko/bucketeer/db"
)

var dbInstance db.Database

func homePage(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "/app/static/index.html")
}

func NewHandler(db db.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db

	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)

	router.Route("/items", items)
        router.Get("/", homePage)

	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
