package musixmatch_test

import (
	"clickaraoke/setup"
	"context"
	"fmt"
	"net/http"
	"testing"

	mxm "github.com/milindmadhukar/go-musixmatch"
	"github.com/milindmadhukar/go-musixmatch/params"
)

func TestCredentials(t *testing.T) {
	setup.Init()
	client := mxm.New(setup.Env.Musixmatch.ApiKey, http.DefaultClient)

	artists, err := client.SearchArtist(context.Background(), params.QueryArtist("Martin Garrix"))

	if err != nil {
		t.Fatalf("Error searching artist: %v", err)
	}

	fmt.Println(artists[0])
}
