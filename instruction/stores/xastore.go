package stores

import (
	"jvm/instruction/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

type AASTORE struct {
	base.NoOperandsInstruction
}

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()

	checkNotNil(arrRef)
	refs := arrRef.Refs()
	checkIndex(len(refs), index)
	refs[index] = val
}

type BASTORE struct {
	base.NoOperandsInstruction
}

func (self *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Bytes()
	checkIndex(len(ints), index)
	ints[index] = int8(val)
}

type CASTORE struct {
	base.NoOperandsInstruction
}

func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Chars()
	checkIndex(len(ints), index)
	ints[index] = uint16(val)
}

type DASTORE struct {
	base.NoOperandsInstruction
}

func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Doubles()
	checkIndex(len(ints), index)
	ints[index] = val
}

type FASTORE struct {
	base.NoOperandsInstruction
}

func (self *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Floats()
	checkIndex(len(ints), index)
	ints[index] = val
}

type IASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Ints()
	checkIndex(len(ints), index)
	ints[index] = val
}

type LASTORE struct {
	base.NoOperandsInstruction
}

func (self *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Longs()
	checkIndex(len(ints), index)
	ints[index] = val
}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	index := stack.PopInt()
	arrRef := frame.OperandStack().PopRef()

	checkNotNil(arrRef)
	ints := arrRef.Shorts()
	checkIndex(len(ints), index)
	ints[index] = int16(val)
}
