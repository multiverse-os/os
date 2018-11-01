// For Linux, get the pagesize from the runtime.
package linux

import "syscall"

func Getpagesize() int {
	return syscall.Getpagesize()
}
