package clusters

var index = 0

func GetIndex() int {
	index++
	return index
}
