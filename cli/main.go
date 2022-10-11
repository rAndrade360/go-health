package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type User struct {
	Name   string
	Height int
}

const MlPerKg = 30
const DefaultVolume = 300

func main() {
	//	if len(os.Args) < 2 {
	//		os.Exit(0)
	//	}

	log.Println(os.Args)

	var height uint

	flag.UintVar(&height, "height", 0, "Your actual height")
	ss := flag.String("bb", "asd", "")

	flag.Parse()

	if ss != nil {
		log.Println(*ss)
	}
	if height == 0 {
		log.Println(height)
		os.Exit(0)
	}

	t := time.Now()

	h := t.Hour()
	m := t.Minute()

	s := fmt.Sprintf("%d:%d", h, m+1)

	c := exec.Command("echo", `"notify-send 'Hey man' 'Toma água lá bro'"`, "|", "at", s)
	log.Println(c.String())
	log.Println(c.Run())
}

func CalcWhaterQtty(h int) int {
	return MlPerKg * h
}

func CalcGlassQtty(v int) int {
	return v / DefaultVolume
}
