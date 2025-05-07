package triggers

import (
	"Adder/gates"
	"fmt"
)

type RSTrigger struct {
	Q, notQ int8 // Q - is output with a lamp; !Q - adds Q and always is opposite to Q
}

func NewRSTrigger() *RSTrigger {
	return &RSTrigger{Q: 0, notQ: 1}
}

// update function without signal which allows changing the state

//func (rs *RSTrigger) Update(S, R int8) { // S and R - enters; S - set, R - reset
//	// saving previous states
//	prevQ := rs.Q
//	prevNotQ := rs.notQ
//
//	// updating the logic using NOR
//	rs.Q = gates.NOR(R, prevNotQ)
//	rs.notQ = gates.NOR(S, prevQ)
//}

// update function with signal which allows changing the state

func (rs *RSTrigger) Update(S, R, enable, clear int8) { // S, R, enable - enters; S - set, R - reset, enable - save this bit;
	s := gates.AND(S, enable)
	r := gates.AND(R, enable)

	r = gates.OR(r, clear)

	newQ := rs.Q
	newNotQ := rs.notQ

	// Several operations — until it will be stabilized
	for i := 0; i < 3; i++ {
		newQ = gates.NOR(r, newNotQ)
		newNotQ = gates.NOR(s, newQ)
	}

	// Check for validity
	if newQ == newNotQ {
		fmt.Println("Warning: Invalid state Q == notQ")
		return
	}

	rs.Q = newQ
	rs.notQ = newNotQ

	fmt.Printf("S=%d R=%d enable=%d | s=%d r=%d | q=%d notQ=%d\n", S, R, enable, s, r, newQ, newNotQ)
}

func (rs *RSTrigger) GetAllStates() (int8, int8) {
	return rs.Q, rs.notQ
}

func (rs *RSTrigger) GetQ() int8 {
	return rs.Q
}
