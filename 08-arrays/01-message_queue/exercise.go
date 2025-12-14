package messagequeue

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// messagequeue returns the an array constructed from the arguments
func messageQueue(a, b, c string) [3]string {
	return [3]string{c, a, b}
}
