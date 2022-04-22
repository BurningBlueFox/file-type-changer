package path_walker

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func WalkPath(dir string) {
	err := filepath.Walk(dir, walkPathCallback())
	if err != nil {
		log.Fatal(err)
	}
}

func walkPathCallback() func(path string, info fs.FileInfo, err error) error {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false {
			newFileName, changeExtErr := changeExtension(path, ".txt")
			if changeExtErr != nil {
				log.Printf("could not change filepath for file at %s \n", path)
			} else {
				logNewFileName(info, newFileName)
			}
		} else {
			log.Printf("\n\npath: %s \n", path)
		}
		return nil
	}
}

func changeExtension(path, newFileExtension string) (newFileName string, err error) {
	oldFileExtension := filepath.Ext(path)

	var newPath string
	if oldFileExtension != "" {
		newPath = strings.Replace(path, oldFileExtension, newFileExtension, 1)
	} else {
		newPath = path + newFileExtension
	}

	err = os.Rename(path, newPath)

	if err == nil {
		newFileName = filepath.Base(newPath)
	}

	return
}

func logNewFileName(info fs.FileInfo, newFileName string) {
	if info.Name() == newFileName {
		fmt.Printf("file %s did not change \n", info.Name())
	} else {
		fmt.Printf("file %s is now %s \n", info.Name(), newFileName)
	}
}
