package cachedir

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

func Get(dir string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	base := ""
	switch runtime.GOOS {
	case "linux", "freebsd", "netbsd", "openbsd", "dragonfly", "solaris":
		base = usr.HomeDir
		dir = "." + dir
	case "darwin":
		base = filepath.Join(usr.HomeDir, "Library/Caches")
	case "windows":
		base = os.Getenv("LOCALAPPDATA")
	case "android", "nacl", "plan9":
		panic("cachedir.Get() not implemented for " + runtime.GOOS)
	}
	return filepath.Join(base, dir), nil
}
