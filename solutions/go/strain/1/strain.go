package strain

func Keep[T any](collection []T, predicate func(T) bool) []T {
	result := []T{}
	for _, num := range collection {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func Discard[T any](collection []T, predicate func(T) bool) []T {
	result := []T{}
	for _, num := range collection {
		if !predicate(num) {
			result = append(result, num)
		}
	}
	return result
}
