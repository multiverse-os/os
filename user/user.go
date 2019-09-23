package filesystem

import (
	"errors"
	"runtime"
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
