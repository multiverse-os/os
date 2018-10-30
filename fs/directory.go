package fs

import "syscall"

// We will store this in a prefix-sorting radix tree to have very fast
// autocomplete

type Dir struct {
	path  string
	files []File
}

func (self *FileInfo) Directory() bool {
	return self.directory
}

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

// Chdir changes the current working directory to the named directory.
// If there is an error, it will be of type *PathError.
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
