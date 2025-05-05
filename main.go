package main

import (
	"Adder/adders"
	"fmt"
)

// todo: think about different types of input (not only binary);
// todo: convert input to binary, use FullAdder and then return the result in specific type

func main() {
	complementer := adders.NewComplementer(true)
	fmt.Println(complementer.Inversion)

	//A := "10101010" // 170
	//B := "11001100" // 204

	A := "111111111111"
	B := "100000000000"

	adder := adders.Adder{}

	sum, carry := adder.AddBits(A, B)
	fmt.Println("Sum:", sum)
	fmt.Println("Carry:", carry)

	fmt.Printf("Result: %d%s", carry, sum)
}
