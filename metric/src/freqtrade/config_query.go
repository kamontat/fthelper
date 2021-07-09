package freqtrade

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/kamontat/fthelper/shared/maps"
)

type QueryConfig struct {
	data maps.Mapper
}

func (c *QueryConfig) Get(key string) url.Values {
	var values = make(url.Values)

	var mapper = c.data.Mi(key)
	mapper.ForEach(func(key string, value interface{}) {
		values.Set(key, fmt.Sprintf("%v", value))
	})

	return values
}

func (c *QueryConfig) Json() string {
	var j, err = json.Marshal(c.data)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func newQueryConfig(query maps.Mapper) *QueryConfig {
	return &QueryConfig{
		data: query,
	}
}
