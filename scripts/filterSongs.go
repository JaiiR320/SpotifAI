package scripts

import (
	"github.com/JaiiR320/SpotifAI/model"
)

func FilterSongs(songs []model.Item, tags []string) []model.Item {

	GenerateTracks(model.LikedSongs.Items, model.Tags)

	// selectedItems := make([]model.Item, 0)
	// for i := 0; i < 15; i++ {
	// 	randomIndex := rand.Intn(len(model.LikedSongs.Items))
	// 	selectedItems = append(selectedItems, model.LikedSongs.Items[randomIndex])
	// }
	return model.LikedSongs.Items
}
