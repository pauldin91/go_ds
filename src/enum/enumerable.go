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

func Reduce[T any](src []T, reduce func(T)) {
	for _, e := range src {
		reduce(e)
	}
}

func Skip[T any](src []T, skip int) []T {
	if len(src) > skip {
		return src[skip : len(src)-1]
	}
	return make([]T, 0)
}

func Take[T any](src []T, take int) []T {
	if len(src) > take {
		return src[0:take]
	}
	return src
}
