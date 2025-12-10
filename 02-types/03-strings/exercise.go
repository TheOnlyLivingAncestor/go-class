package strings

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

func multilineString() string {
	return `some
multiline
string`
}

func stringLen(s string) int {
	return len(s)
}

func trimFirstChar(s string) string {
	if stringLen(s) == 0 {
		return s
	}
	return s[1:]
}

func trimLastChar(s string) string {
	if stringLen(s) == 0 {
		return s
	}
	return s[:stringLen(s)-1]
}

func swapFirstChar(s string) string {
	if stringLen(s) == 0 {
		return s
	}
	return string("A" + trimFirstChar(s))
}

func swapLastChar(s string) string {
	if stringLen(s) == 0 {
		return s
	}
	return string(trimLastChar(s) + "A")
}

func prependChar(s string) string {
	return string("A" + s)
}

func appendChar(s string) string {
	return string(s + "A")
}
