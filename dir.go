package os

import (
	"io"
	"runtime"
	"syscall"
)

// Auxiliary information if the File describes a directory
type dirInfo struct {
	buf  []byte // buffer for directory I/O
	nbuf int    // length of buf; return value from Getdirentries
	bufp int    // location of next record in buf.
}

const (
	// More than 5760 to work around https://golang.org/issue/24015.
	blockSize = 8192
)

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

func (d *dirInfo) close() {}

func (f *File) readdirnames(n int) (names []string, err error) {
	// If this file has no dirinfo, create one.
	if f.dirinfo == nil {
		f.dirinfo = new(dirInfo)
		// The buffer must be at least a block long.
		f.dirinfo.buf = make([]byte, blockSize)
	}
	d := f.dirinfo

	size := n
	if size <= 0 {
		size = 100
		n = -1
	}

	names = make([]string, 0, size) // Empty with room to grow.
	for n != 0 {
		// Refill the buffer if necessary
		if d.bufp >= d.nbuf {
			d.bufp = 0
			var errno error
			d.nbuf, errno = f.pfd.ReadDirent(d.buf)
			runtime.KeepAlive(f)
			if errno != nil {
				return names, wrapSyscallError("readdirent", errno)
			}
			if d.nbuf <= 0 {
				break // EOF
			}
		}

		// Drain the buffer
		var nb, nc int
		nb, nc, names = syscall.ParseDirent(d.buf[d.bufp:d.nbuf], n, names)
		d.bufp += nb
		n -= nc
	}
	if n >= 0 && len(names) == 0 {
		return names, io.EOF
	}
	return names, nil
}
