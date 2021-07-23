package main

import (
	"flag"
	"fmt"
	"github.com/mohemohe/dockerhub-image-keeper/lib"
)

func main() {
	imagePtr := flag.String("image", "", "image name")
	flag.Parse()

	if *imagePtr == "" {
		panic("image name must be specified")
	}

	tags := lib.ListImageTags(*imagePtr)

	fmt.Println("pull targets:")
	for _, v := range tags {
		fmt.Printf("  %s\n", v)
	}

	total := len(tags)
	for i, v := range tags {
		fmt.Printf("\npull %s:%s ( %d / %d )\n", *imagePtr, v, i+1, total)
		lib.Pull(*imagePtr, v)
		fmt.Printf("\nrm %s:%s ( %d / %d )\n", *imagePtr, v, i+1, total)
		lib.Delete(*imagePtr, v)
	}
}
