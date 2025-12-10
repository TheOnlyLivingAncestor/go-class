package grades

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// gradeExam returns the grade of an exam with the given percentage
func gradeExam(percent float32) int {
	switch {
	case percent >= 90.0:
		return 5
	case percent >= 75.0:
		return 4
	case percent >= 50.0:
		return 3
	case percent >= 30.0:
		return 2
	default:
		return 0
	}
}
