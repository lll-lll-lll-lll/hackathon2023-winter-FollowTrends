package db

import (
	"testing"
)

func Test_SplitKeyWord(t *testing.T) {
	t.Parallel()
	usecase := map[string]struct {
		input []Data
		want  []string
	}{
		"success": {
			input: []Data{
				{KeyWord: "KDDI株式会社,音楽ライブ,地域経済,04 Limited Sazabys,BiSH,自治体,開催", URL: "url"},
			},
			want: []string{"KDDI株式会社", "音楽ライブ", "地域経済", "04 Limited Sazabys", "BiSH", "自治体", "開催"},
		},
	}

	for name, tt := range usecase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := SplitKeyWord(tt.input)
			for _, v := range got {
				for i, s := range v {
					if s != tt.want[i] {
						t.Errorf("got %v, want %v", s, tt.want[i])
					}
				}
			}
		})
	}

}
