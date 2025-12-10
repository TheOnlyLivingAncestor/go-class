package digits

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func multiplyDigits(n int) int {
	product := 1
	for n > 0 {
		product *= n % 10
		n = n / 10
	}
	return product
}
