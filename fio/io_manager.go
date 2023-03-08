package fio

type IOManager interface {
	// Read ReadAt reads len(b) bytes from the File starting at byte offset off.
	Read(b []byte, offset int64) (int, error)
	// Write WriteAt writes len(b) bytes
	Write(b []byte) (int, error)
	// Sync commits the current contents of the file to stable storage.
	Sync() error
	// Close closes the File, rendering it unusable for I/O.
	Close() error
}
