package scripts

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Track struct {
	Name  string `json:"name"`
	Album struct {
		Artists []struct {
			Name string `json:"name"`
		} `json:"artists"`
		Name string `json:"name"`
	} `json:"album"`
}

type Item struct {
	Track `json:"track"`
}

type Playlist struct {
	Items []Item `json:"items"`
}

func ParsePlaylist(path string) {
	log.Println("Parsing playlist...")
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var playlist Playlist
	json.Unmarshal(byteValue, &playlist)

	for index, item := range playlist.Items {
		log.Println(index, " ", item.Track.Name)
		for _, artist := range item.Track.Album.Artists {
			log.Println("  By: ", artist.Name)
		}
	}
}
