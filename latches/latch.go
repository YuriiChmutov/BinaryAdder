package latches

import (
	"Adder/triggers"
	"fmt"
	"strings"
)

type Latch8 struct {
	triggers [8]*triggers.DTrigger
}

func NewLatch8() *Latch8 {
	var latch Latch8
	for i := range latch.triggers {
		latch.triggers[i] = triggers.NewDTrigger()
	}
	return &latch
}

func (latch *Latch8) Update(data string, enable int8) {

	if len(data) > 8 {
		panic("Input data can't be longer than 8 bits")
	}

	if len(data) < 8 {
		padding := strings.Repeat("0", 8-len(data))
		data = padding + data
	}

	for i := range latch.triggers {
		item := int8(data[i] - '0')
		latch.triggers[i].Update(item, enable)
	}
}

func (latch *Latch8) Reset() {
	for i := range latch.triggers {
		latch.triggers[i].Update(0, 1)
	}
}

func (latch *Latch8) GetData() string {
	var result string
	for i := range latch.triggers {
		result += fmt.Sprintf("%d", latch.triggers[i].GetQ())
	}

	return result
}
