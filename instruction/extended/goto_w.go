package extended

import (
	"jvm/instruction/base"
	"jvm/rtda"
)

type GOTOW struct {
	offset int
}

func (self *GOTOW) FetchOperands(reader *base.BytecodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTOW) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}
