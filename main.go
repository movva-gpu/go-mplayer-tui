package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const audioFilePath string = "./Adikop - Bring Me Back (feat. Nieulotni) [NCS Release].mp3"

func main() {

	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		panic(err)
	}

	streamer, format, err := mp3.Decode(audioFile)
	if err != nil {
		panic(err)
	}

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	for {
		select {
		case <-done:
			return
		case <-time.After(time.Second):
			speaker.Lock()
			fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
			speaker.Unlock()
		}
	}

	fmt.Println()
	fmt.Printf("%s, v%s\n", APP_NAME_SHORT, APP_VERSION)

}
