package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PRTIMES/nassm/db"
	"github.com/PRTIMES/nassm/openai"
	"github.com/PRTIMES/nassm/prtimes"
	"github.com/PRTIMES/nassm/match"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}
	postgresql := db.NewPostgreSql()
	postgresql.Init()
	defer postgresql.Db.Close()

	// OpenAI, PRTimesのクライアントの初期化
	prtimesClient := prtimes.New()
	openaiClinet := openai.New()

	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					res, err := openaiClinet.GenerateSentence(message.Text)
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusBadRequest)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("文章の自動生成に失敗しました")).Do(); err != nil {
							log.Print(err)
							return
						}
					}

					keyWords, err := openaiClinet.ExtractKeyWords(res.Choices[0].Text)
					if err != nil {
						log.Println(err)
						return
					}

					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(keyWords.Choices[0].Text)).Do(); err != nil {
						log.Print(err)
						return
					}
				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf("ステキなスタンプをありがとうございます。")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
						return
					}
				case *linebot.FlexMessage:
					items, err := prtimesClient.GetItems("1")
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusBadRequest)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("入力文字列に間違いがあります")).Do(); err != nil {
							log.Print(err)
							return
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(items.String())).Do(); err != nil {
						log.Print(err)
						return
					}
				}
			}
		}
	})

	http.HandleFunc("/", matchHandler)

	log.Print("サーバスタート")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

// Path: backend/match/match.go
func matchHandler(w http.ResponseWriter, req *http.Request) {
	// マッチング処理を待った構造体を初期化
	match := match.New()
	// 検索対象の配列
	var data = [][]string{
		{"apple", "banana", "cherry"},
		{"apricot", "blackberry", "cherry", "date"},
		{"avocado", "blackberry", "cherry", "date", "elderberry", "fig"},
		{"avocado", "blackberry", "cherry", "date", "elderberry", "fig", "grape"},
		{"avocado", "blackberry", "cherry", "date", "elderberry", "fig", "fig"},
	}

	// 比較対象の配列
	var compare = []string{"cherry", "date", "fig"}

	// マッチング処理を行う
	match.GetItems(data, compare)
}
