package utils

import (
	"io/ioutil"
	"net/http"
)

func Request(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()
	return ioutil.ReadAll(r.Body)
}
