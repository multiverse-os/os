package fs

import "syscall"

func Getpagesize() int { return syscall.Getpagesize() }
