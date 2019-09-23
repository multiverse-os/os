package os

import (
	"runtime"
	"syscall"
)

var Args []string

func init()                  { Args = runtime_args() }
func runtime_args() []string // in package runtime
func Getuid() int            { return syscall.Getuid() }
func Geteuid() int           { return syscall.Geteuid() }
func Getgid() int            { return syscall.Getgid() }
func Getegid() int           { return syscall.Getegid() }

func Getgroups() ([]int, error) {
	gids, e := syscall.Getgroups()
	return gids, NewSyscallError("getgroups", e)
}

func Exit(code int) {
	if code == 0 {
		runtime_beforeExit()
	}
	syscall.Exit(code)
}

func runtime_beforeExit()
