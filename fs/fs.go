package fs

import (
	"syscall"
	"time"
)

func Open(name string) (*File, error) {
	return OpenFile(name, O_RDONLY, 0)
}

func Create(name string) (*File, error) {
	return OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
}

func Rename(oldpath, newpath string) error {
	return rename(oldpath, newpath)
}

func Remove(name string) error {
	// System call interface forces us to know
	// whether name is a file or directory.
	// Try both: it is cheaper on average than
	// doing a Stat plus the right one.
	e := syscall.Unlink(name)
	if e == nil {
		return nil
	}
	e1 := syscall.Rmdir(name)
	if e1 == nil {
		return nil
	}

	// Both failed: figure out which error to return.
	// OS X and Linux differ on whether unlink(dir)
	// returns EISDIR, so can't use that. However,
	// both agree that rmdir(file) returns ENOTDIR,
	// so we can use that to decide which error is real.
	// Rmdir might also return ENOTDIR if given a bad
	// file path, like /etc/passwd/foo, but in that case,
	// both errors will be ENOTDIR, so it's okay to
	// use the error from unlink.
	if e1 != syscall.ENOTDIR {
		e = e1
	}
	return &PathError{"remove", name, e}
}

func (self *FileInfo) ModifiedAt() time.Time {
	return self.modifiedAt
}
