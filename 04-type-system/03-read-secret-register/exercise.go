package readsecretregister

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

// parseChannelControlRegister constructs 4 octets (8-bit long uint) based on the parameter register.
func parseChannelControlRegister(charCtrl uint32) (uint8, uint8, uint8, uint8) {
	// charCtrl a következő mezőket tartalmaza: RX_PCODE, TX_PCODE, RX_CHAN, TX_CHAN
	// return in the following order: TX_CHAN, RX_CHAN, RX_PCODE, TX_PCODE
	TX_CHAN := uint8(charCtrl)
	RX_CHAN := uint8(charCtrl >> 8)
	TX_PCODE := uint8(charCtrl >> 16)
	RX_PCODE := uint8(charCtrl >> 24)

	return TX_CHAN, RX_CHAN, RX_PCODE, TX_PCODE

}
