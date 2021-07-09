package maps_test

import (
	"reflect"
	"testing"

	"github.com/kamontat/fthelper/shared/maps"
)

type TestCase struct {
	Name     string
	InputA   map[string]interface{}
	InputB   map[string]interface{}
	Strategy maps.Mapper

	Expected map[string]interface{}
}

func TestMergeNormalJson(t *testing.T) {
	var emptyStrategy = make(maps.Mapper)
	testcases := []TestCase{
		{
			Name:     "merge 2 empty json",
			InputA:   make(map[string]interface{}),
			InputB:   make(map[string]interface{}),
			Strategy: emptyStrategy,

			Expected: make(map[string]interface{}),
		},
		{
			Name:   "merge when first json is empty",
			InputA: make(map[string]interface{}),
			InputB: map[string]interface{}{
				"string": "non-empty string",
			},
			Strategy: emptyStrategy,

			Expected: map[string]interface{}{
				"string": "non-empty string",
			},
		},
		{
			Name: "merge when second json is empty",
			InputA: map[string]interface{}{
				"string": "non-empty string",
			},
			InputB:   make(map[string]interface{}),
			Strategy: emptyStrategy,

			Expected: map[string]interface{}{
				"string": "non-empty string",
			},
		},
		{
			Name: "merge normal json",
			InputA: map[string]interface{}{
				"string": "text",
				"int":    25,
			},
			InputB: map[string]interface{}{
				"float64": float64(88),
				"int8":    int8(3),
				"bool":    false,
			},
			Strategy: emptyStrategy,

			Expected: map[string]interface{}{
				"string":  "text",
				"int":     25,
				"float64": float64(88),
				"int8":    int8(3),
				"bool":    false,
			},
		},
		{
			Name: "merge with replace string/int/bool",
			InputA: map[string]interface{}{
				"data":            "string",
				"replace_float":   float32(99),
				"replace_bool":    true,
				"replace_string":  "old-value",
				"change_datatype": 12,
			},
			InputB: map[string]interface{}{
				"replace_float":   float64(77),
				"replace_bool":    false,
				"replace_string":  "newvalue",
				"change_datatype": "12",
			},
			Strategy: emptyStrategy,

			Expected: map[string]interface{}{
				"data":            "string",
				"replace_float":   float64(77),
				"replace_bool":    false,
				"replace_string":  "newvalue",
				"change_datatype": "12",
			},
		},
		{
			Name: "nested json",
			InputA: map[string]interface{}{
				"data": map[string]interface{}{
					"nested":   "string",
					"override": "A",
				},
			},
			InputB: map[string]interface{}{
				"bData": "hello",
				"data": map[string]interface{}{
					"override": "B",
					"bNested":  "bString",
				},
			},
			Strategy: emptyStrategy,

			Expected: map[string]interface{}{
				"bData": "hello",
				"data": map[string]interface{}{
					"nested":   "string",
					"bNested":  "bString",
					"override": "B",
				},
			},
		},
		{
			Name: "nested array",
			InputA: map[string]interface{}{
				"data": map[string]interface{}{
					"nested": []string{"a", "b", "c"},
				},
			},
			InputB: map[string]interface{}{
				"data": map[string]interface{}{
					"nested": []string{"d", "e", "f"},
				},
			},
			Strategy: emptyStrategy,

			Expected: map[string]interface{}{
				"data": map[string]interface{}{
					"nested": []interface{}{"a", "b", "c", "d", "e", "f"},
				},
			},
		},
		{
			Name: "override json with strategy",
			InputA: map[string]interface{}{
				"data": map[string]interface{}{
					"a": "a",
				},
			},
			InputB: map[string]interface{}{
				"data": map[string]interface{}{
					"b": "b",
				},
			},
			Strategy: maps.Mapper{
				"data": maps.MERGER_OVERRIDE,
			},

			Expected: map[string]interface{}{
				"data": maps.Mapper{
					"b": "b",
				},
			},
		},
		{
			Name: "override nested json with strategy",
			InputA: map[string]interface{}{
				"data": map[string]interface{}{
					"a": "a",
					"aNested": map[string]interface{}{
						"value": "a",
					},
					"nested": map[string]interface{}{
						"a": "a",
					},
				},
			},
			InputB: map[string]interface{}{
				"data": map[string]interface{}{
					"nested": map[string]interface{}{
						"b": "b",
					},
				},
			},
			Strategy: maps.Mapper{
				"data": maps.Mapper{
					"nested": maps.MERGER_OVERRIDE,
				},
			},

			Expected: map[string]interface{}{
				"data": map[string]interface{}{
					"a": "a",
					"aNested": map[string]interface{}{
						"value": "a",
					},
					"nested": maps.Mapper{
						"b": "b",
					},
				},
			},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.Name, func(t *testing.T) {
			var expected = testcase.Expected
			var actual = maps.Merge(
				testcase.InputA,
				testcase.InputB,
				testcase.Strategy,
			)

			if !reflect.DeepEqual(actual, expected) {
				t.Errorf("%v != %v", actual, expected)
			}
		})
	}
}
