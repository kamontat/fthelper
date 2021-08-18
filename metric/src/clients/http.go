package clients

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/loggers"
	"github.com/kamontat/fthelper/shared/maps"
	"github.com/kamontat/fthelper/shared/utils"
)

type Http struct {
	Enabled bool
	Cluster string
	Query   *Query

	base     *url.URL
	version  string
	username string
	password string

	logger *loggers.Logger
}

func (c *Http) Request(method, name string, body io.Reader) (*http.Request, error) {
	if !c.Enabled {
		return nil, fmt.Errorf("http client of '%s' is disabled", c.Cluster)
	}

	url, err := c.base.Parse(fmt.Sprintf("/api/%s/%s", c.version, name))
	if err != nil {
		return nil, err
	}

	query := c.Query.Get(name)
	c.logger.Debug("found request query for %s: '%v'", name, query.Encode())

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

func (c *Http) Call(method, name string, body io.Reader) (*http.Response, error) {
	var req, err = c.Request(method, name, body)
	if err != nil {
		return nil, err
	}
	c.logger.Debug("request: %s %s", req.Method, req.URL.String())

	caches.Global.Increase(HTTP_CALL + c.Cluster)
	resp, err := http.DefaultClient.Do(req)
	if err == nil &&
		resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusAccepted {
		err = fmt.Errorf("freqtrade (%s) return error (status %s)", c.Cluster, resp.Status)
	}

	if err == nil {
		caches.Global.Increase(HTTP_SUCCESS + c.Cluster)
	} else {
		c.logger.Debug("error: %s", err.Error())
		caches.Global.Increase(HTTP_FAILURE + c.Cluster)
	}

	return resp, err
}

func (c *Http) GET(name string, target interface{}) error {
	var resp, err = c.Call(http.MethodGet, name, nil)
	if err != nil {
		return err
	}

	return toJson(resp, target)
}

func (c *Http) POST(name string, body io.Reader, target interface{}) error {
	var resp, err = c.Call(http.MethodPost, name, body)
	if err != nil {
		return err
	}

	return toJson(resp, target)
}

func (c *Http) String() string {
	if !c.Enabled {
		return "disabled"
	}

	return fmt.Sprintf("%s '%s:%s'", c.base.String(), c.username, utils.MaskString(c.password, utils.MEDIUM))
}

func NewHttp(cluster string, config maps.Mapper) (*Http, error) {
	var enabled = config.Bo("enabled", true)

	urlString, err := config.Se("url")
	if err != nil {
		return nil, err
	}
	baseUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	apiver := config.So("apiver", "v1")

	username, err := config.Se("username")
	if err != nil {
		return nil, err
	}

	password, err := config.Se("password")
	if err != nil {
		return nil, err
	}

	return &Http{
		Enabled: enabled,
		Cluster: cluster,
		Query:   newQuery(config.Mi("query")),

		base:     baseUrl,
		version:  apiver,
		username: username,
		password: password,

		logger: loggers.Get("client", "http", cluster),
	}, nil
}
