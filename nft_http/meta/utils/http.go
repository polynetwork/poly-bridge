package utils

import (
	"io/ioutil"
	"net/http"
	"time"
)

func Request(url string) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * 12,
	}
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()
	return ioutil.ReadAll(r.Body)
}
