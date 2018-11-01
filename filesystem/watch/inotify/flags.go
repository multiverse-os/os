package inotify

const (
	// class: inotify flags
	IN_ACCESS        = 0x1
	IN_ALL_EVENTS    = 0xfff
	IN_ATTRIB        = 0x4
	IN_CLASSA_HOST   = 0xffffff
	IN_CLASSA_MAX    = 0x80
	IN_CLASSA_NET    = 0xff000000
	IN_CLASSA_NSHIFT = 0x18
	IN_CLASSB_HOST   = 0xffff
	IN_CLASSB_MAX    = 0x10000
	IN_CLASSB_NET    = 0xffff0000
	IN_CLASSB_NSHIFT = 0x10
	IN_CLASSC_HOST   = 0xff
	IN_CLASSC_NET    = 0xffffff00
	IN_CLASSC_NSHIFT = 0x8
	IN_CLOEXEC       = 0x80000
	IN_CLOSE         = 0x18
	IN_CLOSE_NOWRITE = 0x10
	IN_CLOSE_WRITE   = 0x8
	IN_CREATE        = 0x100
	IN_DELETE        = 0x200
	IN_DELETE_SELF   = 0x400
	IN_DONT_FOLLOW   = 0x2000000
	IN_EXCL_UNLINK   = 0x4000000
	IN_IGNORED       = 0x8000
	IN_ISDIR         = 0x40000000
	IN_LOOPBACKNET   = 0x7f
	IN_MASK_ADD      = 0x20000000
	IN_MODIFY        = 0x2
	IN_MOVE          = 0xc0
	IN_MOVED_FROM    = 0x40
	IN_MOVED_TO      = 0x80
	IN_MOVE_SELF     = 0x800
	IN_NONBLOCK      = 0x800
	IN_ONESHOT       = 0x80000000
	IN_ONLYDIR       = 0x1000000
	IN_OPEN          = 0x20
	IN_Q_OVERFLOW    = 0x4000
	IN_UNMOUNT       = 0x2000

	Accesssed = IN_ACCESS
	Modified  = IN_MODIFY
	Opened    = IN_OPEN
	Created   = IN_CREATE
	Deleted   = IN_DELETE
	Moved     = IM_MOVE

	MetadataModified = IN_ATTRIB        // node metadata changed
	WriteClosed      = IN_CLOSE_WRITE   // writable fd to file / to object was closed
	ReadClosed       = IN_CLOSE_NOWRITE // readonly fd to file /
	MovedFrom        = IN_MOVED_FROM    // (directory) had a file moved from it
	MovedTo          = IN_MOVED_TO      // (directory) had a file moved to it
	RootDeleted      = IN_DELETE_SELF   // object itself was deleted
	RootMoved        = IN_MOVE_SELF     // object itself was moved

	IsDirectory = IN_ISDIR

	AllEvents = (Accessed | Modified | MetadataModified | WriteClosed | ReadClosed | Opened | MovedFrom |
		MovedTo | MovedTo | Created | Deleted | RootDeleted | RootMoved | IsDirectory)

	DirectoryEvent   = CloseWrite | Modified | AttributeModified | IsDirectory
	DirectoryRemoved = MovedFrom | Delete | IsDirectory
	DirectoryCreated = MovedTo | Create | IsDirectory

	FileEvent   = CloseWrite | Modified | AttributeModified
	FileRemoved = MovedFrom | Delete
	FileCreated = MovedTo | Create

	RootEvent = RootDeleted | RootMoved
)
