package enum

func Map[T any, K any](src []T, mapping func(T) K) []K {
	result := make([]K, 0)
	for _, e := range src {
		result = append(result, mapping(e))
	}
	return result
}

func Filter[T any](src []T, filter func(T) bool) []T {
	result := make([]T, 0)
	for _, e := range src {
		if filter(e) {
			result = append(result, e)
		}
	}
	return result
}
