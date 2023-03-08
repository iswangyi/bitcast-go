package fio

import "os"

type FileIo struct {
	fd *os.File
}

// NewFileIo returns a new FileIo
func NewFileIo(path string) (*FileIo, error) {
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}
	return &FileIo{fd: fd}, nil
}

func (f *FileIo) Read(b []byte, offset int64) (int, error) {
	return f.fd.ReadAt(b, offset)
}

func (f *FileIo) Write(b []byte) (int, error) {
	return f.Write(b)
}

func (f *FileIo) Sync() error {
	return f.fd.Sync()
}

func (f *FileIo) Close() error {
	return f.fd.Close()
}
