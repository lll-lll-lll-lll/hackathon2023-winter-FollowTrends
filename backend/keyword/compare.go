package keyword

import (
	"errors"
	"log"
)

type Item struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Compare 一番マッチした
func Compare(dbKeyWords [][]string, lineKeyWords []string) (int, error) {
	// 各データ配列ごとに、検索条件に一致した要素の数を格納する配列を初期化する
	matchCounts := make([]int, len(dbKeyWords))

	// データ配列をループし、各要素が検索条件に含まれるかどうかを確認し、一致する要素があれば、対応するmatchCountsの要素を1つ増やす
	for i, d := range dbKeyWords {
		for _, term := range lineKeyWords {
			if contains(d, term) {
				matchCounts[i]++
			}
		}
	}

	// matchCountsの中で最大の値を取得し、最大値を持つ要素のインデックスを取得する
	maxCount := 0
	maxCountIndex := 0
	for i, count := range matchCounts {
		if count > maxCount {
			maxCount = count
			maxCountIndex = i
		}
	}
	log.Println(matchCounts)

	// マッチしたカウント最大の値が 0 以上かどうかを判定しておく
	if maxCount >= 0 {
		return maxCountIndex, nil
	} else {
		return 0, errors.New("一件も見つかりませんでした")
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
