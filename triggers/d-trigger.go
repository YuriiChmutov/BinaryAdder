package triggers

import "Adder/gates"

type DTrigger struct {
	rsTrigger *RSTrigger
}

func NewDTrigger() *DTrigger {
	return &DTrigger{rsTrigger: NewRSTrigger()}
}

// the same as RSTrigger, but it's unnecessary to have S and R the same values

func (dt *DTrigger) Update(D, enable, clear int8) {
	S := D
	R := gates.NOT(D)

	dt.rsTrigger.Update(S, R, enable, clear)
}

func (dt *DTrigger) GetQ() int8 {
	return dt.rsTrigger.GetQ()
}
