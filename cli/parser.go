package cli

import (
	"file-type-changer/path_walker"
	"flag"
	"fmt"
)

func ParseCli() *path_walker.ExtChangeJob {
	pathPtr := flag.String("path", "", "the system path for the desired application")
	extTypePtr := flag.String("new_extension_type", ".txt", "the new extension type for the files")
	whitelistPtr := flag.String("whitelist", "()(.changeme)", "the file extensions where the new type will be applied, should be encapsulated in ()")

	flag.Parse()

	fmt.Printf("path: %s - new_extension_type: %s - whitelist: %s", *pathPtr, *extTypePtr, *whitelistPtr)
	return nil
}
