package adders

import "Adder/gates"

type FullAdder struct{}

func (f FullAdder) Add(A int8, B int8, carryIn int8) (int8, int8) {
	ha := HalfAdder{}
	sum1, carry1 := ha.Add(A, B)               // 8 gates
	sumResult, carry2 := ha.Add(carryIn, sum1) // 8 gates
	carryResult := gates.OR(carry2, carry1)    // 2 gates
	return sumResult, carryResult              // 18 gates
}
