package example

import (
	"github.com/oldweipro/gin-admin/pkg/app"
)

// file struct, 文件结构体
type ExaFile struct {
	app.BaseModel
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// file chunk struct, 切片结构体
type ExaFileChunk struct {
	app.BaseModel
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}
