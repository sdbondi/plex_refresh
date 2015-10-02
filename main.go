package main

import (
	"flag"
	"fmt"
	"os"

	"sdbondi/plex"
	"sdbondi/watcher"
)

const (
	SECTION_ID int = 1
)

func printUsage() {
	fmt.Printf("Usage: %s [watch dir]\n", os.Args[0])
	os.Exit(1)
}

func refreshPlex(sectionId int) watcher.CaptureEvent {
	return func(_mask uint32, _name string) bool {
		return plex.Refresh(sectionId)
	}
}

func main() {
	if len(os.Args) <= 1 {
		printUsage()
	}

	flag.Parse()

	watchDir := flag.Arg(0)

	src, err := os.Stat(watchDir)
	if err != nil {
		fmt.Println(watchDir + " does not exist")
		os.Exit(1)
	}

	if !src.IsDir() {
		fmt.Println(watchDir + " is not a directory")
		os.Exit(1)
	}

	watcher.WatchCreateDelete(watchDir, refreshPlex(SECTION_ID))
}
