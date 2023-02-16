// PRTimes API周りの処理全般を担当
package prtimes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// PRTimes APIを叩いたり、周りの処理を担当する
type PRTimes struct {
	Response Items `json:"response"`
}

func New() *PRTimes {
	return &PRTimes{}
}

type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// String LINE上に表示するフォーマット
func (i Item) String() string {
	return fmt.Sprintf("%s:%s", i.Title, i.URL)
}

type Items []Item

// String スライスのItem構造体を文字列にするメソッド
func (i Items) String() string {
	var s string
	for _, v := range i {
		s += v.String()
	}
	return s
}

// GetItems PRTimes APIを叩いてItem構造体を返す
func (pt *PRTimes) GetItems(categoryID string) (Items, error) {
	client := &http.Client{}
	url := fmt.Sprintf("https://hackathon.stg-prtimes.net/api/categories/%s/releases?from_date=2023-02-13", categoryID)
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
	var items Items
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}
	pt.Response = items
	return items, nil
}
