package utils

import (
	"fmt"
	"strings"
)

func TrimString(str string, limit int) string {
	if len(str) > limit {
		return str[0:limit]
	}
	return str
}

func JoinString(sep string, args ...string) string {
	var base strings.Builder

	var first = true
	for _, arg := range args {
		if arg != "" {
			if first {
				base.WriteString(arg)
				first = false
			} else {
				base.WriteString(fmt.Sprintf("%s%s", sep, arg))
			}
		}
	}

	return base.String()
}

type MaskLevel int

const (
	NONE MaskLevel = iota
	LOW
	MEDIUM
	HIGH
)

func maskString(dot rune, prefix int, s string, suffix int, max int) string {
	counter := 0
	rs := []rune(s)
	for i := prefix; i < len(rs)-suffix; i++ {
		if max > 0 && counter > max {
			rs = append(rs[:i], rs[i+1:]...)
			i-- // reverse index
		} else {
			rs[i] = dot
		}

		counter++
	}
	return string(rs)
}

// mask last 20% of string
func lowMaskString(s string) string {
	var percent = 0.2

	var size = len(s)
	var mask = int(percent * float64(size))
	return maskString('*', size-mask, s, 0, -1)
}

// mask last 80% of string but remain last 15% as unmask
func mediumMaskString(s string) string {
	var percent = 0.8
	var remain = 0.15

	var size = len(s)
	var mask = int(percent * float64(size))
	var remainMask = int(remain * float64(size))

	return maskString('*', size-mask, s, remainMask, -1)
}

// mask last 90% of string but remain last 5% as unmask
// and max dot to only 3
func highMaskString(s string) string {
	var percent = 0.90
	var remain = 0.05
	var maxdot = 3

	var size = len(s)
	var mask = int(percent * float64(size))
	var remainMask = int(remain * float64(size))

	return maskString('.', size-mask, s, remainMask, maxdot)
}

func MaskString(s string, sensitive MaskLevel) string {
	switch sensitive {
	case LOW:
		return lowMaskString(s)
	case MEDIUM:
		return mediumMaskString(s)
	case HIGH:
		return highMaskString(s)
	}

	return s
}
