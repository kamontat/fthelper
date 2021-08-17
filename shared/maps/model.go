package maps

import (
	"github.com/kamontat/fthelper/shared/datatype"
)

type Mapper map[string]interface{}

func (m Mapper) Z(key string) (interface{}, bool) {
	d, ok := m[key]
	return d, ok
}

func (m Mapper) Ze(key string) (interface{}, error) {
	if d, ok := m.Z(key); ok {
		return d, nil
	}
	return nil, missError(m, key)
}

func (m Mapper) Zo(key string, i interface{}) interface{} {
	if d, ok := m.Z(key); ok {
		return d
	}
	return i
}

func (m Mapper) Zi(key string) interface{} {
	return m.Zo(key, nil)
}

func (m Mapper) M(key string) (Mapper, bool) {
	return ToMapper(m[key])
}

func (m Mapper) Me(key string) (Mapper, error) {
	if d, ok := m.M(key); ok {
		return d, nil
	}

	return New(), convertError(key, m[key], "map[string]interface{}")
}

func (m Mapper) Mo(key string, i Mapper) Mapper {
	if d, ok := m.M(key); ok {
		return d
	}
	return i
}

func (m Mapper) Mi(key string) Mapper {
	return m.Mo(key, New())
}

func (m Mapper) A(key string) ([]interface{}, bool) {
	return datatype.ToArray(m[key])
}

func (m Mapper) Ae(key string) ([]interface{}, error) {
	if d, ok := m.A(key); ok {
		return d, nil
	}

	return make([]interface{}, 0), convertError(key, m[key], "array")
}

func (m Mapper) Ao(key string, i []interface{}) []interface{} {
	if d, ok := m.A(key); ok {
		return d
	}
	return i
}

func (m Mapper) Ai(key string) []interface{} {
	return m.Ao(key, make([]interface{}, 0))
}

func (m Mapper) S(key string) (string, bool) {
	return datatype.ToString(m[key])
}

func (m Mapper) Se(key string) (string, error) {
	if d, ok := m.S(key); ok {
		return d, nil
	}

	return "", convertError(key, m[key], "string")
}

func (m Mapper) So(key string, i string) string {
	if d, ok := m.S(key); ok {
		return d
	}
	return i
}

func (m Mapper) Si(key string) string {
	return m.So(key, "")
}

func (m Mapper) F(key string) (float64, bool) {
	return datatype.ToFloat(m[key])
}

func (m Mapper) Fe(key string) (float64, error) {
	if d, ok := m.F(key); ok {
		return d, nil
	}

	return 0, convertError(key, m[key], "float64")
}

func (m Mapper) Fo(key string, i float64) float64 {
	if d, ok := m.F(key); ok {
		return d
	}
	return i
}

func (m Mapper) Fi(key string) float64 {
	return m.Fo(key, 0)
}

func (m Mapper) I(key string) (int64, bool) {
	return datatype.ToInt(m[key])
}

func (m Mapper) Ie(key string) (int64, error) {
	if d, ok := m.I(key); ok {
		return d, nil
	}

	return 0, convertError(key, m[key], "int64")
}

func (m Mapper) Io(key string, i int64) int64 {
	if d, ok := m.I(key); ok {
		return d
	}
	return i
}

func (m Mapper) Ii(key string) int64 {
	return m.Io(key, 0)
}

func (m Mapper) N(key string) (float64, bool) {
	if f, ok := datatype.ToFloat(m[key]); ok {
		return f, true
	}
	if i, ok := datatype.ToInt(m[key]); ok {
		return float64(i), true
	}
	return -1, false
}

func (m Mapper) Ne(key string) (float64, error) {
	if d, ok := m.N(key); ok {
		return d, nil
	}

	return 0, convertError(key, m[key], "number")
}

func (m Mapper) No(key string, i float64) float64 {
	if d, ok := m.N(key); ok {
		return d
	}
	return i
}

func (m Mapper) Ni(key string) float64 {
	return m.No(key, -1)
}

func (m Mapper) B(key string) (bool, bool) {
	return datatype.ToBool(m[key])
}

func (m Mapper) Be(key string) (bool, error) {
	if d, ok := m.B(key); ok {
		return d, nil
	}

	return false, convertError(key, m[key], "bool")
}

func (m Mapper) Bo(key string, i bool) bool {
	if d, ok := m.B(key); ok {
		return d
	}
	return i
}

func (m Mapper) Bi(key string) bool {
	return m.Bo(key, false)
}

func (m Mapper) Struct(target interface{}) error {
	return ToStruct(m, target)
}

func (m Mapper) Size() int {
	return len(m)
}

func (m Mapper) IsEmpty() bool {
	return m.Size() == 0
}

func (m Mapper) Has(key string) bool {
	var d, err = Get(m, key)
	return err == nil && d != nil
}

func (m Mapper) Set(key string, value interface{}) Mapper {
	Set(m, key, value)
	return m
}

func (m Mapper) Get(key string) (interface{}, error) {
	return Get(m, key)
}

func (m Mapper) Gets(keys ...string) (interface{}, error) {
	return Gets(m, keys...)
}

func (m Mapper) ForEach(fn ForEachFn) {
	ForEach(m, fn)
}

func (m Mapper) Copy() Mapper {
	return Copy(m)
}

// Keys will return key as dot notation on every map level
func (m Mapper) Keys() []string {
	var keys = make([]string, 0)
	m.ForEach(func(key string, value interface{}) {
		keys = append(keys, key)
	})
	return keys
}

func (m Mapper) String() string {
	var json, err = ToJson(m)
	if err != nil {
		return err.Error()
	}
	return string(json)
}
