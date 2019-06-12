package handler

import (
	"errors"
	"net/http"
)

type Payload struct {
	URL string
}

func assertHTTPMethod(w http.ResponseWriter, r *http.Request, method string) error {
	if r.Method != method {
		return errors.New("Method not allowed")
	}

	return nil
}