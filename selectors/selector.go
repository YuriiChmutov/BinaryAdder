package selectors

import (
	"Adder/gates"
	"fmt"
)

// 2-Line-to-1-line Selector

type Selector struct{}

func NewSelector() *Selector {
	return &Selector{}
}

func (s *Selector) selectBit(A, B, S int8) int8 {
	part1 := gates.AND(gates.NOT(S), A)
	part2 := gates.AND(S, B)
	return gates.OR(part1, part2)
}

func (s *Selector) Select(A, B string, S int8) string {
	if len(A) != len(B) {
		panic("Input strings A and B must be of the same length")
	}

	var result string

	for i := 0; i < len(A); i++ {
		a := int8(A[i] - '0')
		b := int8(B[i] - '0')

		selected := s.selectBit(a, b, S)

		result += fmt.Sprintf("%d", selected)
	}

	return result
}
