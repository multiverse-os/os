package user

import (
	"errors"
	"runtime"
	"strconv"
)

var (
	userImplemented  = true // set to false by lookup_stubs.go's init
	groupImplemented = true // set to false by lookup_stubs.go's init
)

func UserCacheDir() (string, error) {
	dir := Getenv("XDG_CACHE_HOME")
	if dir == "" {
		dir = Getenv("HOME")
		if dir == "" {
			return "", errors.New("neither $XDG_CACHE_HOME nor $HOME are defined")
		}
		dir += "/.cache"
	}
	return dir, nil
}

func UserHomeDir() string {
	switch runtime.GOOS {
	case "nacl", "android":
		return "/"
	default:
		return Getenv("HOME")
	}
}

// User represents a user account.
type User struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid string
	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid string
	// Username is the login name.
	Username string
	// Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	Name string
	// HomeDir is the path to the user's home directory (if they have one).
	HomeDir string
}

// Group represents a grouping of users.
//
// On POSIX systems Gid contains a decimal number representing the group ID.
type Group struct {
	Gid  string // group ID
	Name string // group name
}

// UnknownUserIdError is returned by LookupId when a user cannot be found.
type UnknownUserIdError int

func (e UnknownUserIdError) Error() string {
	return "user: unknown userid " + strconv.Itoa(int(e))
}

// UnknownUserError is returned by Lookup when
// a user cannot be found.
type UnknownUserError string

func (e UnknownUserError) Error() string {
	return "user: unknown user " + string(e)
}

// UnknownGroupIdError is returned by LookupGroupId when
// a group cannot be found.
type UnknownGroupIdError string

func (e UnknownGroupIdError) Error() string {
	return "group: unknown groupid " + string(e)
}

// UnknownGroupError is returned by LookupGroup when
// a group cannot be found.
type UnknownGroupError string

func (e UnknownGroupError) Error() string {
	return "group: unknown group " + string(e)
}
