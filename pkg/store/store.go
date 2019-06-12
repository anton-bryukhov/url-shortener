package store

import "errors"

var urls []string

func Get(id int) (string, error) {
	if (id - 1 >= len(urls)) {
		return "", errors.New("Not found")
	}

	return urls[id - 1], nil
}

func GetSize() int {
	return len(urls)
}

func Put(url string) {
	urls = append(urls, url)
}
