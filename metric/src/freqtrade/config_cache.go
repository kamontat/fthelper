package freqtrade

import (
	"encoding/json"

	"github.com/kamontat/fthelper/shared/maps"
)

type CacheConfig struct {
	data maps.Mapper
}

var defaultValues map[string]string = map[string]string{
	"ping":      "1s",
	"balance":   "3m",
	"version":   "10m",
	"profit":    "5m",
	"whitelist": "15m",
}

func (c *CacheConfig) GetOrElse(key string, def string) string {
	return c.data.So(key, def)
}

func (c *CacheConfig) Get(key string) string {
	var def, ok = defaultValues[key]
	if !ok {
		def = "1m"
	}

	return c.GetOrElse(key, def)
}

func (c *CacheConfig) Json() string {
	var j, err = json.Marshal(c.data)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}
