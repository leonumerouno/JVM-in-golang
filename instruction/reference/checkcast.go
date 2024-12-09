package reference

import (
	"jvm/instruction/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

type CHECKCAST struct {
	base.Index16Instruction
}

func (self *CHECKCAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}

	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
