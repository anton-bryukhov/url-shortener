package store

import (
	"github.com/anton-bryukhov/url-shortener/pkg/converter"
)

type Source interface {
	Get(int) (string, error)
}

type Store struct {
	DataSource Source
}

func (s *Store) Get(url string) (string, error) {
	id := converter.ConvertToID(url)
	return s.DataSource.Get(id)
}
