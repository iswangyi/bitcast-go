package data

type LogRecordPos struct {
	Offset int64  // offset,表示文件中的偏移量
	Fid    uint32 // file id,表示文件的编号
}
