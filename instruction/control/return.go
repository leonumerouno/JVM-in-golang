package control

import (
	"jvm/instruction/base"
	"jvm/rtda"
)

type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct {
	base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retval := currentFrame.OperandStack().PopRef()
	invokerFrame.OperandStack().PushRef(retval)
}

type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retval := currentFrame.OperandStack().PopDouble()
	invokerFrame.OperandStack().PushDouble(retval)
}

type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retval := currentFrame.OperandStack().PopFloat()
	invokerFrame.OperandStack().PushFloat(retval)
}

type IRETURN struct {
	base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retval := currentFrame.OperandStack().PopInt()
	invokerFrame.OperandStack().PushInt(retval)
}

type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.CurrentFrame()
	retval := currentFrame.OperandStack().PopLong()
	invokerFrame.OperandStack().PushLong(retval)
}
