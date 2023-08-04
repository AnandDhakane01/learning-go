package learninggo 

import "testing"


func TestHello(t *testing.T) {
	got := Hello("Anand")
	want := "Hello, Anand"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
