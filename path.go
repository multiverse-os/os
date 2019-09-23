package os

import (
	"syscall"
)

const (
	PathSeparator     = '/' // OS-specific path separator
	PathListSeparator = ':' // OS-specific path list separator
)

func MkdirAll(path string, perm FileMode) error {
	dir, err := Stat(path)
	if err == nil {
		if dir.IsDir() {
			return nil
		}
		return &PathError{"mkdir", path, syscall.ENOTDIR}
	}
	i := len(path)
	for i > 0 && IsPathSeparator(path[i-1]) { // Skip trailing path separator.
		i--
	}

	j := i
	for j > 0 && !IsPathSeparator(path[j-1]) { // Scan backward over element.
		j--
	}

	if j > 1 {
		err = MkdirAll(fixRootDirectory(path[:j-1]), perm)
		if err != nil {
			return err
		}
	}

	err = Mkdir(path, perm)
	if err != nil {
		dir, err1 := Lstat(path)
		if err1 == nil && dir.IsDir() {
			return nil
		}
		return err
	}
	return nil
}

var removeAllTestHook = func(err error) error { return err }

func RemoveAll(path string) error { return removeAll(path) }

// endsWithDot reports whether the final component of path is ".".
func endsWithDot(path string) bool {
	if path == "." {
		return true
	}
	if len(path) >= 2 && path[len(path)-1] == '.' && IsPathSeparator(path[len(path)-2]) {
		return true
	}
	return false
}

func IsPathSeparator(c uint8) bool {
	return PathSeparator == c
}

func basename(name string) string {
	i := len(name) - 1
	for ; i > 0 && name[i] == '/'; i-- {
		name = name[:i]
	}
	for i--; i >= 0; i-- {
		if name[i] == '/' {
			name = name[i+1:]
			break
		}
	}
	return name
}

func splitPath(path string) (string, string) {
	dirname := "."
	for len(path) > 1 && path[0] == '/' && path[1] == '/' {
		path = path[1:]
	}
	i := len(path) - 1
	for ; i > 0 && path[i] == '/'; i-- {
		path = path[:i]
	}
	basename := path
	for i--; i >= 0; i-- {
		if path[i] == '/' {
			if i == 0 {
				dirname = path[:1]
			} else {
				dirname = path[:i]
			}
			basename = path[i+1:]
			break
		}
	}

	return dirname, basename
}

func fixRootDirectory(p string) string { return p }
