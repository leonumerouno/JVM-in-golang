package rtda

import "jvm/rtda/heap"

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: NewStack(10240000),
	}
}

func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

func (self *Thread) ClearStack() {
	self.stack.clear()
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.Push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.Pop()
}

func (self *Thread) CurrentFrame() *Frame {
	return self.stack.Top()
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return NewFrame(self, method)
}

func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) IsStackEmpty() bool {
	return self.stack.IsEmpty()
}
