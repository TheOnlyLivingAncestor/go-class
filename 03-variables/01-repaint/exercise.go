package repaint

import "fmt"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// repaintColor returns the complementary of the color received as argument or an error for an unknown error.
func repaintColor(color string) (string, error) {
	switch currentColor := color; currentColor {
	case "chartreuse":
		return "magenta", nil
	case "magenta":
		return "chartreuse", nil
	default:
		return "", fmt.Errorf("unknown color")
	}
}
