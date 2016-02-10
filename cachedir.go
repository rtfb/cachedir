package cachedir

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

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
