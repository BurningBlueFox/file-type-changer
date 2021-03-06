package cli

import (
	"file-type-changer/path_walker"
	"flag"
	"fmt"
	"regexp"
)

func ParseCli() (*path_walker.ExtChangeJob, *path_walker.LowercaseChangeJob) {
	pathPtr := flag.String("path", "", "the system path for the desired application")
	extTypePtr := flag.String("new_extension_type", "", "the new extension type for the files")
	whitelistPtr := flag.String("whitelist", "", "the file extensions where the new type will be applied, should be encapsulated in ()")
	lowercasePtr := flag.Bool("lowercase", false, "true if all files should be replaced with lowercase")

	flag.Parse()

	if *pathPtr == "" || *whitelistPtr == "" {
		fmt.Println("path or whitelist invalid, please fix the arguments")
		return nil, nil
	}

	fmt.Printf("Path: %s\nNew Extension Type: %s\n", *pathPtr, *extTypePtr)
	job := path_walker.ExtChangeJob{Path: *pathPtr, NewExtensionType: *extTypePtr, Whitelist: map[string]bool{}}

	match, _ := regexp.MatchString("\\((.*?)\\)", *whitelistPtr)
	if match {
		regex, _ := regexp.Compile("\\((.*?)\\)")
		matches := regex.FindAllString(*whitelistPtr, -1)
		fmt.Println("Whitelist:")
		for i, s := range matches {
			group := s[1 : len(s)-1]
			fmt.Printf("%d - %s\n", i, group)
			job.Whitelist[group] = true
		}
	}

	if *lowercasePtr == true {
		lowercaseJob := path_walker.LowercaseChangeJob{
			Path: *pathPtr,
		}
		return &job, &lowercaseJob
	} else {
		return &job, nil
	}
}
