// +build go1.9
package linux

import "syscall"

type Signal = syscall.Signal
type Errno = syscall.Errno
type SysProcAttr = syscall.SysProcAttr
