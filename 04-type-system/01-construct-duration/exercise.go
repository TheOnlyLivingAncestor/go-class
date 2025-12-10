package constructduration

import "time"

//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// constructTime constructs a `Time` instant based on its two arguments (arg1, arg2)
func constructDuration(arg1 int, arg2 int) time.Duration {
	//arg1 is the number of seconds, arg2 is the number of milliseconds
	return time.Duration(arg1*int(time.Second)) + time.Duration(arg2*int(time.Millisecond))
}
