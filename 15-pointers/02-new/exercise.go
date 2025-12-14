package pointernew

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func newValue() *string {

	return_value := new(string)
	*return_value = "test"

	return return_value
}
