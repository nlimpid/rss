package main

import (
	"net/http"

	"github.com/nlimpid/rss/handler"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/zhihu/:articleName", handler.GetArticle)

	http.ListenAndServe(":6334", r)
}
