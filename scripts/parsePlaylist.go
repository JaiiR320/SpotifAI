package scripts

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/JaiiR320/SpotifAI/model"
)

func ParsePlaylist(path string) {
	log.Println("Parsing playlist...")
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	json.Unmarshal(byteValue, &model.LikedSongs)
}
