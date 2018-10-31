package watch

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"sync"
	"unsafe"
)

type WatcherOptions struct {
	Recursive       bool
	UseWatcherFlags bool
}

type Watcher struct {
	Mutex          sync.Mutex
	Root           string
	FileDescriptor int
	DefaultMask    int
	Descriptors    map[string]*FileDescriptor
	Events         chan *Event
	Errors         chan error
	Options        *WatcherOptions
}

// Start() start starts a WatchDescriptors inotify even watcher
func (d *WatchDescriptor) Start(fd int) error {
	var err error
	if d.Running == true {
		return ErrDescAlreadyRunning
	}
	d.WatchDescriptor, err = unix.InotifyAddWatch(fd, d.Path, uint32(d.Mask))
	if d.WatchDescriptor == -1 || err != nil {
		d.Running = false
		return fmt.Errorf("%s: %s", ErrDescNotStart, err)
	}
	return nil
}

// Stop() Stop a running watch descriptor
func (self *FileDescriptor) Stop(fd int) error {
	if d.Running == false {
		return ErrDescNotRunning
	}
	_, err := unix.InotifyRmWatch(fd, uint32(d.WatchDescriptor))
	if err != nil {
		return fmt.Errorf("%s: %s", ErrDescNotStopped, err)
	}
	d.Running = false
	return nil
}

func (self *Descriptor) Exist() bool {
	_, err := os.Lstat(d.Path)
	return os.IsExist(err)
}

// DescriptorExists() returns true if a WatchDescriptor exists in w.Descriptors, false otherwise
func (w FileDescriptor) DescriptorExists(watchPath string) bool {
	w.Lock()
	defer w.Unlock()
	if _, exists := w.Descriptors[watchPath]; exists {
		return true
	}
	return false
}

// ListDescriptors() returns a string array of all WatchDescriptors in w *Watcher
func (w *Watcher) ListDescriptors() []string {
	list := make([]string, len(w.Descriptors))
	w.Lock()
	defer w.Unlock()
	for path, _ := range w.Descriptors {
		list = append(list, path)
	}
	return list
}

// RemoveDescriptor() removes the WatchDescriptor with the path matching path from the watcher,
// and stops the inotify watcher
func (w *Watcher) RemoveDescriptor(path string) error {
	if w.DescriptorExists(path) == false {
		return ErrDescNotFound
	}
	w.Lock()
	defer w.Unlock()
	descriptor := w.Descriptors[path]
	if descriptor.DoesPathExist() == true {
		if err := descriptor.Stop(descriptor.WatchDescriptor); err != nil {
			return err
		}
	}
	delete(w.Descriptors, path)
	return nil
}

// AddDescriptor() will add a descriptor to Watcher w. The descriptor is not started
// if UseWatcherFlags is true in Watcher.Options, the descriptor will use the Watcher's inotify flags
func (w *Watcher) AddDescriptor(dirPath string, mask int) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return fmt.Errorf("%s: %s", ErrDescNotCreated, "directory does not exist")
	}
	if w.DescriptorExists(dirPath) == true {
		return ErrDescNotFound
	}
	var inotifymask int
	if w.Options.UseWatcherFlags == true {
		inotifymask = w.DefaultMask
	} else {
		inotifymask = mask
	}
	w.Lock()
	w.Descriptors[dirPath] = newWatchDescriptor(dirPath, inotifymask)
	w.Unlock()
	return nil
}

// RecursiveAdd() add the directory at rootPath, and all directories below it, using the flags provided in mask
func (self Watcher) Add(path string, mask int) error {
	dirStat, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	var inotifymask int
	if w.Options.UseWatcherFlags == true {
		inotifymask = w.DefaultMask
	} else {
		inotifymask = mask
	}
	for _, child := range directoryStat {
		if child.IsDir() == true {
			childPath := path.Clean(path.Join(rootPath, child.Name()))
			if err := w.RecursiveAdd(childPath, inotifymask); err != nil {
				return err
			}
			if err := w.AddDescriptor(childPath, inotifymask); err != nil {
				return err
			}
		}
	}
	return nil
}

// NewWatcher() allocate a new watcher at path rootPath, with the default mask defaultMask
// This function initializes inotify, so it must be run first
func NewWatcher(rootPath string, defaultMask int, options *WatcherOptions) (*Watcher, error) {
	// func InotifyInit() (fd int, err error)
	fd, err := unix.InotifyInit()
	if fd == -1 || err != nil {
		return nil, fmt.Errorf("%s: %s", ErrWatchNotCreated, err)
	}
	w := &Watcher{
		RootPath:       path.Clean(rootPath),
		FileDescriptor: fd,
		DefaultMask:    defaultMask,
		Descriptors:    make(map[string]*WatchDescriptor),
		Events:         make(chan *FsEvent),
		Errors:         make(chan error),
		Options:        options,
	}
	if options.Recursive == true {
		if err := w.AddDescriptor(w.RootPath, defaultMask); err != nil {
			return w, err
		}
		return w, w.RecursiveAdd(w.RootPath, defaultMask)
	}
	return w, w.AddDescriptor(w.RootPath, defaultMask)
}

// StartAll() Start all inotify watches described by this Watcher
func (w *Watcher) StartAll() error {
	w.Lock()
	defer w.Unlock()
	for _, d := range w.Descriptors {
		if err := d.Start(w.FileDescriptor); err != nil {
			return err
		}
	}
	return nil
}

// StopAll() Stop all running watch descriptors. Does not remove descriptors from the watch
func (w *Watcher) StopAll() error {
	w.Lock()
	defer w.Unlock()
	for _, d := range w.Descriptors {
		if d.Running == true {
			if err := d.Stop(w.FileDescriptor); err != nil {
				return err
			}
		}
	}
	return nil
}

// GetDesccriptorByWatch() searches a Watcher instance for a watch descriptor.
// Searches by inotify watch descriptor
func (w *Watcher) GetDescriptorByWatch(wd int) *WatchDescriptor {
	w.Lock()
	defer w.Unlock()
	for _, d := range w.Descriptors {
		if d.WatchDescriptor == wd {
			return d
		}
	}
	return nil
}

// GetDescriptorByPath() searches a Watcher instance for a watch descriptor.
// Searches by WatchDescriptor's path
func (w *Watcher) GetDescriptorByPath(watchPath string) *WatchDescriptor {
	if w.DescriptorExists(watchPath) == true {
		w.Lock()
		defer w.Unlock()
		return w.Descriptors[watchPath]
	}
	return nil
}

// Watch() Read events from inotify, and send them over the w.Events channel
// All errors are reported over the w.Errors channel
func (w *Watcher) Watch() {
	var buffer [unix.SizeofInotifyEvent * 4096]byte
	for {
		bytesRead, err := unix.Read(w.FileDescriptor, buffer[:])
		if bytesRead < unix.SizeofInotifyEvent {
			w.Errors <- ErrIncompleteRead
			continue
		} else if bytesRead == -1 || err != nil {
			w.Errors <- fmt.Errorf("%s: %s", ErrReadError.Error(), err)
			continue
		}
		// Offset in the event data pointer - reset to 0 every loop
		var offset uint32
		// Pointer to the event
		var rawEvent *unix.InotifyEvent
		var descriptor *WatchDescriptor
		for offset <= uint32(bytesRead-unix.SizeofInotifyEvent) {
			rawEvent = (*unix.InotifyEvent)(unsafe.Pointer(&buffer[offset]))
			descriptor = w.GetDescriptorByWatch(int(rawEvent.Wd))
			if descriptor == nil {
				w.Errors <- ErrDescForEventNotFound
				continue
			}
			var eventName string
			var eventPath string
			if rawEvent.Len > 0 {
				// Grab the event name and make the event path
				bytes := (*[unix.PathMax]byte)(unsafe.Pointer(&buffer[offset+unix.SizeofInotifyEvent]))
				eventName = strings.TrimRight(string(bytes[0:rawEvent.Len]), "\000")
				eventPath = path.Clean(path.Join(descriptor.Path, eventName))
			}
			// Make our event and send it over the channel
			event := &FsEvent{
				Name:       eventName,
				Path:       eventPath,
				Descriptor: descriptor,
				RawEvent:   rawEvent,
			}
			w.Events <- event
			offset += (unix.SizeofInotifyEvent + rawEvent.Len)
		}
	}
}
