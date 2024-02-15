package scripts

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/JaiiR320/SpotifAI/model"
	"github.com/imroc/req/v3"
	"github.com/joho/godotenv"
)

func GenerateTracks(tracks []model.Item, tags []string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}

	OpenAIKey := os.Getenv("OPENAI_KEY")

	if len(tracks) == 0 || len(tags) == 0 {
		return "", errors.New("no tracks or tags to filter")
	}

	baseURL := "https://api.openai.com/v1/chat/completions"

	data, err := json.Marshal(&tracks)
	if err != nil {
		return "", errors.New("error marshalling data to JSON")
	}
	str := string(data)

	str = "Filter these songs :" + str + " by these tags: "
	for _, tag := range tags {
		str += tag + ", "
	}

	str += `You are an assistant that generates a JSON. You always return just the JSON with no additional description or context.

	The format of the JSON should follow 
	
	{"songs":[{"title":"title"}]}`

	body := Body{
		Model: "gpt-3.5-turbo",
		Messages: []Message{{
			Role:    "user",
			Content: str,
		}},
	}

	var response Response

	client := req.C()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", "Bearer "+OpenAIKey).
		SetBody(body).
		SetSuccessResult(&response).
		EnableDump().
		Post(baseURL)
	if err != nil {
		return "", errors.New("error with Request to OpenAI GPT-3")
	}

	if resp.GetStatusCode() != http.StatusOK {
		return "", errors.New("no response from OpenAI GPT-3")
	}
	return response.Choices[0].Message.Content, nil
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
