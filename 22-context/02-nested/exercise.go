package subtask

import (
	"context"
	"errors"
	"time"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// INSERT YOUR CODE HERE
func StartTask(ctx context.Context) (result string, err error) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	result_channel := make(chan string, 1)
	error_channel := make(chan error, 1)

	go func() {
		defer close(result_channel)
		defer close(error_channel)

		res, err := SubTask(ctxWithTimeout)
		if err != nil {
			error_channel <- err
			return
		}
		result_channel <- res
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case err := <-error_channel:
		if err != nil {
			return "", err
		}
	case res := <-result_channel:
		return "Main task status:" + res, nil
	}

	return "", errors.New("unexpected state")
}

func SubTask(ctx context.Context) (result string, err error) {
	select {
	case <-time.After(200 * time.Millisecond):
		return "Subtask completed successfully", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
