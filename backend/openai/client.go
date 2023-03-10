// OpenAI周りの処理を担当
package openai

import (
	"context"
	"fmt"
	"os"
	"regexp"
	"strings"

	gogpt "github.com/sashabaranov/go-gpt3"
)

// キーワード抽出の際に先頭に必ずつける必要があるテキスト
const ExtractKeyWordsHead = "Extract keywords from this text:\n\n"

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

// GenerateSentence OpenAIにmessageから文章を生成するリクエストを送信するメソッド
func (oa *OpenAI) GenerateSentence(message string) (gogpt.CompletionResponse, error) {
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 300,
		Prompt:    message,
		// ランダムな文章を生成したいので、1に近づけた
		Temperature: 0.69,
	}
	resp, err := oa.Client.CreateCompletion(ctx, req)
	if err != nil {
		return gogpt.CompletionResponse{}, fmt.Errorf("文章の自動生成に失敗しました。%w", err)
	}
	return resp, nil
}

// ExtractKeyWords 引数の文字列からキーワードを抽出
func (oa *OpenAI) ExtractKeyWords(text string) (gogpt.CompletionResponse, error) {
	ctx := context.Background()
	req := gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 100,
		// キーワードのランダム性を抑えるために0に近づけた
		Temperature: 0.2,
		// 同じようなキーワードを抽出させないために2に近づけた
		FrequencyPenalty: 1.5,
		Prompt:           fmt.Sprintf("%s %s", ExtractKeyWordsHead, text),
	}
	resp, err := oa.Client.CreateCompletion(ctx, req)
	if err != nil {
		return gogpt.CompletionResponse{}, fmt.Errorf("自動生成した文章からキーワードを抽出できませんでした。%w", err)
	}
	convertedText, err := ConvertRegexKeyWords(resp.Choices[0].Text)
	if err != nil {
		return gogpt.CompletionResponse{}, fmt.Errorf("キーワードを正規表現で整形するコンバート時に失敗しました%w", err)
	}
	resp.Choices[0].Text = convertedText
	return resp, nil
}

// ConvertRegexKeyWords 正規表現でKeywords以降の文字列を取得
func ConvertRegexKeyWords(keyWords string) (string, error) {
	re, err := regexp.Compile(`Keywords:(.*)`)
	if err != nil {
		return "", err
	}
	match := re.FindStringSubmatch(keyWords)
	if len(match) > 1 {
		return match[1], nil
	}
	return "", nil
}

func SplitWord(text string) []string {
	return strings.Split(text, ",")
}
