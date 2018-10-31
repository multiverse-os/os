package watch

// The running bool can be replaced with being on or off based on files in
// registered with add and remove watched files. if zero watched fiels, wwe
// are not running, so Runing bcomes a method function
type FileDescriptor struct {
	Path           string
	Mask           int
	FileDescriptor int
}
