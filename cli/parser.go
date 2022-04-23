package cli

import (
	"file-type-changer/path_walker"
	"flag"
	"fmt"
	"regexp"
)

func ParseCli() *path_walker.ExtChangeJob {
	pathPtr := flag.String("path", "", "the system path for the desired application")
	extTypePtr := flag.String("new_extension_type", ".txt", "the new extension type for the files")
	whitelistPtr := flag.String("whitelist", "()(.changName)", "the file extensions where the new type will be applied, should be encapsulated in ()")

	flag.Parse()

	fmt.Printf("Path: %s\nNew Extension Type: %s\n", *pathPtr, *extTypePtr)

	match, _ := regexp.MatchString("\\((.*?)\\)", *whitelistPtr)
	if match {
		regex, _ := regexp.Compile("\\((.*?)\\)")
		matches := regex.FindAllString(*whitelistPtr, -1)
		fmt.Println("Whitelist:")
		for i, s := range matches {
			group := s[1 : len(s)-1]
			matches[i] = group
			fmt.Printf("%d - %s\n", i, group)
		}
	}
	return nil
}
