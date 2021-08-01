package freqtrade

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/kamontat/fthelper/metric/v4/src/connection"
	"github.com/kamontat/fthelper/metric/v4/src/constants"
	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/datatype"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
)

type Connection struct {
	Cluster string
	Config  *Config

	base     *url.URL
	version  string
	username string
	password string

	cache  *caches.Service
	logger *loggers.Logger
}

func (c *Connection) ExpireAt(name string) string {
	return c.Config.Cache.Get(name)
}

func (c *Connection) QueryValues(name string) url.Values {
	return c.Config.Query.Get(name)
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
	caches.Global.Increase(constants.FTCONN_CALL + c.Cluster)
	resp, err := http.DefaultClient.Do(req)

	if err == nil {
		caches.Global.Increase(constants.FTCONN_CALL_SUCCESS + c.Cluster)
	} else {
		caches.Global.Increase(constants.FTCONN_CALL_FAILURE + c.Cluster)
	}

	if err != nil {
		c.logger.Debug("error: %s", err.Error())
	}
	return resp, err
}

func (c *Connection) Cache(name string, expireAt string, fn func() (interface{}, error)) (interface{}, error) {
	caches.Global.Increase(constants.FTCONN_CACHE_TOTAL + c.Cluster)
	err := c.cache.Fetch(name, func(o interface{}) (interface{}, error) {
		caches.Global.Increase(constants.FTCONN_CACHE_MISS + c.Cluster)
		return fn()
	}, expireAt)
	var data = c.cache.Get(name).Data
	// ensure that if err is nil data must not be nil
	if data == nil {
		err = fmt.Errorf("receive data as nil value (%s)", name)
	}

	if err != nil {
		c.logger.Warn(err.Error())
	}

	return data, err
}

func (c *Connection) GET(name string, query url.Values, target interface{}) error {
	var resp, err = c.Connect("GET", name, query, nil)
	if err != nil {
		return err
	}

	return toJson(resp, target)
}

func (c *Connection) POST(name string, query url.Values, body io.Reader, target interface{}) error {
	var resp, err = c.Connect("GET", name, query, body)
	if err != nil {
		return err
	}

	return toJson(resp, target)
}

func (c *Connection) String() string {
	var output = `
Connection: %s (%s)
  cache: %s
  query: %s
`

	return fmt.Sprintf(
		output,
		c.base.String(),
		c.Cluster,
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
		Cluster: freqtrade.So("cluster", ""),
		Config:  newConfig(freqtrade),

		base:     baseUrl,
		version:  freqtrade.So("apiver", "v1"),
		username: freqtrade.So("username", "freqtrader"),
		password: freqtrade.So("password", ""),

		cache:  cache,
		logger: loggers.Get("freqtrade", "connection"),
	}, nil
}

func NewConnections(data maps.Mapper) ([]*Connection, error) {
	var clusters = data.Ai("clusters")
	if len(clusters) <= 0 {
		var conn, err = NewConnection(data, caches.New())
		if err != nil {
			return make([]*Connection, 0), err
		}
		return []*Connection{conn}, nil
	}

	var logger = loggers.Get("freqtrade", "connection")
	logger.Info("currently you using 'multiple clusters mode' which still on alpha release")
	var connections = make([]*Connection, 0)
	for _, raw := range clusters {
		var cluster = datatype.ForceString(raw)
		var raw, err = data.Mi("cluster").Gets(cluster, strings.ToLower(cluster))
		if err != nil {
			return nil, err
		}
		var setting, _ = maps.ToMapper(raw)

		baseUrl, err := url.Parse(setting.So("url", "http://localhost:8080"))
		if err != nil {
			return nil, err
		}

		connections = append(connections, &Connection{
			Cluster: cluster,
			Config:  newConfig(data.Mi("freqtrade")),

			base:     baseUrl,
			version:  setting.So("apiver", "v1"),
			username: setting.So("username", "freqtrader"),
			password: setting.So("password", ""),

			cache:  caches.New(),
			logger: logger,
		})
	}

	return connections, nil
}

func ToConnection(conn connection.Http) *Connection {
	return conn.(*Connection)
}
