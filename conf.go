package main

import "fmt"

const PATCH_VERSION uint8 = 2
const MINOR_VERSION uint8 = 2
const MAJOR_VERSION uint8 = 0

const APP_NAME string = "Go Music Player TUI"
const APP_NAME_SHORT string = "gmp"

const APP_LICENSE string = "MIT"
const APP_AUTHOR string = "Danyella Strikann"
const APP_URL string = "https://github.com/danyell/go-mplayer-tui"

var APP_VERSION string = fmt.Sprintf("%d.%d.%d", MAJOR_VERSION, MINOR_VERSION, PATCH_VERSION)
