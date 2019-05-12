package ginhelper

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func RootPath() string {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		panic(err)
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		panic(errors.New("failed to get root path"))
	}
	return string(path[0 : i+1])
}
