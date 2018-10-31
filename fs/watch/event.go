package watch

type Event struct {
	Name       string
	Path       string
	Raw        *InotifyEvent
	Descriptor *WatchDescriptor
}
