package watcher

import (
	"log"

	"golang.org/x/exp/inotify"
)

type CaptureEvent func(uint32, string) bool

func WatchCreateDelete(dir string, onEvent CaptureEvent) {
	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	if err = watcher.Watch(dir); err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case ev := <-watcher.Event:
			if ev.Mask&(inotify.IN_CREATE|inotify.IN_DELETE|inotify.IN_MOVE) != 0 {
				success := onEvent(ev.Mask, ev.Name)
				if !success {
					log.Fatal("Callback failed")
				}
			}
		case err := <-watcher.Error:
			log.Println("error:", err)
		}
	}
}
