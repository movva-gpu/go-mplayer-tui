package main

import (
	"fmt"
	"math"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"

	"github.com/nsf/termbox-go"
)

func drawUI(ctrl *beep.Ctrl, volume *effects.Volume, MIN_VOLUME float64) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	_, height := termbox.Size()

	pausedHelpLabel := "[Space] to pause"
	pausedLabel := "Playing"
	if ctrl.Paused {
		pausedHelpLabel = "[Space] to resume"
		pausedLabel = "Paused"
	}

	volumePercentage := int((volume.Volume - MIN_VOLUME) * 100 / (0 - MIN_VOLUME))

	muteHelpLabel := "[M] to mute"
	mutedLabel := ""
	if volume.Silent {
		muteHelpLabel = "[M] to un-mute"
		mutedLabel = "Muted"
	}

	pausedLabel = fmt.Sprintf("State: %s %s", pausedLabel, mutedLabel)
	volumeLabel := fmt.Sprintf("Volume: %d%%", volumePercentage)

	helpLabels := []string{
		"__Help:__",
		pausedHelpLabel,
		"[+]/[-] or [↑]/[↓] to increase the volume",
		muteHelpLabel,
		"[Q] to quit",
	}

	// Draw UI
	x, y := 4, 2

	for _, c := range pausedLabel {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}

	x = 4
	y++

	for _, c := range volumeLabel {
		termbox.SetCell(x, y, c, termbox.ColorDefault, termbox.ColorDefault)
		x++
	}

	x = 4
	y = height - 1 - len(helpLabels)
	for _, s := range helpLabels {
		currentFgColor := termbox.ColorDefault
		var prevC rune
		underlining := false
		for i, c := range s {
			nextC := rune(s[int(math.Min(float64(i+1), float64(len(s)-1)))])
			if c == '[' && prevC != '\\' {
				currentFgColor = termbox.ColorGreen
				prevC = c
				continue
			}
			if c == ']' && prevC != '\\' {
				currentFgColor = termbox.ColorDefault
				prevC = c
				continue
			}

			if c == '_' && prevC == '_' {
				underlining = !underlining
				prevC = c
				continue
			}
			if c == '_' && nextC == '_' && prevC != '\\' {
				prevC = c
				continue
			}
			if underlining {
				currentFgColor = termbox.ColorCyan | termbox.AttrUnderline
			}

			if c == '/' && prevC != '\\' {
				currentFgColor = termbox.AttrBold
			}

			if c == '\\' {
				prevC = c
				continue
			}

			termbox.SetCell(x, y, c, currentFgColor, termbox.ColorDefault)
			x++
			prevC = c
		}
		x = 4
		y++
	}

	termbox.Flush()
}
