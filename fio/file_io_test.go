package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func destroyFile(f *FileIo) {
	if err := os.Remove(f.fd.Name()); err != nil {
		panic(err)
	}
	f.Close()
}

func TestNewFileIo(t *testing.T) {

	f, err := NewFileIo("./0000.data")
	//defer destroyFile(f)
	if err != nil {
		t.Error(err)
	}
	n, err := f.Write([]byte("hello"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 5, n)
	n, err = f.Write([]byte(" world\n"))
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 7, n)

}

func TestFileIo_Sync(t *testing.T) {
	f, err := NewFileIo("./0000.data")
	defer destroyFile(f)
	assert.Equal(t, nil, err)

	nw, err := f.Write([]byte("hello\n"))
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, nw)

	err = f.Sync()
	assert.Equal(t, nil, err)

}

func TestFileIo_Read(t *testing.T) {
	f, err := NewFileIo("./0000.data")
	defer destroyFile(f)
	assert.Equal(t, nil, err)

	f.Write([]byte("hello\n"))

	b := make([]byte, 6)
	n, err := f.Read(b, 0)
	assert.Equal(t, nil, err)
	assert.Equal(t, 6, n)
	assert.Equal(t, "hello\n", string(b[:n]))

}

func TestFileIo_Close(t *testing.T) {
	f, err := NewFileIo("./0000.data")
	defer destroyFile(f)
	assert.Equal(t, nil, err)

	err = f.Close()
	assert.Equal(t, nil, err)
}
