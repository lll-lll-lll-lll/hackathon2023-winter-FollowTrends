package prtimes

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// PRTimes APIを叩いたり、周りの処理を担当する
type PRTimes struct {}

func New()*PRTimes {
	return &PRTimes{}
}

// GetItems PRTimes APIを叩いてItem構造体を返す
func (pt *PRTimes)GetItems(categoryID int) ([]Item, error) {
	client := &http.Client{}
	url := "https://hackathon.stg-prtimes.net/api/categories/5/releases?from_date=2023-02-13"
	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Authorization", "Bearer b655dffbe1b2c82ca882874670cb110995c6604151e1b781cf5c362563eb4e12")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var items []Item
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}
	return items, nil
}