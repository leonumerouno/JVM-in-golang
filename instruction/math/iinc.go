package math

import (
	"jvm/instruction/base"
	"jvm/rtda"
)

type IINC struct {
	Index uint
	Const int32
}

func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localvars := frame.LocalVars()
	val := localvars.GetInt(self.Index)
	val += self.Const
	localvars.SetInt(self.Index, val)
}
