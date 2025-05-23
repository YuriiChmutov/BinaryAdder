package main

import (
	"Adder/adders"
	"Adder/gates"
	"Adder/helplers"
	"Adder/latches"
	"Adder/selectors"
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
	//testingLatch8()
	//testingSelector()

	testingAdderWithLatch()
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

	trigger.Update(1, 0, 1, 0) // setting 1 and enabling to save
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 0, 0, 0) // no matter S and R, NOT ENABLE to save new state
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 0, 1, 0) // S and R = 0, it's a Saving mode in trigger, the state doesn't change
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 1, 0, 0) // enable = 0, then state doesn't change
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 1, 1, 0) // resetting trigger, Q become = 0
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 0, 1, 0) // setting 1 and enabling to save
	fmt.Println(trigger.GetQ())
}

func testingDTrigger() {
	trigger := triggers.NewDTrigger()

	trigger.Update(1, 1, 0)
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 1, 0)
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 0, 0)
	fmt.Println(trigger.GetQ())

	trigger.Update(1, 1, 0)
	fmt.Println(trigger.GetQ())

	trigger.Update(0, 0, 1)
	fmt.Println(trigger.GetQ())
}

func testingLatch8() {
	latch := latches.NewLatch8()
	var enable int8 = 1

	latch.Update("10101010", enable, 0)

	data := latch.GetData()
	fmt.Println(data)

	latch.Reset()

	data = latch.GetData()
	fmt.Println(data)

	latch.Update("11", enable, 0)

	data = latch.GetData()
	fmt.Println(data)

	latch.Update("11111111", enable, 0)

	data = latch.GetData()
	fmt.Println(data)
}

func testingSelector() {
	A := "10000001"
	B := "01111100"

	selector := selectors.NewSelector()
	result := selector.Select(A, B, 0)
	fmt.Println(result)
}

func testingAdderWithLatch() {
	latch := latches.NewLatch8()
	complementer := adders.NewComplementer(false)
	adder := adders.Adder{}
	//selector := selectors.NewSelector()

	a := "00000001"
	b := "00000001"
	c := "01111100"

	b = complementer.Complete(b)

	sum1, carry1 := adder.AddBits(a, b, 0)
	latch.Update(sum1, 1, 0)

	fmt.Println(latch.GetData())

	sum2, carry2 := adder.AddBits(latch.GetData(), c, carry1)

	//latch.Reset()
	//latch.Update(sum2, 1, 0)
	//fmt.Println(carry2, latch.GetData())
	fmt.Println(carry2, sum2)
}
