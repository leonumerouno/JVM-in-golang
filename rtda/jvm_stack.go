package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0, self.size)
	for frame := self._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}

func (self *Stack) clear() {
	for !self.IsEmpty() {
		self.Pop()
	}
}

func (self *Stack) Push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

func (self *Stack) Pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	return top
}

func (self *Stack) Top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty")
	}
	return self._top
}

func (self *Stack) IsEmpty() bool {
	return self._top == nil
}
