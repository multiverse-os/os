package filesystem

// TODO: Since this library is meant to completely replace the Go 'os' stdlib we
// we will implement all calls to os and reduce our overall footprint by only
// supporting linux, android and unix
import (
	"os"
	"syscall"
	"time"
)

type File struct {
	name        string
	size        int64
	mode        os.FileMode
	modifiedAt  time.Time
	sys         interface{} // What is this?
	pfd         poll.FD
	dirinfo     *dirInfo // nil unless directory being read
	stdoutOrErr bool     // whether this is stdout or stderr
	nonblock    bool     // whether we set nonblocking mode
}

func (f *File) Fd() uintptr {
	if f == nil {
		return ^(uintptr(0))
	}

	// If we put the file descriptor into nonblocking mode,
	// then set it to blocking mode before we return it,
	// because historically we have always returned a descriptor
	// opened in blocking mode. The File will continue to work,
	// but any blocking operation will tie up a thread.
	if f.nonblock {
		f.pfd.SetBlocking()
	}

	return uintptr(f.pfd.Sysfd)
}

func (f *File) Truncate(size int64) error {
	if err := f.checkValid("truncate"); err != nil {
		return err
	}
	if e := f.pfd.Ftruncate(size); e != nil {
		return f.wrapErr("truncate", e)
	}
	return nil
}

func (f *File) Chown(uid, gid int) error {
	if err := f.checkValid("chown"); err != nil {
		return err
	}
	if e := f.pfd.Fchown(uid, gid); e != nil {
		return f.wrapErr("chown", e)
	}
	return nil
}

func (f *File) Chmod(mode FileMode) error { return f.chmod(mode) }

func (f *File) SetDeadline(t time.Time) error {
	return f.setDeadline(t)
}

func (f *File) SetReadDeadline(t time.Time) error {
	return f.setReadDeadline(t)
}

func (f *File) SetWriteDeadline(t time.Time) error {
	return f.setWriteDeadline(t)
}

func (f *File) Sync() error {
	if err := f.checkValid("sync"); err != nil {
		return err
	}
	if e := f.pfd.Fsync(); e != nil {
		return f.wrapErr("sync", e)
	}
	return nil
}

func Chtimes(name string, atime time.Time, mtime time.Time) error {
	var utimes [2]syscall.Timespec
	utimes[0] = syscall.NsecToTimespec(atime.UnixNano())
	utimes[1] = syscall.NsecToTimespec(mtime.UnixNano())
	if e := syscall.UtimesNano(fixLongPath(name), utimes[0:]); e != nil {
		return &PathError{"chtimes", name, e}
	}
	return nil
}

// Chdir changes the current working directory to the file,
// which must be a directory.
// If there is an error, it will be of type *PathError.
func (f *File) Chdir() error {
	if err := f.checkValid("chdir"); err != nil {
		return err
	}
	if e := f.pfd.Fchdir(); e != nil {
		return f.wrapErr("chdir", e)
	}
	return nil
}

// setDeadline sets the read and write deadline.
func (f *File) setDeadline(t time.Time) error {
	if err := f.checkValid("SetDeadline"); err != nil {
		return err
	}
	return f.pfd.SetDeadline(t)
}

// setReadDeadline sets the read deadline.
func (f *File) setReadDeadline(t time.Time) error {
	if err := f.checkValid("SetReadDeadline"); err != nil {
		return err
	}
	return f.pfd.SetReadDeadline(t)
}

// setWriteDeadline sets the write deadline.
func (f *File) setWriteDeadline(t time.Time) error {
	if err := f.checkValid("SetWriteDeadline"); err != nil {
		return err
	}
	return f.pfd.SetWriteDeadline(t)
}

// checkValid checks whether f is valid for use.
// If not, it returns an appropriate error, perhaps incorporating the operation name op.
func (f *File) checkValid(op string) error {
	if f == nil {
		return ErrInvalid
	}
	return nil
}
