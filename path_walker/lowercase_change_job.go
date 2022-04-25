package path_walker

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type LowercaseChangeJob struct {
	Path string
}

func (e *LowercaseChangeJob) Run() {
	err := filepath.Walk(e.Path, lowercaseCallback())
	if err != nil {
		log.Fatal(err)
	}
}

func lowercaseCallback() func(path string, info fs.FileInfo, err error) error {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false {
			applyLowercase(path, info)
		}
		return nil
	}
}

func applyLowercase(path string, info fs.FileInfo) {
	originalFileName := info.Name()

	newPath := strings.Replace(path, originalFileName, strings.ToLower(originalFileName), 1)

	err := os.Rename(path, newPath)

	if err != nil {
		fmt.Println("err on rename")
	}

	return
}
