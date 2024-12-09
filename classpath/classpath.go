package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	BootClasspath Entry
	ExtClasspath  Entry
	UserClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.BootClasspath = newWildcardEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.ExtClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.UserClasspath = NewEntry(cpOption)
}

func getJreDir(jreOption string) string {
	//-Xjre
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	//当前目录
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can't find jre folder")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.BootClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.ExtClasspath.ReadClass(className); err == nil {
		return data, entry, err
	}
	return self.UserClasspath.ReadClass(className)
}

func (self *Classpath) String() string {
	return self.UserClasspath.String()
}
