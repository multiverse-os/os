package os

import (
	"errors"
	"runtime"
)

var executablePath, executablePathErr = func() (string, error) {
	var procfn string
	switch runtime.GOOS {
	default:
		return "", errors.New("Executable not implemented for " + runtime.GOOS)
	case "linux", "android":
		procfn = "/proc/self/exe"
	case "netbsd":
		procfn = "/proc/curproc/exe"
	case "dragonfly":
		procfn = "/proc/curproc/file"
	}
	return Readlink(procfn)
}()

func executable() (string, error) {
	return executablePath, executablePathErr
}

var initWd, errWd = Getwd()

func executable() (string, error) {
	var exePath string
	if len(Args) == 0 || Args[0] == "" {
		return "", ErrNotExist
	}
	if IsPathSeparator(Args[0][0]) {
		// Args[0] is an absolute path, so it is the executable.
		// Note that we only need to worry about Unix paths here.
		exePath = Args[0]
	} else {
		for i := 1; i < len(Args[0]); i++ {
			if IsPathSeparator(Args[0][i]) {
				// Args[0] is a relative path: prepend the
				// initial working directory.
				if errWd != nil {
					return "", errWd
				}
				exePath = initWd + string(PathSeparator) + Args[0]
				break
			}
		}
	}
	if exePath != "" {
		if err := isExecutable(exePath); err != nil {
			return "", err
		}
		return exePath, nil
	}
	// Search for executable in $PATH.
	for _, dir := range splitPathList(Getenv("PATH")) {
		if len(dir) == 0 {
			dir = "."
		}
		if !IsPathSeparator(dir[0]) {
			if errWd != nil {
				return "", errWd
			}
			dir = initWd + string(PathSeparator) + dir
		}
		exePath = dir + string(PathSeparator) + Args[0]
		switch isExecutable(exePath) {
		case nil:
			return exePath, nil
		case ErrPermission:
			return "", ErrPermission
		}
	}
	return "", ErrNotExist
}

// isExecutable returns an error if a given file is not an executable.
func isExecutable(path string) error {
	stat, err := Stat(path)
	if err != nil {
		return err
	}
	mode := stat.Mode()
	if !mode.IsRegular() {
		return ErrPermission
	}
	if (mode & 0111) == 0 {
		return ErrPermission
	}
	return nil
}

// splitPathList splits a path list.
// This is based on genSplit from strings/strings.go
func splitPathList(pathList string) []string {
	if pathList == "" {
		return nil
	}
	n := 1
	for i := 0; i < len(pathList); i++ {
		if pathList[i] == PathListSeparator {
			n++
		}
	}
	start := 0
	a := make([]string, n)
	na := 0
	for i := 0; i+1 <= len(pathList) && na+1 < n; i++ {
		if pathList[i] == PathListSeparator {
			a[na] = pathList[start:i]
			na++
			start = i + 1
		}
	}
	a[na] = pathList[start:]
	return a[:na+1]
}
