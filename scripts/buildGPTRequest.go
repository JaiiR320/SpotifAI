package scripts

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/JaiiR320/SpotifAI/model"
	"github.com/imroc/req/v3"
	"github.com/joho/godotenv"
)

func GenerateTracks(tracks []model.Song, tags []string) []model.Song {
	log.Println("Generating tracks")
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
		return model.LikedSongs
	}

	OpenAIKey := os.Getenv("OPENAI_KEY")

	if len(tracks) == 0 || len(tags) == 0 {
		log.Println("No tracks or tags to filter by")
		return model.LikedSongs
	}

	baseURL := "https://api.openai.com/v1/chat/completions"

	str := `You are an assistant that generates a list. 
	You always return just the list with no additional description, formatting or context.
	Duplicates are not allowed in the list.
	The format of the list should follow the pattern: title1,...,titleN. Do not add any extra spaces or characters.`

	str += "\nGenerate a list of songs that match these tags: "
	for _, tag := range tags {
		str += "\"" + tag + "\","
	}

	str += " using these songs:"

	for _, song := range tracks {
		str += "\"" + song.Track.Name + "\","
	}

	body := Body{
		Model: "gpt-4-0125-preview",
		Messages: []Message{{
			Role:    "user",
			Content: str,
		}},
	}

	var response Response
	req.DevMode()
	client := req.C()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+OpenAIKey).
		SetBody(body).
		SetSuccessResult(&response).
		EnableDump().
		Post(baseURL)
	if err != nil {
		log.Println(err)
		return model.LikedSongs
	}

	if resp.GetStatusCode() != http.StatusOK {
		log.Println("Bad response from OpenAI")
		return model.LikedSongs
	}
	str = strings.ReplaceAll(response.Choices[0].Message.Content, "\"", "")

	strings.Split(str, ",")

	filteredSongs := FilterSongsByTitle(strings.Split(str, ","))
	log.Println("Filtered songs")
	return filteredSongs
}

func FilterSongsByTitle(songs []string) []model.Song {
	var filteredSongs []model.Song
	for _, song := range songs {
		for _, likedSong := range model.LikedSongs {
			if likedSong.Track.Name == song {
				filteredSongs = append(filteredSongs, likedSong)
				break
			}
		}
	}
	return filteredSongs
}

type Body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     *string `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Response struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int      `json:"created"`
	Model             string   `json:"model"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
	SystemFingerprint *string  `json:"system_fingerprint"`
}
