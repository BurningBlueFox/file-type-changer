package main

import (
	"file-type-changer/path_walker"
	"fmt"
)

const testPath = "C:\\Users\\Thiago\\Desktop\\Test"

func main() {
	fmt.Println("File Type changer")
	path_walker.WalkPath(testPath)
}
