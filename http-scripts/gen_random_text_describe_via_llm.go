package http_scripts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PayloadSt struct {
	Model            string         `json:"model"`
	Messages         []MessageSt    `json:"messages"`
	Stream           interface{}    `json:"stream"`
	MaxTokens        int            `json:"max_tokens"`
	Stop             []string       `json:"stop"`
	Temperature      float64        `json:"temperature"`
	TopP             float64        `json:"top_p"`
	TopK             int            `json:"top_k"`
	FrequencyPenalty float64        `json:"frequency_penalty"`
	N                int            `json:"n"`
	ResponseFormat   ResponseFormat `json:"response_format"`
}

type MessageSt struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}

type RespSt struct {
	Id      string      `json:"id"`
	Choices []ChoicesSt `json:"choices"`
	Usage   UsageSt     `json:"usage"`
	Created int         `json:"created"`
	Model   string      `json:"model"`
	Object  string      `json:"object"`
}

type ChoicesSt struct {
	Index        int       `json:"index"`
	Message      MessageSt `json:"message"`
	FinishReason string    `json:"finish_reason"`
}

type UsageSt struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

func GenRandomTextDescribe(model, url string) string {
	payload := PayloadSt{
		Model: model,
		Messages: []MessageSt{
			{
				Role:    "user",
				Content: "随机描述一个画面",
			},
		},
		Stream:           false,
		MaxTokens:        512,
		Temperature:      0.7,
		TopP:             0.7,
		TopK:             50,
		FrequencyPenalty: 0.5,
		ResponseFormat: ResponseFormat{
			Type: "text",
		},
	}

	reqBody, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	resp, err := SendHttpRequest(url, http.MethodPost, reqBody)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	response := RespSt{}
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return response.Choices[0].Message.Content
}
