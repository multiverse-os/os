package watch

// TODO: Since this library is meant to completely replace the Go 'os' stdlib we
// we will implement all calls to os and reduce our overall footprint by only
// supporting linux, android and unix
import (
	"os"
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
