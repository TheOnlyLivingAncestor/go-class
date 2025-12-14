package channelmultiplexer

import (
	"context"
	"sync"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func channelMultiplex(ctx context.Context, inputs []chan any) chan any {
	output := make(chan any)

	var wg sync.WaitGroup

	for _, input := range inputs {
		wg.Add(1)
		go func(ch chan any) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					return
				case v, ok := <-ch:
					if !ok {
						return
					}
					output <- v
				}
			}
		}(input)
	}

	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}
