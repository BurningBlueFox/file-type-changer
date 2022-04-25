package main

import (
	"file-type-changer/cli"
	"fmt"
)

func main() {
	fmt.Println("File Type changer")
	job, lowercaseJob := cli.ParseCli()

	if job == nil {
		fmt.Println("invalid inputs, please use the correct values")
		return
	}

	job.Run()

	if lowercaseJob != nil {
		lowercaseJob.Run()
	}
}
