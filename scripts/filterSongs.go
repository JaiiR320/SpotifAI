package scripts

import (
	"math/rand"

	"github.com/JaiiR320/SpotifAI/model"
)

func FilterSongs(songs []model.Item, tags []string) []model.Item {
	selectedItems := make([]model.Item, 0)
	for i := 0; i < 15; i++ {
		randomIndex := rand.Intn(len(model.LikedSongs.Items))
		selectedItems = append(selectedItems, model.LikedSongs.Items[randomIndex])
	}
	return selectedItems
}
