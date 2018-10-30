package fs

import "syscall"

func Chmod(name string, mode FileMode) error { return chmod(name, mode) }

func Chown(name string, uid, gid int) error {
	if e := syscall.Chown(name, uid, gid); e != nil {
		return &PathError{"chown", name, e}
	}
	return nil
}
