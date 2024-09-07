package spotify

import (
	"clickaraoke/setup"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

func initAuth() {
	authenticator = spotifyauth.New(spotifyauth.WithRedirectURL("http://localhost:8080/callback"), spotifyauth.WithScopes(spotifyauth.ScopeUserReadPrivate, spotifyauth.ScopeUserReadPlaybackState), spotifyauth.WithClientID(setup.Env.Spotify.ClientId))
}

func Auth() *spotify.Client {
	tok := loadAuth()

	initAuth()

	if tok != nil {
		client := spotify.New(authenticator.Client(context.Background(), tok))
		return client
	}

	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := authenticator.AuthURL(state,
		oauth2.SetAuthURLParam("code_challenge_method", "S256"),
		oauth2.SetAuthURLParam("code_challenge", codeChallenge),
	)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	return client
}

func getAuthPath() string {
	tmp := os.TempDir()
	return filepath.Join(tmp, "spotify_auth.json")
}

// TODO: Considerar expiração do token
func loadAuth() *oauth2.Token {
	file, err := os.Open(getAuthPath())
	if err != nil {
		return nil
	}
	defer file.Close()

	var tok oauth2.Token
	err = json.NewDecoder(file).Decode(&tok)
	if err != nil {
		log.Fatal(err)
	}

	return &tok
}

func saveAuth(tok *oauth2.Token) {
	file, err := os.Create(getAuthPath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	serialized, err := json.Marshal(tok)

	if err != nil {
		log.Fatal(err)
	}

	_, err = file.Write(serialized)

	if err != nil {
		log.Fatal(err)
	}
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := authenticator.Token(r.Context(), state, r,
		oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}
	// use the token to get an authenticated client
	client := spotify.New(authenticator.Client(r.Context(), tok))
	saveAuth(tok)
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}
