package watch

import (
	"fmt"
	"os"

	inotify "github.com/multiverse-os/os/filesystem/watch/inotify"
)

type FileDescriptor struct {
	Path            string
	Mask            int
	WatchDescriptor int
	Running         bool
}

func NewFileDescriptor(path string, mask int) *FileDescriptor {
	return &WatchDescriptor{
		Path:            path,
		WatchDescriptor: -1,
		Mask:            mask,
	}
}

func (self *FileDescriptor) PathExists() bool {
	_, err := os.Lstat(self.Path)
	return os.IsExist(err)
}

func (self *FileDescriptor) Start(fd int) (err error) {
	if self.Running {
		return ErrDescAlreadyRunning
	}
	self.FileDescriptor, err = inotify.AddWatch(fd, self.Path, uint32(self.Mask))
	if self.FileDescriptor == -1 || err != nil {
		self.Running = false
		return fmt.Errorf("%s: %s", ErrDescNotStart, err)
	}
	return nil
}

func (self *FileDescriptor) Stop(fd int) (err error) {
	if !self.Running {
		return ErrDescNotRunning
	}
	_, err = inotify.RemoveWatch(fd, uint32(self.FileDescriptor))
	if err != nil {
		return fmt.Errorf("%s: %s", ErrDescNotStopped, err)
	}
	self.Running = false
	return nil
}
