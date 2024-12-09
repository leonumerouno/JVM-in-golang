package main

import (
	"fmt"
	"jvm/instruction"
	"jvm/instruction/base"
	"jvm/rtda"
	"jvm/rtda/heap"
)

func interpret(thread *rtda.Thread, logInst bool) {
	defer catchErr(thread)
	loop(thread, logInst)
}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := loader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		//fmt.Println("LocalVars:%v\n", frame.LocalVars())
		//fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &base.BytecodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instruction.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		inst.Execute(frame)
		if thread.IsStackEmpty() {
			return
		}
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%.2d %T %v\n", className, methodName, pc, inst, inst)
}

//func loop(thread *rtda.Thread, bytecode []byte) {
//	frame := thread.PopFrame()
//	reader := &base.BytecodeReader{}
//	for {
//		pc := frame.NextPC()
//		thread.SetPC(pc)
//
//		//decode
//		reader.Reset(bytecode, pc)
//		opcode := reader.ReadUint8()
//		inst := instruction.NewInstruction(opcode)
//		inst.FetchOperands(reader)
//		frame.SetNextPC(reader.PC())
//
//		//execute
//		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
//		inst.Execute(frame)
//	}
//}
