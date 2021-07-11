package maps

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/kamontat/fthelper/shared/datatype"
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
		} else if array, ok := datatype.ToArray(v); ok {
			copied[k] = utils.CloneArray(array)
		} else {
			copied[k] = v
		}
	}

	return copied
}

// Normalize will remove key listed in removed.
// warning: input will be mutation to remove data
// if you don't want that, copy first.
func Normalize(input Mapper, removed []string) Mapper {
	for key, value := range input {
		for _, remove := range removed {
			if key == remove {
				// fmt.Printf("removing key=%s in %v\n", key, input)
				delete(input, key)
				continue
			}
		}
		if output, ok := ToMapper(value); ok {
			Normalize(output, removed)
		}
	}
	return input
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

func ToJson(m interface{}) ([]byte, error) {
	return json.Marshal(m)
}

func ToFormatJson(m interface{}) ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
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
