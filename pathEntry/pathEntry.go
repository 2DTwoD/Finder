package pathEntry

import (
	"os"
)

type DirEntryWithPath struct {
	dirEntry *os.DirEntry
	path     string
}

func New(entry *os.DirEntry, path string) *DirEntryWithPath {
	return &DirEntryWithPath{entry, path}
}

func (f *DirEntryWithPath) PathWithName() string {
	return f.path + (*f.dirEntry).Name()
}

func (f *DirEntryWithPath) Path() string {
	return f.path
}

func (f *DirEntryWithPath) AppendPath() {
	f.path += f.Name() + "/"
}

func (f *DirEntryWithPath) Name() string {
	return (*f.dirEntry).Name()
}

func (f *DirEntryWithPath) IsDir() bool {
	return (*f.dirEntry).IsDir()
}
