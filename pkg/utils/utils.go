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

func Separate[T any](slice []T, predicate func(T) bool) ([]T, []T) {
	selected := Filter(slice, predicate)
	removed := Filter(slice, func(t T) bool { return !predicate(t) })

	return selected, removed
}

func Reduce[T any, O any](slice []T, predicate func(O, T) O, intialValue O) O {
	output := intialValue
	for _, item := range slice {
		output = predicate(output, item)
	}
	return output
}

func Sum(slices []int, initialValue int) int {
	return Reduce(slices, func(o int, t int) int { return o + t }, initialValue)
}

func Map[T any, O any](slices []T, predicate func(T) O) []O {
	output := []O{}
	for _, item := range slices {
		output = append(output, predicate(item))
	}
	return output
}
