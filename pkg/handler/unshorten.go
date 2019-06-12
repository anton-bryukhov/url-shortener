package handler

import (
	"net/http"
	"github.com/anton-bryukhov/url-shortener/pkg/converter"
	"github.com/anton-bryukhov/url-shortener/pkg/store"
)

func unshortenHandler(w http.ResponseWriter, r *http.Request) {
	if err := assertHTTPMethod(w, r, "GET"); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	shortURL := r.URL.Path[1:]
	id := converter.ConvertToID(shortURL)

	originalURL, err := store.Get(id)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
