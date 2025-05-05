package main

import (
	"Adder/adders"
	"Adder/gates"
	"Adder/helplers"
	"fmt"
)

// todo: think about different types of input (not only binary);
// todo: convert input to binary, use FullAdder and then return the result in specific type

const subtraction = true

func main() {

	complementer := adders.NewComplementer(subtraction)
	adder := adders.Adder{}

	A := "110"
	B := "100"

	// if complementer.Inversion = 1 (true), then B will complement to once (inverted), if not - leave the same value
	complementOnesB := complementer.Complete(B)
	fmt.Println(complementOnesB)

	sum, carry := adder.AddBits(A, complementOnesB, helplers.BoolToBit(subtraction))

	// carry overflow or disappearance -> if [+] then overflow, if [-] then disappearance
	carry = gates.XOR(carry, helplers.BoolToBit(subtraction))

	fmt.Println("Sum:", sum)
	fmt.Println("Carry:", carry)

	fmt.Printf("Result: %d%s", carry, sum)
}
