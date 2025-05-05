package adders

import "Adder/gates"

type HalfAdder struct{}

func (h HalfAdder) Add(A int8, B int8) (int8, int8) {
	sum := gates.XOR(A, B)   // 6 gates
	carry := gates.AND(A, B) // 2 gates
	return sum, carry        // 8 gates
}
