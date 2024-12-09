package lang

import (
	"jvm/native"
	"jvm/rtda"
	"jvm/rtda/heap"
)

func init() {
	native.Register("java/lang/String", "intern", "()Ljava/lang/String;", intern)
}

func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
