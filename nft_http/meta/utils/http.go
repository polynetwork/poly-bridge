package utils

import (
	"io/ioutil"
	"net/http"
	"poly-bridge/conf"
	"time"
)

func Request(url string) ([]byte, error) {
	if conf.GlobalConfig.HttpConfig.Timeout == 0 {
		conf.GlobalConfig.HttpConfig.Timeout = 30
	}
	client := &http.Client{
		Timeout: time.Duration(conf.GlobalConfig.HttpConfig.Timeout) * time.Second,
	}
	r, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() { _ = r.Body.Close() }()
	return ioutil.ReadAll(r.Body)
}
