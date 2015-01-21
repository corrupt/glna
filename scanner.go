package main

import (
	"path/filepath"
	"log"
	"os"
	"strings"
)

var FileTypes = []string {".mp3", ".flac", ".ogg"}

func Scan(path string) (err error) {

	err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if ! info.IsDir() {
			log.Println("checking " + path + " for being a media file")
			if IsMediaFile(path) {
				af := ParseAudioFile(path)
				log.Print(af)
			}
		}
		return nil
	})
	return
}


func IsMediaFile(file string) bool {
	ext := strings.ToLower(filepath.Ext(file))

	for _,ex := range FileTypes {
		if ex == ext {
			log.Println("ex: " + ex + ", ext: " + ext)
			return true
		}
	}
	return false
}
