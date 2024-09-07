package spotify

import (
	"context"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

var (
	ch = make(chan *spotify.Client)
	//TODO Pra que serve esse state?
	state         = "abc123"
	authenticator *spotifyauth.Authenticator
	//TODO: Gerar a cada request
	//! PKCE Realmente é o melhor método de autenticação? Já que vai utilizar secret via .env do Musixmatch, então secret do Spotify não seria um problema
	// These should be randomly generated for each request
	//  More information on generating these can be found here,
	// https://www.oauth.com/playground/authorization-code-with-pkce.html
	codeVerifier, codeChallenge = GeneratePKCE()
)

type Spotify struct {
	Client *spotify.Client
	ctx    context.Context
}

func NewSpotify() *Spotify {
	client := Auth()
	ctx := context.Background()
	return &Spotify{
		Client: client,
		ctx:    ctx,
	}
}

func (s *Spotify) GetCurrentPlaying() (*spotify.PlayerState, error) {
	return s.Client.PlayerState(s.ctx)
}
