package fs

import "runtime"

func TempDir() string {
	dir := Getenv("TMPDIR")
	if dir == "" {
		if runtime.GOOS == "android" {
			dir = "/data/local/tmp"
		} else {
			dir = "/tmp"
		}
	}
	return dir
}
