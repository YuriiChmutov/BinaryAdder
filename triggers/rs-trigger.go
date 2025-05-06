package triggers

import "Adder/gates"

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

func (rs *RSTrigger) Update(S, R, enable int8) { // S, R, enable - enters; S - set, R - reset, enable - save this bit;
	r := gates.AND(R, enable)
	s := gates.AND(enable, S)

	q := gates.XOR(r, rs.notQ)
	nq := gates.XOR(rs.Q, s)

	rs.Q = q
	rs.notQ = nq
}

func (rs *RSTrigger) GetAllStates() (int8, int8) {
	return rs.Q, rs.notQ
}

func (rs *RSTrigger) GetQ() int8 {
	return rs.Q
}
