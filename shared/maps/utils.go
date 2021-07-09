package maps

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kamontat/fthelper/shared/utils"
)

func Get(m map[string]interface{}, key string) (interface{}, error) {
	return getByKey(m, ToKeys(key))
}

func Gets(m map[string]interface{}, keys ...string) (interface{}, error) {
	for _, key := range keys {
		var v, err = getByKey(m, ToKeys(key))
		if err == nil {
			return v, nil
		}
	}

	return nil, fmt.Errorf("cannot found any data of keys (%v)", keys)
}

func Set(m map[string]interface{}, key string, value interface{}) {
	var keys = ToKeys(key)
	var length = len(keys)

	var counter, _ = ToMapper(m)
	for i, k := range keys {
		if i == length-1 {
			if value == nil {
				delete(counter, k)
			} else {
				counter[k] = value
			}
		} else {
			v, ok := counter.M(k)
			if !ok {
				counter[k] = make(map[string]interface{})
				counter, _ = counter.M(k)
			} else {
				counter = v
			}
		}
	}
}

func Copy(m map[string]interface{}) Mapper {
	var copied = make(Mapper)
	for k, v := range m {
		if mapper, ok := ToMapper(v); ok {
			copied[k] = Copy(mapper)
		} else if array, ok := utils.ToArray(v); ok {
			copied[k] = utils.CloneArray(array)
		} else {
			copied[k] = v
		}
	}

	return copied
}

func ForEach(m map[string]interface{}, fn ForEachFn) {
	forEach(m, []string{}, fn)
}

func ToKeys(key string) []string {
	return strings.Split(key, ".")
}

func ToMapper(d interface{}) (Mapper, bool) {
	// try with map[string]interface{}
	data, ok := d.(map[string]interface{})
	if ok {
		return data, ok
	}

	// try with Mapper
	data, ok = d.(Mapper)
	if ok {
		return data, ok
	}

	return make(Mapper), false
}

func ToJson(m map[string]interface{}) string {
	var j, err = json.Marshal(m)
	if err != nil {
		return err.Error()
	} else {
		return string(j)
	}
}

func FromJson(content []byte) (Mapper, error) {
	var empty = make(map[string]interface{})
	var err = json.Unmarshal(content, &empty)
	return Mapper(empty), err
}

func ToStruct(m map[string]interface{}, target interface{}) error {
	body, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, target)
	if err != nil {
		return err
	}

	return nil
}

func New() Mapper {
	return make(Mapper)
}
