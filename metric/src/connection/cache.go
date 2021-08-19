package connection

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/caches"
	"github.com/kamontat/fthelper/shared/errors"
	"github.com/kamontat/fthelper/shared/maps"
)

type cacheState string

const (
	CREATED = "created"
	INITIAL = "initial"
)

type cache struct {
	state cacheState
	err   error

	connector Connector
	service   *caches.Service
	config    maps.Mapper
}

func (c *cache) Cluster() string {
	return c.connector.Cluster()
}

func (c *cache) Initial() error {
	if c.state == CREATED {
		c.err = c.connector.Initial()
		c.state = INITIAL
	}

	return c.err
}

func (c *cache) Cleanup() error {
	if c.state == INITIAL {
		c.err = c.connector.Cleanup()
		c.state = CREATED
	}

	return c.err
}

// cache connector should never has parent
func (c *cache) Parent() Connector {
	return nil
}

// WithParent of cache will do nothing
// cache connector should be the top level of connector
func (c *cache) WithParent(parent Connector) Connector {
	return c
}

func (c *cache) Save(name string, data interface{}) Connector {
	c.connector.Save(name, data)
	return c
}

func (c *cache) Connect(name string) (interface{}, error) {
	var expiredAt, err = c.config.Se(name)
	if err != nil {
		return nil, fmt.Errorf("you forget to set cache duration for '%s'", name)
	}

	caches.Global.Increase(CACHE_TOTAL + c.Cluster())
	updated, _ := c.service.Fetch(name, func(o interface{}) (interface{}, error) {
		var data = c.service.Get(name)
		if data.IsExist() {
			// Save data only if data is exist
			c.connector.Save(name, data)
		}

		return c.connector.Connect(name)
	}, expiredAt)
	if updated {
		caches.Global.Increase(CACHE_MISS + c.Cluster())
	}

	var data = c.service.Get(name)
	return data.Get()
}

func (c *cache) ConnectAll() *errors.Handler {
	return c.connector.ConnectAll()
}

func (c *cache) String() string {
	return fmt.Sprintf("[with-cache] %v", c.connector.String())
}

func WithCache(c Connector, service *caches.Service, config maps.Mapper) Connector {
	var val = &cache{
		state: CREATED,
		err:   nil,

		connector: c,
		service:   service,
		config:    config,
	}

	c.WithParent(val)
	return val
}
