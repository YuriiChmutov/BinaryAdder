package main

import (
	"Adder/adders"
	"Adder/gates"
	"Adder/helplers"
	"Adder/latches"
	"Adder/triggers"
	"fmt"
)

// todo: think about different types of input (not only binary);
// todo: convert input to binary, use FullAdder and then return the result in specific type

const subtraction = false

func main() {
	//testingComlementerAndAdder()
	//testingRsTrigger()
	//testingDTrigger()

	testingLatch8()
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
	trigger := triggers.NewRSTrigger()
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 0, 1) // setting 1 and enabling to save
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 0, 0) // no matter S and R, NOT ENABLE to save new state
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 0, 1) // S and R = 0, it's a Saving mode in trigger, the state doesn't change
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 1, 0) // enable = 0, then state doesn't change
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 1, 1) // resetting trigger, Q become = 0
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 0, 1) // setting 1 and enabling to save
	fmt.Println(trigger.GetQ())
}

func testingDTrigger() {
	trigger := triggers.NewDTrigger()

	trigger.Update(1, 1)
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 1)
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 0)
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 1)
	fmt.Println(trigger.GetQ())
}

func testingLatch8() {
	latch := latches.NewLatch8()
	var enable int8 = 1

	latch.Update("10101010", enable)

	data := latch.GetData()
	fmt.Println(data)

	latch.Reset()

	data = latch.GetData()
	fmt.Println(data)

	latch.Update("11", enable)

	data = latch.GetData()
	fmt.Println(data)

	latch.Update("11111111", enable)

	data = latch.GetData()
	fmt.Println(data)
}
