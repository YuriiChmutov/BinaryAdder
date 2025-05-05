package main

import (
	"Adder/adder"
	"fmt"
)

// todo: think about different types of input (not only binary);
// todo: convert input to binary, use FullAdder and then return the result in specific type

func main() {
	//A := "10101010" // 170
	//B := "11001100" // 204

	A := "111111111111"
	B := "100000000000"

	sum, carry := AddBits(A, B)
	fmt.Println("Sum:", sum)
	fmt.Println("Carry:", carry)

	fmt.Printf("Result: %d%s", carry, sum)
}

func AddBits(A, B string) (string, int8) {
	var sum string
	var carry int8 = 0

	for i := len(A) - 1; i >= 0; i-- {
		bitA := int8(A[i] - '0')
		bitB := int8(B[i] - '0')

		bitSum, c := adder.FullAdder(bitA, bitB, carry)
		carry = c

		sum = fmt.Sprintf("%d", bitSum) + sum
	}

	return sum, carry
}
