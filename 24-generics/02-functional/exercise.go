package functional

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func filter[E any](input []E, predicate func(E) bool) []E {
	var result []E
	for _, elem := range input {
		if predicate(elem) {
			result = append(result, elem)
		}
	}
	return result
}
