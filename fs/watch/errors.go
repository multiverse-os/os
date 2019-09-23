package watch

import "errors"

// TODO: Support localiziation
var (
	ErrDirDoesNotExist = errors.New("directory does not exist")
	//Top-level Watcher errors
	ErrWatchNotCreated = errors.New("watcher could not be created")
	//Descriptor errors
	ErrDescNotCreated       = errors.New("descriptor could not be created")
	ErrDescNotStart         = errors.New("descriptor could not be started")
	ErrDescAlreadyRunning   = errors.New("descriptor already running")
	ErrDescNotStopped       = errors.New("descriptor could not be stopped")
	ErrDescAlreadyExists    = errors.New("descriptor for that directory already exists")
	ErrDescNotRunning       = errors.New("descriptor not running")
	ErrDescForEventNotFound = errors.New("descriptor for event not found")
	ErrDescNotFound         = errors.New("descriptor not found")
	//Inotify interface errors
	ErrIncompleteRead = errors.New("incomplete event read")
	ErrReadError      = errors.New("error reading an event")
)
