package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	fmt.Println("Hello World")
	var path = "/home/svars/test"

	rw, err := NewRecursiveWatcher(path)
	if err != nil {
		log.Fatal("Could not create recursive watcher")
	}

	Scan(path)
	rw.Run(true)
	defer rw.Close()

	dq := NewDirectoryQueue()
	dq.Run()

	pups := make(chan int)
	for {
		select {
		case folder := <-rw.Folders:
			log.Println("watching " + folder)
			dq.events <- folder
		case file := <-rw.Files:
			log.Println("watching " + file)
			dq.events <- filepath.Dir(file)
		case dir := <-dq.scan:
			Scan(dir)
		}
	}
	<-pups
}
