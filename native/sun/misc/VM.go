package misc

import (
	"jvm/instruction/base"
	"jvm/native"
	"jvm/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	//vmClass := frame.Method().Class()
	//savedProps := vmClass.GetRefVar("savedProps", "Ljava/util/Properties;")
	//key := heap.JString(vmClass.Loader(), "foo")
	//val := heap.JString(vmClass.Loader(), "bar")
	//
	//frame.OperandStack().PushRef(savedProps)
	//frame.OperandStack().PushRef(key)
	//frame.OperandStack().PushRef(val)
	//
	//propsClass := vmClass.Loader().LoadClass("java/util/Properties")
	//setPropMethod := propsClass.GetInstanceMethod("setProperty",
	//	"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	//base.InvokeMethod(frame, setPropMethod)
	classLoader := frame.Method().Class().Loader()
	jlSysClass := classLoader.LoadClass("java/lang/System")
	initSysClass := jlSysClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSysClass)
}
