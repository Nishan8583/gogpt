package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

const (
	COMPLETIONS_URL = "https://api.openai.com/v1/chat/completions"
)

// OpenAIResponse holds generic response from OPENAI
type OpenAIResponse struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Usage   struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
	Choices []struct {
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
		Index        int    `json:"index"`
	} `json:"choices"`
	Error interface{} `json:"error"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Request struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature"`
}

// Chat takes a slice of Message, requests openAI completion URL based on the msg context and returns response with error if any
func (oa OpenAI) Chat(msg []Message) (OpenAIResponse, error) {
	responseMsg := OpenAIResponse{} // holds response from OPENAI API

	// Crafting message body for openAI
	req := Request{
		Model:       "gpt-3.5-turbo",
		Temperature: 0.2,
		Messages:    msg,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return responseMsg, errors.Wrap(err, "while crafting message")
	}

	bearer := "Bearer " + oa.APIKey // OPENAI uses bearer token

	body := io.NopCloser(bytes.NewReader(reqBody))
	// Crafting request
	networkReq, err := http.NewRequest(http.MethodPost, COMPLETIONS_URL, body)
	if err != nil {
		return responseMsg, errors.Wrap(err, "while crafting request")
	}
	networkReq.Header.Set("Content-Type", "application/json")
	networkReq.Header.Set("Authorization", bearer)

	// Performing the request
	response, err := oa.client.Do(networkReq)
	if err != nil {
		return responseMsg, errors.Wrap(err, "while performing request")
	}

	if err := json.NewDecoder(response.Body).Decode(&responseMsg); err != nil {
		return responseMsg, errors.Wrap(err, "while unmarshalling response")
	}

	return responseMsg, nil
}
