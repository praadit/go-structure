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
