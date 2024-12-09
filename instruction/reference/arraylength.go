package reference

import (
	"jvm/instruction/base"
	"jvm/rtda"
)

type ARRAYLENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAYLENGTH) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}

	arrLen := arrRef.ArrayLength()
	stack.PushInt(arrLen)
}
