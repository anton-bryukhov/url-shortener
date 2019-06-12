package handler

import (
	"encoding/json"
	"fmt"
	"github.com/anton-bryukhov/url-shortener/pkg/converter"
	"github.com/anton-bryukhov/url-shortener/pkg/store"
	"net/http"
	"strings"
)

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if err := assertHTTPMethod(w, r, "POST"); err != nil {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
		return
	}

	var data Payload
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	shortURL := converter.ConvertToURL(store.GetSize() + 1)

	
	if strings.HasPrefix(data.URL, "http") {
		store.Put(data.URL)
	} else {
		store.Put(fmt.Sprintf("http://%s", data.URL))
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(Payload{fmt.Sprintf("http://%s/%s", r.Host, shortURL)})
}
