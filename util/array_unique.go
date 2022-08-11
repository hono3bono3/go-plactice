package util

func ArrayUnique[T comparable](arr []T) []T {
	m := make(map[T]bool)
	var uniq []T

	for _, ele := range arr {
		if !m[ele] {
			m[ele] = true
			uniq = append(uniq, ele)
		}
	}
	return uniq
}
