package filesystem

import "syscall"

func Link(oldname, newname string) error {
	e := syscall.Link(oldname, newname)
	if e != nil {
		return &LinkError{"link", oldname, newname, e}
	}
	return nil
}

func Symlink(oldname, newname string) error {
	e := syscall.Symlink(oldname, newname)
	if e != nil {
		return &LinkError{"symlink", oldname, newname, e}
	}
	return nil
}

func Lchown(name string, uid, gid int) error {
	if e := syscall.Lchown(name, uid, gid); e != nil {
		return &PathError{"lchown", name, e}
	}
	return nil
}
