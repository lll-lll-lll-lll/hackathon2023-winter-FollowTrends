// OpenAI周りの処理を担当
package openai

import (
	"context"
	"os"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type OpenAI struct {
	Client *gogpt.Client
}

// New OpenAIのシークレートの読み込みとクライアントの初期化
func New() *OpenAI {
	openai := &OpenAI{}
	openAIAPIKEY := os.Getenv("OPENAIAPIKEY")
	client := gogpt.NewClient(openAIAPIKEY)
	openai.Client = client
	return openai
}

// Do OpenAIにmessageから文章を生成するリクエストを送信するメソッド
func (oa *OpenAI) Do(message string) (gogpt.CompletionResponse, error) {
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 120,
		Prompt:    message,
	}
	resp, err := oa.Client.CreateCompletion(ctx, req)
	if err != nil {
		return gogpt.CompletionResponse{}, err
	}
	return resp, nil
}
