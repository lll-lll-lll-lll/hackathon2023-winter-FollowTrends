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

func (m *Match) GetItems(data [][]string, searchTerms []string) (Items, error) {
	// 各データ配列ごとに、検索条件に一致した要素の数を格納する配列を初期化する
	matchCounts := make([]int, len(data))

	// データ配列をループし、各要素が検索条件に含まれるかどうかを確認し、一致する要素があれば、対応するmatchCountsの要素を1つ増やす
	for i, d := range data {
		for _, term := range searchTerms {
			if contains(d, term) {
				matchCounts[i]++
			}
		}
	}

	// matchCountsの中で最大の値を取得し、最大値を持つ要素のインデックスを取得する
	maxCount := 0
	maxCountIndex := -1
	for i, count := range matchCounts {
		if count > maxCount {
			maxCount = count
			maxCountIndex = i
		}
	}


	// マッチしたカウント最大の値が 0 以上かどうかを判定しておく
	if maxCount >= 0 {
		fmt.Println("一番近しいデータ: ", data[maxCountIndex])
		fmt.Println("一番近しいデータ: ", matchCounts)
		return nil, nil
	} else {
		fmt.Println("一件も見つかりませんでした。")
		return nil, nil
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}


	// ruby で書くとこんな感じ
	// data.each_with_index do |d, i|
	// 	searchTerms.each do |term|
	// 		if d.include?(term)
	// 			matchCounts[i] += 1
	// 		end
	// 	end
	// end

	// ruby で書くとこんな感じ
	// maxCounts の中で最大の値を取得する（複数ある場合は最初の値）
	//  maxCountIndex = matchCounts.each_with_index.max[0]

	// ruby で書くとこんな感じ
	// マッチしたカウント最大の値が 0 以上かどうかを判定しておく
	// if matchCounts[maxCountIndex] > 0
	// 	puts "一番近しいデータ: #{data[maxCountIndex]}"
	// else
	// 	puts "一件も見つかりませんでした。"
	// end




// // GetItemsは、検索対象のデータ配列から、指定された比較配列に一致するアイテムを探し、
// // 見つかった場合にはそのアイテムを返します。
// func (m *Match) GetItems(data [][]string, searchTerms []string) (Items, error) {
// 	// 各データ配列ごとに、検索条件に一致した要素の数を格納する配列を初期化する
// 	matchCounts := make([]int, len(data))

// 	// 各データ配列ごとに、検索条件に一致する要素の数をカウントする
// 	for i, d := range data {
// 		for _, term := range searchTerms {
// 			fmt.Printf("d: %v, term: %v", d, term)
// 			if contains(d, term) {
// 				matchCounts[i]++
// 			}
// 		}
// 	}

// 	// マッチした数が最も多い配列番号を取得する
// 	maxCountIndex := 0
// 	for i := 1; i < len(data); i++ {
// 		if matchCounts[i] > matchCounts[maxCountIndex] {
// 			maxCountIndex = i
// 		}
// 	}

// 	fmt.Printf("マッチした数: %v", matchCounts)
// 	// マッチした数が最も多い配列を出力する
// 	if matchCounts[maxCountIndex] > 0 {
// 		fmt.Printf("一番近しいデータ: %s\n", data[maxCountIndex])
// 	} else {
// 		fmt.Println("一件も見つかりませんでした。")
// 	}

// 	m.Response = Items{
// 		Item{
// 			Title: data[maxCountIndex][0],
// 			URL:   data[maxCountIndex][1],
// 		},
// 	}
// 	// マッチした数が最も多い配列を返す
// 	return m.Response, nil
// }

// // sliceに要素が含まれるかを判定する関数
// func contains(slice []string, item string) bool {
// 	for _, s := range slice {
// 		if s == item {
// 			return true
// 		}
// 	}
// 	return false
// }
