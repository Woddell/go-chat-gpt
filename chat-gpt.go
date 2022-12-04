package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type ModelResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
}

var apiUrl string = "https://api.openai.com/v1"

func Chat(message string) *ModelResponse {
	response, err := client().Do(request(message, "completions"))
	if err != nil {
		panic(err)
	}
	jsonString, err := io.ReadAll(response.Body)
	response.Body.Close()

	// Unmarshal JSON data
	var modelResponse ModelResponse
	LogToFile("API Response: " + string(jsonString[:]))
	err = json.Unmarshal([]byte(jsonString), &modelResponse)
	if err != nil {
		panic(err)
	}
	return &modelResponse
}

func request(message string, endpoint string) *http.Request {
	body, _ := json.Marshal(map[string]interface{}{
		"model":      "text-davinci-003",
		"prompt":     message,
		"max_tokens": 512,
	})

	buffer := bytes.NewBuffer(body)
	url := fmt.Sprintf("%s/%s", apiUrl, endpoint)

	request, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("CHAT_GPT_API_KEY")
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	request.Header.Add("Content-Type", "application/json")

	return request
}

func client() *http.Client {
	return &http.Client{
		Timeout: time.Second * 60,
	}
}
