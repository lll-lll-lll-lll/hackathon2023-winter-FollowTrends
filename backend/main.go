package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PRTIMES/nassm/db"
	"github.com/PRTIMES/nassm/prtimes"
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

	// PRTimes API 周りの処理を待った構造体を初期化
	prtimes := prtimes.New()

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
					log.Println(message.Text)
					items, err := prtimes.GetItems(message.Text)
					if err != nil {
						log.Println(err)
						w.WriteHeader(http.StatusBadRequest)
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("入力文字列に間違いがあります")).Do(); err != nil {
							log.Print(err)
						}
					}
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(items.String())).Do(); err != nil {
						log.Print(err)
					}

				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf("ステキなスタンプをありがとうございます。")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	log.Print("サーバスタート")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
