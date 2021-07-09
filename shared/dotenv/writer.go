package dotenv

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/kamontat/fthelper/shared/fs"
)

// Marshal outputs the given environment as a dotenv-formatted environment file.
// Each line is in the format: KEY="VALUE" where VALUE is backslash-escaped.
func Marshal(envMap map[string]string) (string, error) {
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		if d, err := strconv.Atoi(v); err == nil {
			lines = append(lines, fmt.Sprintf(`%s=%d`, k, d))
		} else {
			lines = append(lines, fmt.Sprintf(`%s="%s"`, k, doubleQuoteEscape(v)))
		}
	}
	sort.Strings(lines)
	return strings.Join(lines, "\n"), nil
}

// Write serializes the given environment and writes it to a file
func Write(envMap map[string]string, file fs.FileSystem) error {
	content, err := Marshal(envMap)
	if err != nil {
		return err
	}

	return file.Write([]byte(content + "\n"))
}
