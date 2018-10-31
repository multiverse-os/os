package watch

import "unsafe"

// Inotify based watch implementation
const SizeofInotifyEvent = 0x10

type Event struct {
	Name             [0]int8
	workingDirectory int32
	Mask             uint32
	Cookie           uint32
	Length           uint32
}

func InotifyInit() (fd int, err error) {
	return InotifyInit1(0)
}

func InotifyInit1(flags int) (fd int, err error) {
	r0, _, e1 := RawSyscall(SYS_INOTIFY_INIT1, uintptr(flags), 0, 0)
	fd = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}

func InotifyRmWatch(fd int, watchdesc uint32) (success int, err error) {
	r0, _, e1 := RawSyscall(SYS_INOTIFY_RM_WATCH, uintptr(fd), uintptr(watchdesc), 0)
	success = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}

	return
}

func InotifyAddWatch(fd int, pathname string, mask uint32) (watchdesc int, err error) {
	var _p0 *byte
	_p0, err = BytePtrFromString(pathname)
	if err != nil {
		return
	}
	r0, _, e1 := Syscall(SYS_INOTIFY_ADD_WATCH, uintptr(fd), uintptr(unsafe.Pointer(_p0)), uintptr(mask))
	watchdesc = int(r0)
	if e1 != 0 {
		err = errnoErr(e1)
	}
	return
}
