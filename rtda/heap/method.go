package heap

import "jvm/classfile"

type Method struct {
	ClassMember
	maxStack        uint
	maxLocals       uint
	code            []byte
	argSlotCount    uint
	exceptionTable  ExceptionTable
	lineNumberTable *classfile.LineNumberTableAttribute
}

func NewMethods(class *Class, cfMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(cfMethods))
	for i, cfMethod := range cfMethods {
		methods[i] = newMethod(class, cfMethod)
	}
	return methods
}

func newMethod(class *Class, cfMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(cfMethod)
	method.copyAttributes(cfMethod)
	md := parseMethodDescriptor(method.descriptor)
	method.calcArgSlotCount(md.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttribute(md.returnType)
	}
	return method
}

func (self *Method) injectCodeAttribute(returnType TypeDescriptor) {
	self.maxStack = 4
	self.maxLocals = self.ArgSlotCount()
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xfe, 0xb1}
	case 'D':
		self.code = []byte{0xfe, 0xaf}
	case 'F':
		self.code = []byte{0xfe, 0xae}
	case 'J':
		self.code = []byte{0xfe, 0xad}
	case 'L', '[':
		self.code = []byte{0xfe, 0xb0}
	default:
		self.code = []byte{0xfe, 0xac}
	}
}

func (self *Method) copyAttributes(cfMethod *classfile.MemberInfo) {
	if codeAttr := cfMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = uint(codeAttr.MaxStack())
		self.maxLocals = uint(codeAttr.MaxLocals())
		self.code = codeAttr.Code()
		self.exceptionTable = newExceptionTable(codeAttr.ExceptionTable(), self.class.constantPool)
		self.lineNumberTable = codeAttr.LineNumberTableAttribute()
	}
}

func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable == nil {
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}

func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (self *Method) calcArgSlotCount(paramTypes []TypeDescriptor) {
	self.argSlotCount = 0
	for _, paramType := range paramTypes {
		self.argSlotCount++
		if paramType == "J" || paramType == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

func (self *Method) Class() *Class {
	return self.class
}

func (self *Method) Name() string {
	return self.name
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) Code() []byte {
	return self.code
}

func (self *Method) ArgSlotCount() uint {
	return self.argSlotCount
}

func (self *Method) Descriptor() string {
	return self.descriptor
}
