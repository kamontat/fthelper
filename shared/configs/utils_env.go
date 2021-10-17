package configs

import (
	"fmt"
	"strings"
)

func IsEnvKey(k string) bool {
	return strings.HasPrefix(k, ENV_PREFIX+"_")
}

func IsCEnvKey(k string) bool {
	return strings.HasPrefix(k, ENV_CUSTOM_PREFIX+"_")
}

func EnvToKey(c string) (string, bool) {
	if (!IsCEnvKey(c) && !IsEnvKey(c)) || strings.Contains(c, "___") {
		return "", false
	}

	var output string = ""
	if IsCEnvKey(c) {
		output = strings.Replace(c, ENV_CUSTOM_PREFIX+"_", "", 1)
	} else if IsEnvKey(c) {
		output = strings.Replace(c, ENV_PREFIX+"_", "", 1)
	}

	output = strings.ReplaceAll(output, "__", ".")
	output = strings.ReplaceAll(output, "_", "-")
	if IsCEnvKey(c) {
		output = "_." + output
	}

	return strings.ToLower(output), true
}

// NOTE: ou should not use _ as a key, since we cannot parse underscroll back after it created environment variable
// Please use - (dash) instead
func KeyToEnv(key string) string {
	if key[0] == '_' {
		var result = fmt.Sprintf("%s-%s", ENV_CUSTOM_PREFIX, key[2:])
		return strings.ToUpper(
			strings.ReplaceAll(strings.ReplaceAll(result, ".", "__"), "-", "_"),
		)
	} else {
		var result = fmt.Sprintf("%s-%s", ENV_PREFIX, key)
		return strings.ToUpper(
			strings.ReplaceAll(strings.ReplaceAll(result, ".", "__"), "-", "_"),
		)
	}
}
