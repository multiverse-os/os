// +build aix darwin dragonfly freebsd linux nacl netbsd openbsd solaris

package linux

import "syscall"

// ParseDirent parses up to max directory entries in buf,
// appending the names to names. It returns the number of
// bytes consumed from buf, the number of entries added
// to names, and the new names slice.
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string) {
	return syscall.ParseDirent(buf, max, names)
}
