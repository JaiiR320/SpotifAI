package scripts

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/JaiiR320/SpotifAI/model"
)

func InitializeSongList(path string) {
	log.Println("Parsing playlist...")
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	var playlist model.Playlist
	err = json.Unmarshal(byteValue, &playlist)
	if err != nil {
		log.Panic(err)
	}

	model.LikedSongs = playlist.Items
	model.FilteredSongs = playlist.Items

	log.Println("Liked songs initialized")
}

// func saveToFile() {
// 	outputFile, err := os.Create("output.json")
// 	if err != nil {
// 		log.Panic(err)
// 	}
// 	defer outputFile.Close()

// 	encodedData, err := json.Marshal(model.LikedSongs)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	_, err = outputFile.Write(encodedData)
// 	if err != nil {
// 		log.Panic(err)
// 	}
// }
