package handler

import "net/http"

// GetHome test for the api
func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello rss and power by docker!"))
}
