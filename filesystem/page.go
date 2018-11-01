package filesystem

import "syscall"

func Getpagesize() int { return syscall.Getpagesize() }
