package main

import (
	"fmt"
	"jvm/classpath"
	"jvm/cmd"
	"jvm/instruction/base"
	"jvm/rtda"
	"jvm/rtda/heap"
	"strings"
)

type JVM struct {
	cmd         *cmd.Cmd
	classLoader *heap.ClassLoader
	mainThread  *rtda.Thread
}

func newJVM(c *cmd.Cmd) *JVM {
	cp := classpath.Parse(c.XjreOption, c.CpOption)
	classLoader := heap.NewClassLoader(cp, c.VerboseClassFlag)

	return &JVM{
		cmd:         c,
		classLoader: classLoader,
		mainThread:  rtda.NewThread(),
	}
}

func (self *JVM) start() {
	self.initVM()
	self.execMain()
}

func (self *JVM) initVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.VerboseClassFlag)
}

func (self *JVM) execMain() {
	className := strings.Replace(self.cmd.Class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Println("Main method not found in class %s\n", self.cmd.Class)
		return
	}

	argsArr := self.CreateArgsArray()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.LocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.VerboseInstFlag)
}

func (self *JVM) CreateArgsArray() *heap.Object {
	stringClass := self.classLoader.LoadClass("java/lang/String")
	argsArr := stringClass.ArrayClass().NewArray(uint(len(self.cmd.Args)))
	jArgs := argsArr.Refs()
	for i, arg := range self.cmd.Args {
		jArgs[i] = heap.JString(self.classLoader, arg)
	}
	return argsArr
}
