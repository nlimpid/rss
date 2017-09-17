package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/go-chi/cors"
	"github.com/nlimpid/rss/handler"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	xx := cors.New(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(xx.Handler)

	// r.Use(middleware.RequestID)
	// r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", handler.GetHome)
	r.Get("/zhihu/{articleName}", handler.GetArticle)
	r.Get("/zhuanlan/{articleName}", handler.ZhihuZhuanlan)
	r.Get("/v1/:dbname/:id", handler.GetDB)
	r.Get("/zhihu_image", handler.GetImage)
	logrus.Info("Rss run at :6334")
	logrus.Error(http.ListenAndServe(":6334", r))
}
