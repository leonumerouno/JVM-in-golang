package main

import (
	"fmt"
	"jvm/cmd"
)

func main() {
	c := cmd.Parsecmd()

	if c.VersionFlag {
		fmt.Println("Version 0.0.1")
	} else if c.HelpFlag || c.Class == "" {
		cmd.PrintUsage()
	} else {
		newJVM(c).start()
	}
}

//func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
//	classData, _, err := cp.ReadClass(className)
//	if err != nil {
//		panic(err)
//	}
//	cf, err := classfile.Parse(classData)
//	if err != nil {
//		panic(err)
//	}
//	return cf
//}

//func startJVM(c *cmd.Cmd) {
//	cp := classpath.Parse(c.XjreOption, c.CpOption)
//	classLoader := heap.NewClassLoader(cp, c.VerboseClassFlag)
//
//	className := strings.Replace(c.Class, ".", "/", -1)
//	mainClass := classLoader.LoadClass(className)
//	mainMethod := mainClass.GetMainMethod()
//	if mainMethod != nil {
//		interpret(mainMethod, c.VerboseClassFlag, c.Args)
//	} else {
//		fmt.Println("Main method not found in class %s\n", c.Class)
//	}
//}
//
//func getMainMethod(cf *classfile.ClassFile) *classfile.MemberInfo {
//	for _, m := range cf.Methods() {
//		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
//			return m
//		}
//	}
//	return nil
//}
