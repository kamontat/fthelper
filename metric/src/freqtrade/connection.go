package freqtrade

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kamontat/fthelper/metric/src/connection"
	"github.com/kamontat/fthelper/metric/src/constants"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type Connection struct {
	Config *Config

	base     *url.URL
	version  string
	username string
	password string

	cache  *caches.Service
	logger *loggers.Logger
}

func (c *Connection) request(method, name string, query url.Values, body io.Reader) (*http.Request, error) {
	url, err := buildPath(c.base, c.version, name)
	if err != nil {
		return nil, err
	}

	url.RawQuery = query.Encode()
	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.username, c.password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func (c *Connection) Connect(method string, url string, query url.Values, body io.Reader) (*http.Response, error) {
	var req, err = c.request(method, url, query, body)
	if err != nil {
		return nil, err
	}
	c.logger.Debug("request: %s %s", req.Method, req.URL.String())
	caches.Global.Increase(constants.FTCONN_CALL)
	resp, err := http.DefaultClient.Do(req)

	if err == nil {
		caches.Global.Increase(constants.FTCONN_CALL_SUCCESS)
	} else {
		caches.Global.Increase(constants.FTCONN_CALL_FAILURE)
	}

	if err != nil {
		c.logger.Debug("error: %s", err.Error())
	}
	return resp, err
}

func (c *Connection) Cache(name string, expireAt string, fn func() (interface{}, error)) (interface{}, error) {
	caches.Global.Increase(constants.FTCONN_CACHE_TOTAL)
	err := c.cache.Fetch(name, func(o interface{}) (interface{}, error) {
		caches.Global.Increase(constants.FTCONN_CACHE_MISS)
		return fn()
	}, expireAt)

	if err != nil {
		c.logger.Warn(err.Error())
	}

	return c.cache.Get(name).Data, err
}

func (c *Connection) GET(name string, query url.Values, target interface{}) error {
	var resp, err = c.Connect("GET", name, query, nil)
	if err != nil {
		return err
	}

	err = toJson(resp, target)
	c.logger.Debug("response body: %v", target)

	return err
}

func (c *Connection) POST(name string, query url.Values, body io.Reader, target interface{}) error {
	var resp, err = c.Connect("GET", name, query, body)
	if err != nil {
		return err
	}

	err = toJson(resp, target)
	c.logger.Debug("response body: %v", target)

	return err
}

func (c *Connection) String() string {
	return fmt.Sprintf(
		"\nConnection: \n  cache: %s\n  query: %s",
		c.Config.Cache.Json(),
		c.Config.Query.Json(),
	)
}

func NewConnection(data maps.Mapper, cache *caches.Service) (*Connection, error) {
	var freqtrade = data.Mi("freqtrade")
	var baseUrl, err = url.Parse(freqtrade.So("url", "http://localhost:8080"))
	if err != nil {
		return nil, err
	}

	return &Connection{
		Config: newConfig(freqtrade),

		base:     baseUrl,
		version:  freqtrade.So("version", "v1"),
		username: freqtrade.So("username", "freqtrader"),
		password: freqtrade.So("password", ""),

		cache:  cache,
		logger: loggers.Get("freqtrade", "connection"),
	}, nil
}

func ToConnection(conn connection.Http) *Connection {
	return conn.(*Connection)
}
