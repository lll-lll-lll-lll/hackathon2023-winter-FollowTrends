package match

import (
	"fmt"
)

// match APIを叩いたり、周りの処理を担当する
type Match struct {
	Response Items `json:"response"`
}

// New Match構造体を返す
func New() *Match {
	return &Match{}
}

// Item 構造体
type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// String LINE上に表示するフォーマット
func (i Item) String() string {
	return fmt.Sprintf("%s:%s", i.Title, i.URL)
}

// Items Item構造体のスライス
type Items []Item

// String スライスのItem構造体を文字列にするメソッド
func (mt *Match) GetItems(data [][]string, compare []string) (Items, error) {
	fmt.Println("GetItems/Match")
	// // 検索対象の配列
	// data := [][]string{
	// 	{"apple", "banana", "cherry"},
	// 	{"apricot", "blackberry", "cherry", "date"},
	// 	{"avocado", "blackberry", "cherry", "date", "elderberry", "fig"},
	// }

	// // 比較対象の配列
	// compare := []string{"cherry", "date", "fig"}
	fmt.Println("data:", data)
	fmt.Println("compare:", compare)
	// 一番近しいデータを格納する変数を初期化
	var closestMatch string

	// 一番近しいデータを探す
	for _, d := range data {
		// 全てのcompareの要素がdataの中に含まれるかをチェックする
		foundAll := true
		for _, c := range compare {
			if !contains(d, c) {
				foundAll = false
				break
			}
		}
		// 全てのcompareの要素がdataの中に含まれる場合、一番近しいデータとして格納する
		if foundAll {
			closestMatch = d[0]
			break
		}
	}

	// closestMatchが空でなければ、結果を出力する
	if closestMatch != "" {
		fmt.Printf("一番近しいデータ: %s\n", closestMatch)
	} else {
		fmt.Println("一件も見つかりませんでした。")
	}
	return nil, nil
}

// sliceに要素が含まれるかを判定する関数
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}