package adders

import (
	"Adder/gates"
	"fmt"
)

type Inverter struct{}

func (i Inverter) Invert(A string) string {
	var result string

	for j := 0; j < len(A); j++ {
		bit := int8(A[j] - '0')
		inverted := gates.NOT(bit)
		result += fmt.Sprintf("%d", inverted)
	}

	return result
}
