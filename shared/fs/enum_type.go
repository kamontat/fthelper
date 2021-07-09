package fs

type Type string

const (
	FILE      Type = "file"
	DIRECTORY Type = "directory"
)

func ToType(s string) (Type, bool) {
	switch s {
	case "file", "f":
		return FILE, true
	case "directory", "dir", "d":
		return DIRECTORY, true
	}

	return FILE, false
}
