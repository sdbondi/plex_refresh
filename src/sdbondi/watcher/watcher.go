package watcher

import (
	"log"

	"gopkg.in/fsnotify.v0"
)

type CaptureEvent func(*fsnotify.FileEvent) bool

func WatchCreateDelete(dir string, onEvent CaptureEvent) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	if err = watcher.Watch(dir); err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
			if ev.IsCreate() || ev.IsDelete() || ev.IsRename() {
				success := onEvent(ev)
				if !success {
					log.Fatal("Callback failed")
				}
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
