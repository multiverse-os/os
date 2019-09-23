// +build !plan9
package net

import (
	"runtime"
)

// rawConn implements syscall.RawConn.
type rawConn struct {
	file *File
}

func (c *rawConn) Control(f func(uintptr)) error {
	if err := c.file.checkValid("SyscallConn.Control"); err != nil {
		return err
	}
	err := c.file.pfd.RawControl(f)
	runtime.KeepAlive(c.file)
	return err
}

func (c *rawConn) Read(f func(uintptr) bool) error {
	if err := c.file.checkValid("SyscallConn.Read"); err != nil {
		return err
	}
	err := c.file.pfd.RawRead(f)
	runtime.KeepAlive(c.file)
	return err
}

func (c *rawConn) Write(f func(uintptr) bool) error {
	if err := c.file.checkValid("SyscallConn.Write"); err != nil {
		return err
	}
	err := c.file.pfd.RawWrite(f)
	runtime.KeepAlive(c.file)
	return err
}

func newRawConn(file *File) (*rawConn, error) {
	return &rawConn{file: file}, nil
}
