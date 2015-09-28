package main

import (
	"flag"
	"fmt"
	"os"

	"sdbondi/plex"
	"sdbondi/watcher"
)

func printUsage() {
	fmt.Printf("Usage: %s [watch dir]\n", os.Args[0])
	os.Exit(1)
}

func refreshPlex(_mask uint32, _name string) <-chan bool {
	out := make(chan bool)

	// Unnecessary but fun ;)
	go func() {
		out <- plex.Refresh(1)
		close(out)
	}()

	return out
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

	watcher.WatchCreateDelete(watchDir, refreshPlex)
}
