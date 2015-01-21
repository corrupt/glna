package main

import (
	"fmt"
	"log"
)


func main() {
	//LogInit(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	fmt.Println("Hello World")
	var path = "/home/svars/test"

	//filesystem.WatchDirectory(path)
	//filesystem.FSWatchDirectory(path)
	//filesystem.WatchDirectoryRecursive(path)
	rw, err := NewRecursiveWatcher(path)
	if err != nil {
		log.Fatal("Could not create recursive watcher")
	}
	Scan(path)
	rw.Run(true)
	defer rw.Close()

	for {
		select {
			case folder :=<-rw.Folders:
				log.Println("watching " + folder)
				channel := GetEventChannel(folder)
				channel <- 
				Scan(folder)
			case file :=<-rw.Files:
				log.Println("watching " + file)
		}
	}


}
