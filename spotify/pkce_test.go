package spotify_test

import (
	"clickaraoke/setup"
	"clickaraoke/spotify"
	"testing"
)

func TestGenerate(t *testing.T) {
	setup.Init()
	codeVerifier, codeChallenge := spotify.GeneratePKCE()

	if codeVerifier == "" {
		t.Errorf("codeVerifier is empty")
	}

	if codeChallenge == "" {
		t.Errorf("codeChallenge is empty")
	}
}

func TestLength(t *testing.T) {
	setup.Init()
	codeVerifier, codeChallenge := spotify.GeneratePKCE()

	if len(codeVerifier) != 43 {
		t.Errorf("codeVerifier length is invalid")
	}

	if len(codeChallenge) != 43 {
		t.Errorf("codeChallenge length is invalid")
	}
}

func TestNotEqual(t *testing.T) {
	setup.Init()
	codeVerifier, codeChallenge := spotify.GeneratePKCE()

	if codeVerifier == codeChallenge {
		t.Errorf("codeVerifier and codeChallenge are equal")
	}

	codeVerifier2, codeChallenge2 := spotify.GeneratePKCE()

	if codeVerifier == codeVerifier2 {
		t.Errorf("codeVerifier is equal")
	}

	if codeChallenge == codeChallenge2 {
		t.Errorf("codeChallenge is equal")
	}
}
