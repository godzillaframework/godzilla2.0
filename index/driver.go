package index

import (
	"errors"
	"io"
	"io/fs"
)

var registry = make(map[string]func(config any) (Driver, error))

type Driver interface {
	Create(path string) (File, error)
	Open(path string) (File, error)
	Remove(path string) error
	Rename(oldPath, newPath string) error
	Stat(path string) (fs.FileInfo, error)
	MakeDir(path string) error
	RemoveDir(path string) error
	RenameDir(oldPath, newPath string) error
	StatDir(path string) (fs.FileInfo, error)
	ListDir(path string) ([]fs.DirEntry, error)
	ReadFile(path string) ([]byte, error)
	WriteFile(path string, data []byte) error
	ReadFileStream(path string) (io.ReadCloser, error)
	WriteFileStream(path string, data io.Reader) error
	RemoveFile(path string) error
	FileSize(path string) (int64, error)
	FileModTime(path string) (int64, error)
}

// Register a driver
func Register(name string, driverCreator func(config any) (Driver, error)) {
	registry[name] = driverCreator
}

// gettin a driver
func Get(name string, config any) (Driver, error) {
	driverCreator, ok := registry[name]

	if !ok {
		return nil, errors.New("driver not found")
	}

	return driverCreator(config)
}
