package lang

import (
	"jvm/native"
	"jvm/rtda"
	"jvm/rtda/heap"
)

func init() {
	native.Register("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

func checkArrayCopy(src, dst *heap.Object) bool {
	srcClass := src.Class()
	dstClass := dst.Class()
	if !srcClass.IsArray() || !dstClass.IsArray() {
		return false
	}
	if srcClass.ComponentClass().IsPrimitive() || dstClass.ComponentClass().IsPrimitive() {
		return srcClass == dstClass
	}
	return true
}

func arraycopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dst := vars.GetRef(2)
	dstPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dst == nil {
		panic("java.lang.NullPointerException")
	}

	if !checkArrayCopy(src, dst) {
		panic("java.lang.ArrayStoreException")
	}

	if srcPos < 0 || dstPos < 0 || length < 0 || srcPos+length > src.ArrayLength() || dstPos+length > dst.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, dst, srcPos, dstPos, length)
}