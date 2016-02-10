// Package cachedir provides an easy way to hide OS specifics when constructing
// a path to a directory suitable for storing files. It picks the path for each
// GOOS in the following fashion:
// 	* on Linux, *BSD and Solaris: $HOME/
// 	* on OSX: $HOME/Library/Caches/
// 	* on Windows: %LOCALAPPDATA%\
package cachedir

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// Get constructs a path by appending any number of subdirectories to an
// OS-dependent root (see the package description). On Linux and other systems
// where root is $HOME, the first parameter to Get will be prepended with a
// '.', to make it a "hidden" dot-file. Remaining path elements are not
// touched.
//
// Since Get relies on path/filepath for implementation, the parameters can
// contain slash-separated path segments and are properly handled by
// filepath.Clean.
//
// Get only constructs the path, it does not touch the file system. The caller
// is responsible for calling os.MkdirAll if necessary.
func Get(elem ...string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	base := ""
	head := elem[0]
	switch runtime.GOOS {
	case "linux", "freebsd", "netbsd", "openbsd", "dragonfly", "solaris":
		base = usr.HomeDir
		head = "." + head
	case "darwin":
		base = filepath.Join(usr.HomeDir, "Library/Caches")
	case "windows":
		base = os.Getenv("LOCALAPPDATA")
	case "android", "nacl", "plan9":
		return "", errors.New("cachedir.Get() not implemented for " + runtime.GOOS)
	}
	chain := append([]string{base, head}, elem[1:]...)
	return filepath.Join(chain...), nil
}
