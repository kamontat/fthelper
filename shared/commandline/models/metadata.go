package models

import "fmt"

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

func EmptyMetadata() *Metadata {
	return &Metadata{
		Name:    "",
		Version: "",
		Commit:  "",
		Date:    "",
		BuiltBy: "",
	}
}
