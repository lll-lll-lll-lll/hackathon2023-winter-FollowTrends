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
	// OpenAI, PRTimesのクライアントの初期化
	prtimesClient := prtimes.New()
	// openaiClinet := openai.New()
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
					log.Println(message)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
						return
					}

					/*
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

						data, err := postgre.GetData()
						if err != nil {
							log.Println(err)
							return
						}
						//自前で用意したキーワード一覧
						splitedKeyWords := db.SplitKeyWord(data)

						// 自動生成されたキーワード一覧
						splitedGeneratedWord := openai.SplitWord(keyWords.Choices[0].Text)

						idx, err := keyword.Compare(splitedKeyWords, splitedGeneratedWord)
						if err != nil {
							log.Println(err)
							return
						}
						if idx == 0 {
							if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("キーワードに一致する記事はありません")).Do(); err != nil {
								log.Print(err)
								return
							}
						}

						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(data[idx].URL)).Do(); err != nil {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(keyWords.Choices[0].Text)).Do(); err != nil {
							log.Print(err)
							return
						}
					*/

				case *linebot.StickerMessage:
					replyMessage := fmt.Sprintf("ステキなスタンプをありがとうございます。")
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
						log.Print(err)
						return
					}
				//TODO: 今は1を決め打ちで入力しているので、FlexMessageの種類によってGetItemsメソッドの引数が変数化してほしいs

				//var
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

	http.HandleFunc("/notify", func(w http.ResponseWriter, req *http.Request) {
		items, _ := prtimesClient.GetItems("2")

		// メッセージを作成
		messages := []linebot.SendingMessage{
				linebot.NewTextMessage("ジョルジョル星人だよ"),
				linebot.NewTextMessage(items[0].String()),
		}

		// ブロードキャストメッセージを作成
		broadcast := bot.BroadcastMessage(messages...)

		// APIリクエストを送信
		if _, err = broadcast.Do(); err != nil {
				fmt.Println(err)
				return
		}
		fmt.Println("Broadcast message sent successfully!")
	})

	log.Print("サーバスタート")
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
