package utilities

func Contains[T comparable](slice []T, search T) (found bool) {
	found = false
	for _, v := range slice {
		if v == search {
			found = true
			break
		}
	}
	return
}

func GetMapKeys[T comparable](maps map[T]interface{}) (slice []T) {
	for key := range maps {
		slice = append(slice, key)
	}
	return
}
