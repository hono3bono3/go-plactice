package util

func InArray[T comparable](val T, arr []T) (exists bool, index int) {
	exists = false
	index = -1
	for i, v := range arr {
		if val == v {
			exists = true
			index = i
			return exists, index
		}
	}
	return exists, index
}
