package connection

import (
	"io"
	"net/http"
	"net/url"
)

type Http interface {
	Connect(method string, url string, query url.Values, body io.Reader) (*http.Response, error)
	Cache(name string, expireAt string, fn func() (interface{}, error)) (interface{}, error)

	GET(name string, query url.Values, target interface{}) error
	POST(name string, query url.Values, body io.Reader, target interface{}) error

	String() string
}
