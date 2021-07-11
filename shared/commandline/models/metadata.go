package models

import (
	"fmt"

	"github.com/kamontat/fthelper/shared/maps"
)

type Metadata struct {
	Name    string
	Version string
	Commit  string
	Date    string
	BuiltBy string
}

func (m *Metadata) String() string {
	return fmt.Sprintf("%s: %s (%s)", m.Name, m.Version, m.Commit)
}

func (m *Metadata) ToMapper() maps.Mapper {
	return maps.Mapper{
		"name":    m.Name,
		"version": m.Version,
		"commit":  m.Commit,
		"date":    m.Date,
		"buildby": m.BuiltBy,
	}
}

func EmptyMetadata() *Metadata {
	return &Metadata{
		Name:    "",
		Version: "",
		Commit:  "",
		Date:    "",
		BuiltBy: "",
	}
}
