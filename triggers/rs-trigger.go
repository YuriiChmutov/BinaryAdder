package triggers

import "Adder/gates"

type RSTrigger struct {
	Q, notQ int8 // Q - is output with a lamp; !Q - adds Q and always is opposite to Q
}

func (rs *RSTrigger) Update(S, R int8) { // S and R - enters; S - set, R - reset
	// saving previous states
	prevQ := rs.Q
	prevNotQ := rs.notQ

	// updating the logic using NOR
	rs.Q = gates.NOR(R, prevNotQ)
	rs.notQ = gates.NOR(S, prevQ)
}

func (rs *RSTrigger) Output() (int8, int8) {
	return rs.Q, rs.notQ
}
