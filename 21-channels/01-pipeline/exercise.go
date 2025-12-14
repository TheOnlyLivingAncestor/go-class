package pipeline

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// generator (function)
func generator(nums []int) <-chan int {
	//pushes its elements to an integer channel
	channel := make(chan int)
	go func() {
		defer close(channel)
		for _, num := range nums {
			channel <- num
		}
	}()
	return channel
}

// adder (function)
func adder(in <-chan int) <-chan float32 {
	//increments digits by 1 of integers read from the in channel
	//writes them to a float32 channel

	channel := make(chan float32)
	go func() {
		defer close(channel)
		for num := range in {
			channel <- float32(num + 1)
		}
	}()
	return channel
}

// collector (function)
func collector(in <-chan float32) []float32 {
	out := []float32{}
	for num := range in {
		out = append(out, num)
	}
	return out
}
