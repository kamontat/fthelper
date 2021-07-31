package utils

import (
	"regexp"
	"strconv"
)

func prenameToNumber(name string) int {
	switch name {
	case "alpha":
		return 1
	case "beta":
		return 2
	case "rc":
		return 3
	case "":
		return 0
	default:
		return -1
	}
}

// Support version string
// `v0.0.1`
// `v0.1.0`
// `v1.0.0`
// `v1.2.3`
// `v1.2.3-beta.1`
// `v1.2.3-alpha.1`
// `v1.2.3-rc.1`
func VersionNumber(version string) float64 {
	var regex = `v?(?P<Major>\d+).(?P<Minor>\d+).(?P<Patch>\d+)-?(?P<Prename>\w+)?.?(?P<Prerelease>\d+)?`
	var compRegEx = regexp.MustCompile(regex)
	match := compRegEx.FindStringSubmatch(version)
	if len(match) < 3 {
		return -1
	}

	var major, _ = strconv.Atoi(match[1])
	var minor, _ = strconv.Atoi(match[2])
	var patch, _ = strconv.Atoi(match[3])

	var base = float64((major * 10000) + (minor * 100) + patch)

	var prename = prenameToNumber(match[4])
	if prename < 0 {
		return -1
	}

	if prename > 0 {
		var prerelease, _ = strconv.Atoi(match[5])
		base = base + (float64(prename) / float64(10)) + (float64(prerelease) / float64(1000))
	}

	return base
}
