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

	outputFile, err := os.Create("output.json")
	if err != nil {
		log.Panic(err)
	}
	defer outputFile.Close()

	encodedData, err := json.Marshal(model.LikedSongs)
	if err != nil {
		log.Panic(err)
	}

	_, err = outputFile.Write(encodedData)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Playlist parsed and output to output.json.")

}
