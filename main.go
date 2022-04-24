package main

import (
	"file-type-changer/cli"
	"fmt"
)

func main() {
	fmt.Println("File Type changer")
	job := cli.ParseCli()

	if job == nil {
		fmt.Println("invalid inputs, please use the correct values")
		return
	}

	job.Run()
}
