package store_test

import (
	"errors"
	"fmt"
	. "github.com/anton-bryukhov/url-shortener/pkg/store"
	"testing"
)

var data = map[int]string{
	64:   "www.google.com",
	512:  "www.facebook.com",
	1024: "www.amazon.com",
	2048: "www.youtube.com",
	4096: "en.wikipedia.org",
	8192: "www.yahoo.com",
}

type MockStore struct{}

func (m MockStore) Get(id int) (string, error) {
	url, ok := data[id]

	if ok {
		return url, nil
	}

	return "", errors.New("Not found")
}

func TestGet(t *testing.T) {
	store := Store{MockStore{}}
	tests := []struct {
		in  string
		out string
	}{
		{"ab", "www.google.com"},
		{"hp", "www.facebook.com"},
		{"pF", "www.amazon.com"},
		{"Gb", "www.youtube.com"},
		{"add", "en.wikipedia.org"},
		{"bhh", "www.yahoo.com"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Gets '%s' for '%s'", test.out, test.in), func(t *testing.T) {
			got, err := store.Get(test.in)

			if err != nil {
				t.Fatalf("error '%v' was returned", err)
			}

			if got != test.out {
				t.Fatalf("got: '%s', want: '%s'", got, test.out)
			}
		})
	}
}

func TestGetError(t *testing.T) {
	store := Store{MockStore{}}
	tests := []struct{ in string }{{"abc"}, {"def"}}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Returns error when '%v' not found", test.in), func(t *testing.T) {
			got, err := store.Get(test.in)

			if got != "" {
				t.Fatalf("returned '%s' instead of empty string", got)
			}

			if err == nil {
				t.Fatal("error was not returned")
			}
		})
	}
}
