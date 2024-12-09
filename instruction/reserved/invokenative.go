package reserved

import (
	"jvm/instruction/base"
	"jvm/native"
	_ "jvm/native/java/lang"
	_ "jvm/native/sun/misc"
	"jvm/rtda"
)

type INVOKENATIVE struct {
	base.NoOperandsInstruction
}

func (self *INVOKENATIVE) Execute(frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	methodDescriptor := method.Descriptor()

	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		methodInfo := className + "~" + methodName + "~" + methodDescriptor
		panic("java.lang.UnsatisifiedLinkError: " + methodInfo)
	}

	nativeMethod(frame)
}
