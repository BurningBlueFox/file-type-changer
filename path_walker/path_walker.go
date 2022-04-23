package path_walker

import (
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Whitelist map[string]bool

// WalkPath Walk the path and it's subdirectories while changing all files ending
// with changedExtensions to the newExtensionType
func WalkPath(dir, newExtensionType string, whitelist Whitelist) {
	err := filepath.Walk(dir, walkPathCallback(newExtensionType, whitelist))
	if err != nil {
		log.Fatal(err)
	}
}

func walkPathCallback(newExt string, whitelist Whitelist) func(path string, info fs.FileInfo, err error) error {
	return func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() == false {
			applyExtensionChange(path, newExt, whitelist, info)
		} else {
			ct.ChangeColor(ct.Blue, true, ct.None, false)
			log.Printf("\n\npath: %s \n", path)
		}
		return nil
	}
}

func applyExtensionChange(path string, newExt string, whitelist Whitelist, info fs.FileInfo) {
	if whitelist[filepath.Ext(path)] == false {
		return
	}
	newFileName, changeExtErr := changeExtension(path, newExt)
	if changeExtErr != nil {
		ct.ChangeColor(ct.Red, true, ct.None, false)
		log.Printf("could not change filepath for file at %s \n", path)
	} else {
		logNewFileName(newFileName, info)
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

func logNewFileName(newFileName string, info fs.FileInfo) {
	if info.Name() == newFileName {
		ct.ChangeColor(ct.Black, true, ct.None, false)
		fmt.Printf("file: %s did not change \n", info.Name())
	} else {
		ct.ChangeColor(ct.Green, true, ct.None, false)
		fmt.Printf("file: %s -> %s \n", info.Name(), newFileName)
	}
}
