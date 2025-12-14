package channelbroadcaster

import "context"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func channelBroadcast(ctx context.Context, input <-chan any, outputs []chan<- any) {

	//The function should start its own goroutine
	go func() {
		defer func() {
			for _, ch := range outputs {
				close(ch)
			}
		}()

		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-input:
				if !ok {
					return
				}
				for _, ch := range outputs {
					ch <- v
				}
			}
		}

	}()
}
