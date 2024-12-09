package classpath

import (
	"os"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir: absDir}
}

func (self *DirEntry) ReadClass(ClassName string) ([]byte, Entry, error) {
	filename := filepath.Join(self.absDir, ClassName)
	data, err := os.ReadFile(filename)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
