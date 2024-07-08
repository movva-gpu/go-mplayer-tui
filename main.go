package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/speaker"

	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/vorbis"
	"github.com/faiface/beep/wav"

	"github.com/nsf/termbox-go"
)

const audioFilePath string = "./Adikop - Bring Me Back (feat. Nieulotni) [NCS Release].mp3"

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	err = termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()

	audioFile, err := os.Open(audioFilePath)
	if err != nil {
		panic(err)
	}

	var streamer beep.StreamSeekCloser
	var format beep.Format

	tmp := strings.Split(audioFile.Name(), ".")
	switch tmp[len(tmp)-1] {
	case "mp3":
		streamer, format, err = mp3.Decode(audioFile)
	case "flac":
		streamer, format, err = flac.Decode(audioFile)
	case "wav":
		streamer, format, err = wav.Decode(audioFile)
	case "ogg":
		streamer, format, err = vorbis.Decode(audioFile)
	default:
		err = fmt.Errorf("unsupported file format: %s", tmp[len(tmp)-1])

	}
	if err != nil {
		panic(err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		panic(err)
	}
	defer speaker.Close()

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   0,
		Silent:   false,
	}

	const MAX_VOLUME = 1
	const MIN_VOLUME = -5

	speaker.Play(volume)

	for {
		drawUI(ctrl, volume, MIN_VOLUME)

		// fmt.Println("Press the key [SPACE] to pause/resume")
		// fmt.Println("Press the keys [+/-] to volume up/down")
		// fmt.Println("Press the key [M] to mute/unmute")
		// fmt.Println("Press the key [Q] to quit")

		// fmt.Println("Paused: ", ctrl.Paused)
		// fmt.Printf("%0.2f%%\n", math.Round((volume.Volume-(MIN_VOLUME))*(100-0)/(0-(MIN_VOLUME))+0))
		// fmt.Println("Muted: ", volume.Silent && !lowVolumeMute)

		ev := termbox.PollEvent()
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeySpace:
				speaker.Lock()
				ctrl.Paused = !ctrl.Paused
				speaker.Unlock()
			case ev.Ch == 'q':
				speaker.Clear()

				fmt.Println()
				fmt.Printf("%s, v%s\n", APP_NAME_SHORT, APP_VERSION)
				os.Exit(0)
			case ev.Ch == 'm':
				volume.Silent = !volume.Silent
			case ev.Ch == '+' || ev.Ch == rune(termbox.KeyArrowUp):
				volume.Silent = false

				if volume.Volume < MAX_VOLUME {
					volume.Volume += 0.5
				}
			case ev.Ch == '-' || ev.Ch == rune(termbox.KeyArrowDown):
				if volume.Volume > MIN_VOLUME {
					volume.Volume -= 0.5
				}

				if volume.Volume <= MIN_VOLUME {
					volume.Silent = true
				}
			}
		}

		time.Sleep(time.Millisecond * 100)
	}

}
