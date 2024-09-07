package main

import (
	"clickaraoke/setup"
	"clickaraoke/spotify"
	"context"
	"fmt"
	"net/http"

	mxm "github.com/milindmadhukar/go-musixmatch"
	mxmParams "github.com/milindmadhukar/go-musixmatch/params"
)

func main() {
	setup.Init()
	sptf := spotify.NewSpotify()

	current, _ := sptf.GetCurrentPlaying()

	if current == nil {
		fmt.Println("No music playing")
		return
	}

	client := mxm.New(setup.Env.Musixmatch.ApiKey, http.DefaultClient)

	lyrics, err := client.GetTrackLyrics(context.Background(), []mxmParams.Param{
		mxmParams.TrackISRC(current.CurrentlyPlaying.Item.ExternalIDs["isrc"]),
	}...)

	if err != nil {
		fmt.Println("Error getting lyrics")
		return
	}

	fmt.Println(lyrics.Body)

}
