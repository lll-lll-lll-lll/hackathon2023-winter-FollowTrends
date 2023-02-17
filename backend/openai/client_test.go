package openai

import "testing"

func Test_SplitWord(t *testing.T) {
	input := "Windows, Linux, Mac OS X"
	want := []string{"Windows", "Linux", "Mac OS X"}
	got := SplitWord(input)
	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Logf("got %s, want %s", got[i], want[i])
			return
		}
	}

}
