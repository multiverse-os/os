package inotify

// Inotify based watch implementation
const EventSize = 0x10

type Event struct {
	Name             [0]int8
	WorkingDirectory int32
	Mask             uint32
	Cookie           uint32
	Length           uint32
}
