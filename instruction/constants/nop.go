package constants

import (
	"jvm/instruction/base"
	"jvm/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {

}
