package secretprotocolheader

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// createPublishFixHeader constructs an octet (8-bit long byte) based on its three arguments and the fix QoS setting.
func createPublishFixHeader(isFirstAttempt bool, isBroadcasted bool, isSecure bool) byte {
	header := byte(0b01001000)
	if isFirstAttempt {
		header |= 0b00010000
	}
	if isBroadcasted {
		header |= 0b00000010
	}
	if isSecure {
		header |= 0b00000001
	}
	return header
}
