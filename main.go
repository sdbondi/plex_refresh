package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
	"gopkg.in/fsnotify.v0"

	"sdbondi/plex"
	"sdbondi/watcher"
)

func printUsage() {
	os.Exit(1)
}

func refreshPlex(token string, sectionId int) watcher.CaptureEvent {
	return func(_event *fsnotify.FileEvent) bool {
		return plex.Refresh(token, sectionId)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "plex refresher"
	app.Action = run
	app.Usage = fmt.Sprintf("%s [flags] watch-dir\n", os.Args[0])
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "section, s",
			Value: 1,
		},
		cli.StringFlag{
			Name:   "plex-token, t",
			Usage:  "Plex token",
			EnvVar: "PLEX_TOKEN",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

type Error struct {
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func run(c *cli.Context) error {
	watchDir := c.Args().Get(0)

	src, err := os.Stat(watchDir)
	if err != nil {
		return Error{
			Message: watchDir + " does not exist",
		}
	}

	if !src.IsDir() {
		return Error{
			Message: watchDir + " is not a directory",
		}
	}

	fmt.Printf("Watching directory: %s\n", src.Name())
	watcher.WatchCreateDelete(
		watchDir,
		refreshPlex(c.String("plex-token"), c.Int("section")),
	)

	return nil
}
