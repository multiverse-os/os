package watch

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"unsafe"

	inotify "github.com/multiverse-os/os/filesystem/watch/inotify"
	linux "github.com/multiverse-os/os/linux"
)

type Options struct {
	Recursive       bool
	UseWatcherFlags bool
}

// TODO: Use the list of files being watched as the running bool, if its not
// running, empty the list, if the list has items, then it is running
type Watcher struct {
	Mutex           sync.Mutex
	Path            string
	WatchDescriptor int
	DefaultMask     int
	FileDescriptors map[string]*FileDescriptor
	Events          chan *Event
	Errors          chan error
	Options         *Options
}

func New(path string, mask int, options *Options) (*Watcher, error) {
	fd, err := inotify.Init()
	if fd == -1 || err != nil {
		return nil, fmt.Errorf("%s: %s", ErrWatchNotCreated, err)
	}
	w := &Watcher{
		Path:            CleanPath(path),
		FileDescriptor:  fd,
		DefaultMask:     mask,
		FileDescriptors: make(map[string]*FileDescriptor),
		Events:          make(chan *Event),
		Errors:          make(chan error),
		Options:         options,
	}
	if options.Recursive == true {
		if err := w.AddFileDescriptor(w.Path, mask); err != nil {
			return w, err
		}
		return w, w.RecursiveAdd(w.Path, mask)
	}
	return w, w.AddFileDescriptor(w.Path, mask)
}

func (self *Watcher) FileDescriptorExists(path string) bool {
	self.Lock()
	defer self.Unlock()
	if _, exists := self.FileDescriptors[path]; exists {
		return true
	}
	return false
}

func (self *Watcher) FileDescriptorList() []string {
	list := make([]string, len(self.Descriptors))
	self.Lock()
	defer self.Unlock()
	for path, _ := range self.FileDescriptors {
		list = append(list, path)
	}
	return list
}

func (self *Watcher) RemoveFileDescriptor(path string) error {
	if !self.DescriptorExists(path) {
		return ErrDescNotFound
	}
	self.Lock()
	defer self.Unlock()
	fd := self.FileDescriptors[path]
	if fd.PathExists() {
		if err := fd.Stop(fd.FileDescriptor); err != nil {
			return err
		}
	}
	delete(self.FileDescriptors, path)
	return nil
}

func (self *Watcher) AddFileDescriptor(path string, mask int) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("%s: %s", ErrDescNotCreated, "directory does not exist")
	}
	if self.DescriptorExists(path) {
		return ErrDescNotFound
	}
	var inotifymask int
	if self.Options.UseWatcherFlags {
		inotifymask = self.DefaultMask
	} else {
		inotifymask = mask
	}
	self.Lock()
	defer self.Unlock()
	self.FileDescriptors[path] = NewFileDescriptor(path, inotifymask)
	return nil
}

func (self *Watcher) Add(rootPath string, mask int) error {
	directory, err := ioutil.ReadDir(rootPath)
	if err != nil {
		return err
	}
	var inotifymask int
	if self.Options.UseWatcherFlags {
		inotifymask = self.DefaultMask
	} else {
		inotifymask = mask
	}
	for _, child := range directory {
		if child.IsDirectory() {
			childPath := CleanPath(JoinPath(rootPath, child.Name()))
			if err := self.RecursiveAdd(childPath, inotifymask); err != nil {
				return err
			}
			if err := w.AddDescriptor(childPath, inotifymask); err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *Watcher) StartAll() error {
	self.Lock()
	defer self.Unlock()
	for _, fd := range self.FileDescriptors {
		if err := fd.Start(self.FileDescriptor); err != nil {
			return err
		}
	}
	return nil
}

func (self *Watcher) StopAll() error {
	self.Lock()
	defer self.Unlock()
	for _, fd := range self.FileDescriptors {
		if fd.Running {
			if err := fd.Stop(self.FileDescriptor); err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *Watcher) FileDescriptorWithWatchDescriptor(watchDescriptor int) *FileDescriptor {
	self.Lock()
	defer self.Unlock()
	for _, fd := range self.FileDescriptors {
		if self.WatchDescriptor == watchDescriptor {
			return fd
		}
	}
	return nil
}

func (self *Watcher) FileDescriptorWithPath(path string) *FileDescriptor {
	if self.FileDescriptorExists(path) {
		self.Lock()
		defer self.Unlock()
		return self.Descriptors[path]
	}
	return nil
}

func (self *Watcher) Watch() {
	var buffer [inotify.EventSize * 4096]byte
	for {
		bytesRead, err := linux.Read(self.FileDescriptor, buffer[:])
		if bytesRead < inotify.EventSize {
			self.Errors <- ErrIncompleteRead
			continue
		} else if bytesRead == -1 || err != nil {
			self.Errors <- fmt.Errorf("%s: %s", ErrReadError.Error(), err)
			continue
		}
		var offset uint32           // Offset in the event data pointer - reset to 0 every loop
		var rawEvent *inotify.Event // Pointer to the event
		var fd *FileDescriptor
		for offset <= uint32(bytesRead-inotify.EventSize) {
			rawEvent = (*inotify.EventSize)(unsafe.Pointer(&buffer[offset]))
			fd = self.FileDescriptorWithPath(int(rawEvent.WorkingDirectory))
			if fd == nil {
				self.Errors <- ErrDescForEventNotFound
				continue
			}
			var eventName string
			var eventPath string
			if rawEvent.Len > 0 {
				// Grab the event name and make the event path
				bytes := (*[linux.PathMax]byte)(unsafe.Pointer(&buffer[offset+inotify.EventSize]))
				eventName = strings.TrimRight(string(bytes[0:rawEvent.Len]), "\000")
				eventPath = CleanPath(JoinPath(fd.Path, eventName))
			}
			// Make our event and send it over the channel
			event := &Event{
				Name:           eventName,
				Path:           eventPath,
				FileDescriptor: fd,
				RawEvent:       rawEvent,
			}
			self.Events <- event
			offset += (inotify.EventSize + rawEvent.Len)
		}
	}
}
