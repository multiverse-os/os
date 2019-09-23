package filesystem

import "syscall"

type Dir struct {
	path  string
	files []File
}

func (f *File) Readdir(n int) ([]FileInfo, error) {
	if f == nil {
		return nil, ErrInvalid
	}
	return f.readdir(n)
}

func (f *File) Readdirnames(n int) (names []string, err error) {
	if f == nil {
		return nil, ErrInvalid
	}
	return f.readdirnames(n)
}

func (self *FileInfo) Directory() bool { return self.directory }

func Mkdir(name string, perm FileMode) error {
	e := syscall.Mkdir(fixLongPath(name), syscallMode(perm))
	if e != nil {
		return &PathError{"mkdir", name, e}
	}
	if !supportsCreateWithStickyBit && perm&ModeSticky != 0 {
		e = setStickyBit(name)
		if e != nil {
			Remove(name)
			return e
		}
	}
	return nil
}

func Chdir(dir string) error {
	if e := syscall.Chdir(dir); e != nil {
		testlog.Open(dir) // observe likely non-existent directory
		return &PathError{"chdir", dir, e}
	}
	if log := testlog.Logger(); log != nil {
		wd, err := Getwd()
		if err == nil {
			log.Chdir(wd)
		}
	}
	return nil
}
