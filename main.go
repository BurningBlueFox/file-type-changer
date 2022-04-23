package main

import (
	"file-type-changer/path_walker"
	"fmt"
)

func main() {
	fmt.Println("File Type changer")

	newWhitelist := path_walker.Whitelist{
		"":          true,
		".material": true,
	}
	extChangeJob := path_walker.ExtChangeJob{
		Path:             "C:\\Users\\Thiago\\Desktop\\Test",
		NewExtensionType: ".txt",
		Whitelist:        newWhitelist,
	}
	extChangeJob.WalkPath()
}
