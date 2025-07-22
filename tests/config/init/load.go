package tests

import (
	"log"
	"os"
	"path"
	"runtime"
)

func init() {
	SetRootDirectory()
}

func SetRootDirectory() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../../..")
	err := os.Chdir(dir)
	if err != nil {
		log.Panic(err)
	}
}
