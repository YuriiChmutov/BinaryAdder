package adders

import (
	"fmt"
)

type Adder struct{}

func (a Adder) AddBits(A, B string) (string, int8) {
	var sum string
	var carry int8 = 0

	fa := FullAdder{}

	for i := len(A) - 1; i >= 0; i-- {
		bitA := int8(A[i] - '0')
		bitB := int8(B[i] - '0')

		bitSum, c := fa.Add(bitA, bitB, carry)
		carry = c

		sum = fmt.Sprintf("%d", bitSum) + sum
	}

	return sum, carry
}
