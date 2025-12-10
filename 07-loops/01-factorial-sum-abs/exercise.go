package factorial

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func calcAbs(n int) int {
	if n >= 0 {
		return n
	}
	return -1 * n
}
