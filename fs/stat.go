package fs

import (
	"internal/testlog"
	"syscall"
	"time"
)

func (fs *fileStat) Name() string { return fs.name }
func (fs *fileStat) IsDir() bool  { return fs.Mode().IsDir() }

// Stat returns a FileInfo describing the named file.
// If there is an error, it will be of type *PathError.
func Stat(name string) (FileInfo, error) {
	testlog.Stat(name)
	return statNolog(name)
}

// Lstat returns a FileInfo describing the named file.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link. Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
func Lstat(name string) (FileInfo, error) {
	testlog.Stat(name)
	return lstatNolog(name)
}

func fillFileStatFromSys(fs *fileStat, name string) {
	fs.name = basename(name)
	fs.size = fs.sys.Size
	fs.modTime = timespecToTime(fs.sys.Mtim)
	fs.mode = FileMode(fs.sys.Mode & 0777)
	switch fs.sys.Mode & syscall.S_IFMT {
	case syscall.S_IFBLK:
		fs.mode |= ModeDevice
	case syscall.S_IFCHR:
		fs.mode |= ModeDevice | ModeCharDevice
	case syscall.S_IFDIR:
		fs.mode |= ModeDir
	case syscall.S_IFIFO:
		fs.mode |= ModeNamedPipe
	case syscall.S_IFLNK:
		fs.mode |= ModeSymlink
	case syscall.S_IFREG:
		// nothing to do
	case syscall.S_IFSOCK:
		fs.mode |= ModeSocket
	}
	if fs.sys.Mode&syscall.S_ISGID != 0 {
		fs.mode |= ModeSetgid
	}
	if fs.sys.Mode&syscall.S_ISUID != 0 {
		fs.mode |= ModeSetuid
	}
	if fs.sys.Mode&syscall.S_ISVTX != 0 {
		fs.mode |= ModeSticky
	}
}

func timespecToTime(ts syscall.Timespec) time.Time {
	return time.Unix(int64(ts.Sec), int64(ts.Nsec))
}

// For testing.
func atime(fi FileInfo) time.Time {
	return timespecToTime(fi.Sys().(*syscall.Stat_t).Atim)
}

// Stat returns the FileInfo structure describing file.
// If there is an error, it will be of type *PathError.
func (f *File) Stat() (FileInfo, error) {
	if f == nil {
		return nil, ErrInvalid
	}
	var fs fileStat
	err := f.pfd.Fstat(&fs.sys)
	if err != nil {
		return nil, &PathError{"stat", f.name, err}
	}
	fillFileStatFromSys(&fs, f.name)
	return &fs, nil
}

// statNolog stats a file with no test logging.
func statNolog(name string) (FileInfo, error) {
	var fs fileStat
	err := syscall.Stat(name, &fs.sys)
	if err != nil {
		return nil, &PathError{"stat", name, err}
	}
	fillFileStatFromSys(&fs, name)
	return &fs, nil
}

// lstatNolog lstats a file with no test logging.
func lstatNolog(name string) (FileInfo, error) {
	var fs fileStat
	err := syscall.Lstat(name, &fs.sys)
	if err != nil {
		return nil, &PathError{"lstat", name, err}
	}
	fillFileStatFromSys(&fs, name)
	return &fs, nil
}
