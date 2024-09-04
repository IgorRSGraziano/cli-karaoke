package spotify_test

import (
	"clickaraoke/spotify"
	"context"
	"fmt"
	"log"
	"testing"
)

func TestAuth(t *testing.T) {

	client := spotify.Auth()

	if client == nil {
		t.Errorf("client is nil")
	}
}

// TODO: Achar uma forma melhor de testar para não ir no terminal autenticar toda vez
func TestClientWork(t *testing.T) {
	ctx := context.Background()

	client := spotify.Auth()

	if client == nil {
		t.Errorf("client is nil")
	}

	msg, page, err := client.FeaturedPlaylists(ctx)
	if err != nil {
		log.Fatalf("couldn't get features playlists: %v", err)
	}

	fmt.Println(msg)

	if err != nil {
		t.Errorf("error in client.FeaturedPlaylists")
	}

	if page == nil {
		t.Errorf("page is nil")
	}

}
