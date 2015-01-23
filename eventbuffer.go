package main

import (
	"time"
	"log"
)

var defaultDuration = 5 * time.Second

type DirectoryQueue struct {
	bufferMap map[string]*time.Timer
	events    chan string
	scan      chan string
}

func (dq *DirectoryQueue) Run() {
	go func() {
		for {
			select {
			case dir := <-dq.events:
				log.Println("blubb")
				timer, ok := dq.bufferMap[dir]
				if ok {
					timer.Reset(defaultDuration)
				} else {
					dq.bufferMap[dir] = time.AfterFunc(defaultDuration, func() {
						dq.scan <- dir
						delete(dq.bufferMap, dir)
					})
				}
			}
		}
	}()
}

func NewDirectoryQueue() (dq *DirectoryQueue) {
	dq = new(DirectoryQueue)
	dq.bufferMap = make(map[string]*time.Timer)
	dq.events = make(chan string)
	dq.scan = make(chan string)
	return
}
