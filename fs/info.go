package fs

import "os"

func (self *FileInfo) Mode() os.FileMode {
	return self.mode
}

func (self *FileInfo) Name() string {
	return self.name
}

func (self *FileInfo) Size() int64 {
	return self.size
}

func (self *FileInfo) Sys() interface{} {
	return self.sys
}
