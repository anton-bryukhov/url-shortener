package handler

import "net/http"

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		shortenHandler(w, r)
	} else {
		unshortenHandler(w, r)
	}
}