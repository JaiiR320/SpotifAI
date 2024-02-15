package scripts

import (
	"log"
	"math/rand"

	"github.com/JaiiR320/SpotifAI/model"
)

func FilterSongs(songs []model.Item, tags []string) []model.Item {
	if len(tags) == 0 {
		log.Println("No tags to filter")
		return model.LikedSongs
	}
	log.Print("Filtering songs")
	selectedItems := make([]model.Item, 0)
	for i := 0; i < 20; i++ {
		randomIndex := rand.Intn(len(model.LikedSongs))
		selectedItems = append(selectedItems, model.LikedSongs[randomIndex])
	}
	//return GenerateTracks(model.LikedSongs, model.Tags)
	return selectedItems
}

type Song struct {
	Title string `json:"title"`
}

type Songs struct {
	Songs []Song `json:"songs"`
}
