package closures

import "errors"

// DO NOT REMOVE THIS COMMENT
//
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate
func proxy(f func(string) int) func(string) (int, error) {
	on := true

	return func(s string) (int, error) {
		if on {
			res := f(s)
			on = !on
			return res, nil
		}
		on = !on
		return 0, errors.New("Error")
	}
}
