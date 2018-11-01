// +build aix darwin dragonfly freebsd linux netbsd openbsd solaris

package linux

import "runtime"

// IoctlSetWinsize performs an ioctl on fd with a *Winsize argument.
//
// To change fd's window size, the req argument should be TIOCSWINSZ.
func IoctlSetWinsize(fd int, req uint, value *Winsize) error {
	// TODO: if we get the chance, remove the req parameter and
	// hardcode TIOCSWINSZ.
	err := ioctlSetWinsize(fd, req, value)
	runtime.KeepAlive(value)
	return err
}

// IoctlSetTermios performs an ioctl on fd with a *Termios.
//
// The req value will usually be TCSETA or TIOCSETA.
func IoctlSetTermios(fd int, req uint, value *Termios) error {
	// TODO: if we get the chance, remove the req parameter.
	err := ioctlSetTermios(fd, req, value)
	runtime.KeepAlive(value)
	return err
}