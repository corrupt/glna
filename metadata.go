package main

import (
	"github.com/wtolson/go-taglib"
	"time"
	"log"
	"path"
	"path/filepath"
	"strings"
	"io/ioutil"
	"bytes"
	"os"
	"strconv"
)

var (
	CoverBase = []string {"cover", "coverart", "albumcover", "albumart"}
	CoverExt = []string {".png", ".jpg", ".jpeg"}
)

type AudioFile struct {
	path string
	artist string
	title string
	album string
	track int
	year int
	comment string
	genre string
	duration time.Duration
	bitrate int
	samplerate int
	channels int
	cover string
}

func ParseAudioFile(fname string) (af AudioFile) {

		file,err := taglib.Read(path.Clean(fname))
		if err != nil {
			log.Println(err)
			log.Fatal("Cannot access " + fname)
		}
		defer file.Close()

		af.path = fname
		af.artist = file.Artist()
		af.title = file.Title()
		af.album = file.Album()
		af.track = file.Track()
		af.year = file.Year()
		af.comment = file.Comment()
		af.genre = file.Genre()
		af.duration = file.Length()
		af.bitrate = file.Bitrate()
		af.samplerate = file.Samplerate()
		af.channels = file.Channels()

		cover := FindAlbumCover(af)
		if cover != nil {
			af.cover = cover.Name()
		}
		return
}

func FindAlbumCover(af AudioFile) os.FileInfo {

	path := filepath.Dir(af.path)
	title := strings.ToLower(af.album)
	CoverBase = append(CoverBase, title)
	CoverBase = append(CoverBase, SpaceReplace(title, ""))
	CoverBase = append(CoverBase, SpaceReplace(title, "_"))
	CoverBase = append(CoverBase, SpaceReplace(title, "+"))
	CoverBase = append(CoverBase, SpaceReplace(title, "-"))
	CoverBase = append(CoverBase, SpaceReplace(title, "."))

	files,_ := ioutil.ReadDir(path)
	for _,f := range files {
		ext := strings.ToLower(filepath.Ext(f.Name())) //lowercase file extension
		bse := strings.ToLower(BaseName(f.Name())) //lowercase file basename
		for _,ex := range CoverExt {
			if ex == ext {
				for _,bs := range CoverBase {
					if bs == bse {
						return f
					}
				}
			}
		}
	}
	return nil
}

func SpaceReplace(s string, c string) string {
	return strings.Replace(s, " ", c, -1)
}

func BaseName(f string) string {
	return strings.TrimSuffix(f, filepath.Ext(f))
}

func (af AudioFile) String() string {
	var buf bytes.Buffer

	buf.WriteString(af.path + ":\n")
	buf.WriteString("\tArtist:\t" + af.artist + "\n")
	buf.WriteString("\tTitle:\t" + af.title + "\n")
	buf.WriteString("\tAlbum:\t" + af.album + "\n")
	buf.WriteString("\tTrack:\t" + strconv.Itoa(af.track) + "\n")
	buf.WriteString("\tYear:\t" + strconv.Itoa(af.year) + "\n")
	buf.WriteString("\tLength:\t" + af.duration.String() + "\n")
	buf.WriteString("\tGenre:\t" + af.genre + "\n")
	buf.WriteString("\tBit Rate:\t" + strconv.Itoa(af.bitrate) + "\n")
	buf.WriteString("\tCover:\t" + af.cover + "\n")
	buf.WriteString("\tSample Rate:\t" + strconv.Itoa(af.samplerate) + "\n")

	return buf.String()
}

