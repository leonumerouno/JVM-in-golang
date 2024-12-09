package classpath

import (
	"os"
	"strings"
)

const pathListSeperator = string(os.PathListSeparator)

type Entry interface {
	ReadClass(ClassName string) ([]byte, Entry, error)
	String() string
}

func NewEntry(path string) Entry {
	if strings.Contains(path, pathListSeperator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
