package handler

import (
	"log"
	"net/http"

	"github.com/pressly/chi"
)

func GetDB(w http.ResponseWriter, r *http.Request) {
	dbname := chi.URLParam(r, "dbname")
	id := chi.URLParam(r, "id")
	log.Println(dbname)
	log.Printf("id=%v\n", id)
}
