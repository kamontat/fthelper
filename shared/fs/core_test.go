package fs_test

import (
	"testing"

	"github.com/kamontat/fthelper/shared/fs"
	"github.com/kamontat/fthelper/shared/maps"
)

func TestCore(t *testing.T) {
	var v, err = fs.Build(maps.Mapper{
		"type": "file",
		"mode": string(fs.MULTIPLE),
		// "fullpath": []string{"{{ .a }}/extra", "/tmp"},
		"paths": []interface{}{
			[]string{
				"/tmp/", "hello/test", "world", "{{ .a }}",
			}, []string{
				"/tmp/", "hello/test", "world2",
			}, []string{
				"hello", "hello/test", "world2",
			},
		},
	}, maps.Mapper{
		"a": "/etc/freqtrade",
	})

	if err == nil {
		for _, fs := range v.Multiple() {
			t.Log(fs.Stat())
		}
	} else {
		t.Log(err)
	}
}
