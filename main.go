package main

import (
	"file-type-changer/path_walker"
	"fmt"
)

func main() {
	fmt.Println("File Type changer")

	var (
		path               string
		targetExtension    string
		affectedExtensions path_walker.Whitelist
	)

	path = "C:\\Users\\Thiago\\Desktop\\Test"
	targetExtension = ".txt"
	affectedExtensions = path_walker.Whitelist{
		"":          true,
		".material": true,
	}
	path_walker.WalkPath(path, targetExtension, affectedExtensions)
}
