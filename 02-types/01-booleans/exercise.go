package logicalops

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func inverse(b bool) bool {
	return !b
}

func or(x, y bool) bool {
	return x || y
}

func deMorgan(a, b bool) bool {
	return or(inverse(a), inverse(b))
}
