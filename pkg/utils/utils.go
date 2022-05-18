package utils

func Filter[T any](slice []T, predicate func(T) bool) []T {
	newSlice := []T{}

	for _, item := range slice {
		if predicate(item) {
			newSlice = append(newSlice, item)
		}
	}
	return newSlice
}
