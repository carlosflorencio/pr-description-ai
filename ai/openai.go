package ai

import (
	"context"
	"os"

	"github.com/sashabaranov/go-openai"
)

const (
	OpenAPIEnvVar = "OPENAI_API_KEY"
)

func ChatGPT(prompt string, model string) (string, error) {
	key := os.Getenv(OpenAPIEnvVar)
	client := openai.NewClient(key)

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: model,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
