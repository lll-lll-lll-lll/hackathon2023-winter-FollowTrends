package keyword

import "testing"

func Test_Compare(t *testing.T) {
	t.Parallel()
	usecase := map[string]struct {
		dbInput   [][]string
		LineInput []string
		want      int
	}{
		"success": {
			dbInput: [][]string{
				{"dd", "Li", "disdfw"},
				{"fae", "Linux", "Mac"},
				{"Windows", "Li", "Mac"},
				{"Windows", "Linux", "Mac OS X"},
			},
			LineInput: []string{"Windows", "Linux", "Mac OS X"},
			want:      3,
		},
		"success2": {
			dbInput: [][]string{
				{"Windows", "0x80070103", "アップデート", "互換性", "リセット", "4DDiG", "Windows Boot Genius", "Tenorshare", "ファイル復元", "重複ファ"},
				{"Windows", " for Restaurants", "予約台帳", "AI", "電話予約代行システム", "2018", "キャンセル料徴収"},
			},
			LineInput: []string{"Windows", "12月", "2018年", "10 October", "2018", "Update ", "1809 ", "2019年1月"},
			want:      1,
		},
	}
	for name, tt := range usecase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Compare(tt.dbInput, tt.LineInput)
			if err != nil {
				t.Error(err)
				return
			}
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
				return
			}
		})
	}

}
