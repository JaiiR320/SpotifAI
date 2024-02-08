package scripts

import (
	"encoding/json"
	"log"

	"github.com/JaiiR320/SpotifAI/model"
)

type Body struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Completion struct {
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Message Message `json:"message"`
}

var OpenAIKey = "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"

func GenerateTracks(tracks []model.Item, tags []string) {
	if len(tracks) == 0 || len(tags) == 0 {
		log.Println("No tracks or tags to filter")
		return
	}

	baseURL := "https://api.openai.com/v1/chat/completions"

	data, err := json.Marshal(&tracks)
	if err != nil {
		log.Panic(err)
	}
	str := string(data)

	str = "Filter these songs :" + str + " by these tags: "
	for _, tag := range tags {
		str += tag + ", "
	}

	str += ". Respond with a JSON object with the filtered songs. You can omit all information besides the song title and artist fields from the original object. Please only respond with the JSON object text and nothing else."

	body := Body{
		Model: "gpt-3.5-turbo",
		Messages: []Message{{
			Role:    "user",
			Content: str,
		}},
	}

	var response Completion

	err = Post(baseURL).
		WithBody(body).
		WithHeader("Content-Type", "application/json").
		WithHeader("Authorization", "Bearer "+OpenAIKey).
		WithObject(&response).Do()

	if err != nil {
		log.Panic(err)
	}

	println(response.Choices[0].Message.Content)
}
