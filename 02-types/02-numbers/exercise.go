package calculator

import (
	"math"
	"strconv"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func gmean(x, y float64) int {
	return int(math.Round(math.Sqrt(x * y)))
}

func gmeanString(x, y string) (int, error) {
	var x_float, err_x = strconv.ParseFloat(x, 64)
	if err_x != nil {
		return 0, err_x
	}
	var y_float, err_y = strconv.ParseFloat(y, 64)
	if err_y != nil {
		return 0, err_y
	}
	return gmean(x_float, y_float), nil
}
