package keyword

import "testing"

func Test_Compare(t *testing.T) {
	t.Parallel()
	usecase := map[string]struct {
		lineInput [][]string
		dbInput   []string
		want      int
	}{
		"success": {
			lineInput: [][]string{
				{"dd", "Li", "disdfw"},
				{"fae", "Linux", "Mac"},
				{"Windows", "Li", "Mac"},
				{"Windows", "Linux", "Mac OS X"},
			},
			dbInput: []string{"Windows", "Linux", "Mac OS X"},
			want:    3,
		},
	}
	for name, tt := range usecase {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Compare(tt.lineInput, tt.dbInput)
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
