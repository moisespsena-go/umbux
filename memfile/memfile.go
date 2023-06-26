package memfile

import (
	"bytes"
	"io/fs"
	"os"
	"time"
)

type FileInfo struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	isDir   bool
	sys     any
}

func NewFileInfo(name string, size int64, mode fs.FileMode, modTime time.Time, isDir bool, sys any) *FileInfo {
	return &FileInfo{name: name, size: size, mode: mode, modTime: modTime, isDir: isDir, sys: sys}
}

func (i *FileInfo) Name() string {
	return i.name
}

func (i *FileInfo) Size() int64 {
	return i.size
}

func (i *FileInfo) Mode() fs.FileMode {
	return i.mode
}

func (i *FileInfo) ModTime() time.Time {
	return i.modTime
}

func (i *FileInfo) IsDir() bool {
	return i.isDir
}

func (i *FileInfo) Sys() any {
	return i.sys
}

type MemFile struct {
	Data *bytes.Reader
	Info os.FileInfo
}

func (f *MemFile) Stat() (fs.FileInfo, error) {
	return f.Info, nil
}

func (f *MemFile) Read(bytes []byte) (int, error) {
	return f.Data.Read(bytes)
}

func (f *MemFile) Close() (err error) {
	_, err = f.Data.Seek(0, 0)
	return
}
