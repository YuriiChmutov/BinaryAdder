package main

import (
	"Adder/adders"
	"Adder/gates"
	"Adder/helplers"
	"Adder/triggers"
	"fmt"
)

// todo: think about different types of input (not only binary);
// todo: convert input to binary, use FullAdder and then return the result in specific type

const subtraction = false

func main() {
	//testingComlementerAndAdder()

	testingRsTrigger()
}

func testingComlementerAndAdder() {
	complementer := adders.NewComplementer(subtraction)
	adder := adders.Adder{}

	A := "10000001"
	B := "01111100"

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

func testingRsTrigger() {
	trigger := triggers.RSTrigger{}

	trigger.Update(1, 0) // setting
	fmt.Println(trigger.Output())

	trigger.Update(0, 0) // saving
	fmt.Println(trigger.Output())

	trigger.Update(0, 1) // reset
	fmt.Println(trigger.Output())

	trigger.Update(0, 0) // saving
	fmt.Println(trigger.Output())
}
