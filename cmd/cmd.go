package cmd

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	HelpFlag         bool
	VersionFlag      bool
	VerboseClassFlag bool
	VerboseInstFlag  bool
	CpOption         string
	Class            string
	Args             []string
	XjreOption       string
}

func Parsecmd() *Cmd {
	c := &Cmd{}

	flag.Usage = PrintUsage
	flag.BoolVar(&c.HelpFlag, "help", false, "print help message")
	flag.BoolVar(&c.HelpFlag, "?", false, "print help message")
	flag.BoolVar(&c.VersionFlag, "version", false, "print version and exit")
	flag.StringVar(&c.CpOption, "classpath", "", "classpath")
	flag.StringVar(&c.CpOption, "cp", "", "classpath")
	flag.StringVar(&c.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		c.Class = args[0]
		c.Args = args[1:]
	}

	return c
}

func PrintUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
