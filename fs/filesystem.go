package filesystem

import (
	"syscall"
	"time"
)

func Open(name string) (*File, error)        { return OpenFile(name, O_RDONLY, 0) }
func Create(name string) (*File, error)      { return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666) }
func Rename(oldpath, newpath string) error   { return rename(oldpath, newpath) }
func (self *FileInfo) ModifiedAt() time.Time { return self.modifiedAt }

func Remove(name string) error {
	e := syscall.Unlink(name)
	if e == nil {
		return nil
	}
	e1 := syscall.Rmdir(name)
	if e1 == nil {
		return nil
	}
	if e1 != syscall.ENOTDIR {
		e = e1
	}
	return &PathError{"remove", name, e}
}
