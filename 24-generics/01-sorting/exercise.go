package sorting

import "sort"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
type Orderable interface {
	~uint8 | ~int32
}

func sortSlice[T Orderable](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	sort.Slice(result, func(i, j int) bool {
		return result[i] < result[j]
	})

	return result
}
