package main

import (
	"clickaraoke/setup"
	"clickaraoke/spotify"
	"fmt"
)

func main() {
	setup.Init()
	sptf := spotify.NewSpotify()

	current, _ := sptf.GetCurrentPlaying()

	if current == nil {
		fmt.Println("No music playing")
		return
	}

	fmt.Println(current)
}
