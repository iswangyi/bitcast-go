package fio

import (
	"fmt"
	"os"
)

type FileIo struct {
	fd *os.File
}

// NewFileIo returns a new FileIo
func NewFileIo(path string) (*FileIo, error) {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0640)
	if err != nil {
		return nil, err
	}
	return &FileIo{fd: fd}, nil
}

func (f *FileIo) Read(b []byte, offset int64) (int, error) {
	n, err := f.fd.ReadAt(b, offset)
	if err != nil {
		return 0, err
	}
	if n != len(b) {
		return 0, fmt.Errorf("read %d bytes, but expected %d bytes", n, len(b))
	}
	return n, nil
}

// Write writes len(b) bytes to the File.
// It returns the number of bytes written and an error, if any.
// Write returns a non-nil error when n != len(b).
func (f *FileIo) Write(b []byte) (int, error) {
	n, err := f.fd.Write(b)
	if err != nil {
		return 0, err
	}
	if n != len(b) {
		return 0, fmt.Errorf("write %d bytes, but expected %d bytes", n, len(b))
	}
	return n, nil
}

// Sync commits the current contents of the file to stable storage.
// Typically, this means flushing the file system's in-memory copy
// of recently written data to disk.
func (f *FileIo) Sync() error {
	return f.fd.Sync()
}

func (f *FileIo) Close() error {
	return f.fd.Close()
}
