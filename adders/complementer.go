package adders

import (
	"Adder/gates"
	"Adder/helplers"
	"fmt"
)

// The type responsible for what value needs to be added to the input number
// To obtain the maximum number in digits

// In the case of binary computing, this is called "One's Complement"

type Complementer struct {
	Inversion int8
}

func NewComplementer(inversion bool) *Complementer {
	binaryInversionFormat := helplers.BoolToBit(inversion)
	return &Complementer{binaryInversionFormat}
}

// XOR operation is used because if the Inversion value is 1, then the input value will be inverted.
// If it is 0, then the value will not change.

func (complementer Complementer) Complete(A string) string {
	var completed string

	for i := 0; i < len(A); i++ {
		bit := int8(A[i] - '0')

		completedBit := gates.XOR(complementer.Inversion, bit)

		completed += fmt.Sprintf("%d", completedBit)
	}

	return completed
}
