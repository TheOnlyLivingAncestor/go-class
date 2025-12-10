package pathsplit

import "path"

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// splitPath returns the directory component of a file path.
func splitPath(fullPath string) string {
	dir, _ := path.Split(fullPath)
	return dir
}
