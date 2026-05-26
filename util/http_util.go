package util

import (
	"net/http"
	"net/url"
)

var (
	client       *http.Client
	Transport    *http.Transport
	SchemePrefix = "http://"
)

func init() {
	Transport = &http.Transport{
		MaxIdleConnsPerHost: 1024,
	}
	client = &http.Client{Transport: Transport}
}

func Post(url string, values url.Values) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Get(url string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func DownloadUrl(fileUrl string) (filename string, content []byte, e error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
