package scripts

import (
	"log"
	"math/rand"

	"github.com/JaiiR320/SpotifAI/model"
)

func FilterSongs(songs []model.Item, tags []string) []model.Item {
	if len(tags) == 0 {
		return model.LikedSongs
	}
	//GenerateTracks(model.LikedSongs, model.Tags)
	log.Print("Filtering songs")
	selectedItems := make([]model.Item, 0)
	for i := 0; i < 15; i++ {
		randomIndex := rand.Intn(len(model.LikedSongs))
		selectedItems = append(selectedItems, model.LikedSongs[randomIndex])
	}
	return selectedItems
}

type Song struct {
	Title string `json:"title"`
}

type Songs struct {
	Songs []Song `json:"songs"`
}

// Filter Songs
// str, err := GenerateTracks(model.LikedSongs, model.Tags)
// if err != nil {
// 	log.Panic(err)
// }

// var jsonSongs Songs
// err = json.Unmarshal([]byte(str), &jsonSongs)
// if err != nil {
// 	log.Panic(err)
// }

// var filteredSongs []model.Item

// for _, s := range jsonSongs.Songs {
// 	for _, song := range songs {
// 		if strings.Contains(song.Track.Name, s.Title) {
// 			filteredSongs = append(filteredSongs, song)
// 		}
// 	}
// }

// return filteredSongs
