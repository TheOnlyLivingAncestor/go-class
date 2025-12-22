package filteringdata

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// filterData filters a slice based in an index slice.
func filterData(keys []string, indices []int) [10]string {
	filtered_array := [10]string{}

	var array_index int
	if len(keys) == len(indices) {
		for i := 0; i < len(keys); i++ {
			if indices[i] > 4 {
				filtered_array[array_index] = keys[i]
				array_index++
			}
		}
	}

	return filtered_array
}
