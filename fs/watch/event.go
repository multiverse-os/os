package watch

import (
	inotify "github.com/multiverse-os/os/filesystem/watch/inotify"
)

type Event struct {
	Name           string
	Path           string
	RawEvent       *inotify.Event
	FileDescriptor *FileDescriptor
}
