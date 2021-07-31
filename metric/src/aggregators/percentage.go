package aggregators

// calculate percent change from previous to current
func PercentChange(previous, current float64) (float64, bool) {
	var increase = current - previous
	return Percentage(increase, previous)
}

// calculate percentage of value in total
func Percentage(value, total float64) (float64, bool) {
	if total == 0 {
		return 0, false
	}

	return value / total, true
}
