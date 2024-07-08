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
	
	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}

	speaker.Play(ctrl)

	for {
		fmt.Println("Press [ENTER] to pause/resume")
		fmt.Println("Press [Q]+[ENTER] to quit")
		
		var key string
		fmt.Scanln(&key)
		if key == "" {
			ctrl.Paused = !ctrl.Paused
			fmt.Println("Paused: ", ctrl.Paused)
		}

		if key == "q" {
			break
		}
	}

	fmt.Println()
	fmt.Printf("%s, v%s\n", APP_NAME_SHORT, APP_VERSION)

}
