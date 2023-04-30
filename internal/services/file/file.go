package file

import (
	"os"
	"sync"
)

type File struct {
	mutex    sync.RWMutex
	filePath string
	perm     os.FileMode
}

func New(name string, perm os.FileMode) *File {
	return &File{filePath: name, perm: perm}
}

func (f *File) Load() ([]byte, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return os.ReadFile(f.filePath)
}

func (f *File) Save(data []byte) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	return os.WriteFile(f.filePath, data, f.perm)
}
