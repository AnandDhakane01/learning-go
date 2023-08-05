package learninggo

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("testing repeat", func(t *testing.T) {

		repeated := Repeat("a")
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("got %q want %q", repeated, expected)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	char := "a"
	for i:=0; i<b.N; i++ {
		Repeat(char)
	}
}