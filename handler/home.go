package handler

import "net/http"

func GetHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello rss and power by docker"))
}
