package xtemplates

import (
	"fmt"
	"text/template"
	"time"
)

// toDuration will parse input format to input unit
// possible unit is ms (millisecond), s (second), m (minute), and h (hour)
func toDuration(format, unit string) (float64, error) {
	duration, err := time.ParseDuration(format)
	if err != nil {
		return -1, err
	}

	switch unit {
	case "ms":
		return float64(duration.Milliseconds()), nil
	case "s":
		return duration.Seconds(), nil
	case "m":
		return duration.Minutes(), nil
	case "h":
		return duration.Hours(), nil
	default:
		return -1, fmt.Errorf("cannot convert to unit '%s'", unit)
	}
}

// dayToDuration will parse input number as days to output base on unit
func dayToDuration(day int, unit string) (float64, error) {
	var hour = day * 24
	return toDuration(fmt.Sprintf("%dh", hour), unit)
}

var durationFuncs template.FuncMap = map[string]interface{}{
	"toDuration":    toDuration,
	"dayToDuration": dayToDuration,
}
