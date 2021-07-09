package repository

import (
	"io"
	"net/http"
)

type Http interface {
	Do(method string, url string, body io.Reader) (*http.Response, error)
	WithCache(name string, expireAt string, fn func() (interface{}, error)) (interface{}, error)

	GET(name string, target interface{}) error

	POST(name string, body io.Reader, target interface{}) error
}
