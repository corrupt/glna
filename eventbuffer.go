package main

import (
	"github.com/wtolson/go-taglib"
	"time"
)

var BufferMap map [string]EventBuffer

type EventBuffer struct {
	directory string
	in <-chan taglib.Event
}

type GetEventChannel(dir string) <-chan taglib.Event {
	buf := BufferMap[dir]
	if dir == nil {
			BufferMap[dir] = &EventBuffer{directory: dir, in := make(chan, taglib.Event)}
	}
	return BufferMap[dir].in
}
