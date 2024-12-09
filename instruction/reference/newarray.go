package reference

import (
	"jvm/instruction/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

type NEWARRAY struct {
	atype uint8
}

const (
	//Array Type  atype
	ATBoolean = 4
	ATChar    = 5
	ATFloat   = 6
	ATDouble  = 7
	ATByte    = 8
	ATShort   = 9
	ATInt     = 10
	ATLong    = 11
)

func (self *NEWARRAY) FetchOperands(reader *base.BytecodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEWARRAY) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	classLoader := frame.Method().Class().Loader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case ATBoolean:
		return loader.LoadClass("[Z")
	case ATByte:
		return loader.LoadClass("[B")
	case ATChar:
		return loader.LoadClass("[C")
	case ATShort:
		return loader.LoadClass("[S")
	case ATInt:
		return loader.LoadClass("[I")
	case ATLong:
		return loader.LoadClass("[J")
	case ATFloat:
		return loader.LoadClass("[F")
	case ATDouble:
		return loader.LoadClass("[D")
	default:
		panic("invalid atype!")
	}
}
