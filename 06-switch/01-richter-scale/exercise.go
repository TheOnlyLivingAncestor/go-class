package richterscale

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// describeEarthquake returns the "description" of a given magnitude value on the Richter scale.
func describeEarthquake(magnitude float32) string {
	switch {
	case magnitude < 2.0:
		return "micro"
	case 2.0 <= magnitude && magnitude < 3.0:
		return "very minor"
	case 3.0 <= magnitude && magnitude < 4.0:
		return "minor"
	case 4.0 <= magnitude && magnitude < 5.0:
		return "light"
	case 5.0 <= magnitude && magnitude < 6.0:
		return "moderate"
	case 6.0 <= magnitude && magnitude < 7.0:
		return "strong"
	case 7.0 <= magnitude && magnitude < 8.0:
		return "major"
	case 8.0 <= magnitude && magnitude < 10.0:
		return "great"
	default:
		return "massive"
	}

}
