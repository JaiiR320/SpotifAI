package scripts

import (
	"log"
	"net/http"

	"github.com/JaiiR320/SpotifAI/model"
	"github.com/imroc/req/v3"
)

func AddTracksToPlaylist() {
	uris := getURIs()
	if len(uris) == 0 {
		return
	}
	playlistId := createPlaylist()

	var response Response
	baseURL := "https://api.spotify.com/v1/playlists/" + playlistId + "/tracks"

	client := req.C()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+model.AccessToken.AccessToken).
		SetBody(AddBody{Uris: uris}).
		SetSuccessResult(&response).
		EnableDump().
		Post(baseURL)
	if err != nil {
		log.Println(err)
	}
	if resp.GetStatusCode() != http.StatusCreated {
		log.Println("Bad response from Spotify")
	}
}

func getURIs() []string {
	uris := make([]string, len(model.FilteredSongs))
	for i, song := range model.FilteredSongs {
		uris[i] = "spotify:track:" + song.Track.ID

	}
	return uris
}

type AddBody struct {
	Uris []string `json:"uris"`
}

func createPlaylist() string {
	baseURL := "https://api.spotify.com/v1/users/" + model.CurrentUser.Id + "/playlists"

	var response createResponse

	body := createBody{Name: "New Playlist"}

	client := req.C()
	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+model.AccessToken.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		SetSuccessResult(&response).
		Post(baseURL)
	if err != nil {
		log.Println(err)
	}
	if resp.GetStatusCode() != http.StatusCreated {
		log.Println("Bad response from Spotify")
	}
	return response.Id
}

type createBody struct {
	Name string `json:"name"`
}

type createResponse struct {
	Id string `json:"id"`
}
