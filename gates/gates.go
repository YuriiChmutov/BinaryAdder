package gates

// For carry

func AND(A int8, B int8) int8 {
	if A == 1 && B == 1 {
		return 1
	}
	return 0
}

func OR(A int8, B int8) int8 {
	if A == 1 || B == 1 {
		return 1
	}
	return 0
}

func NOR(A int8, B int8) int8 {
	if A == 0 && B == 0 {
		return 1
	}
	return 0
}

func NOT(A int8) int8 {
	if A == 0 {
		return 1
	}
	return 0
}

func NAND(A int8, B int8) int8 {
	if A == 1 && B == 1 {
		return 0
	}
	return 1
}

func BUFFER(A int8) int8 {
	return A
}

//For sum

func XOR(A int8, B int8) int8 {
	return AND(OR(A, B), NAND(A, B)) // OR * NAND
}

func EQUAL(A int8, B int8) int8 {
	if A == B {
		return 1
	}
	return 0
}
