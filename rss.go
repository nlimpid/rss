package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
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
	r.Get("/v1/:dbname/:id", handler.GetDB)
	r.Get("/zhihu_image", handler.GetImage)
	logrus.Info("Rss run at :6334")
	logrus.Error(http.ListenAndServe(":6334", r))
}
