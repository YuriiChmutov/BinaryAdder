package adder

import "Adder/gates"

func HalfAdder(A int8, B int8) (int8, int8) {
	sum := gates.XOR(A, B)   // 6 gates
	carry := gates.AND(A, B) // 2 gates
	return sum, carry        // 8 gates
}

func FullAdder(A int8, B int8, carryIn int8) (int8, int8) {
	sum1, carry1 := HalfAdder(A, B) // 8 gates

	sumResult, carry2 := HalfAdder(carryIn, sum1) // 8 gates

	carryResult := gates.OR(carry2, carry1) // 2 gates

	return sumResult, carryResult // 18 gates
}
