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

	fmt.Println(current)
	//fetch the lyrics
	searchParams := []mxmParams.Param{
		mxmParams.QueryArtist(current.CurrentlyPlaying.Item.Artists[0].Name),
		mxmParams.QueryTrack(current.CurrentlyPlaying.Item.Name),
	}

	searchedTracks, err := client.SearchTrack(context.Background(), searchParams...)

	if err != nil {
		fmt.Println("Error searching lyrics lyrics")
		return
	}

	track := searchedTracks[0]

	lyrics, err := client.GetTrackLyrics(context.Background(), []mxmParams.Param{
		mxmParams.TrackID(track.ID),
	}...)

	if err != nil {
		fmt.Println("Error getting lyrics")
		return
	}

	fmt.Println(lyrics.Body)

}
