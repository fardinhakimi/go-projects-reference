package hello

import "testing"

func TestHello(t *testing.T) {
	want := "hi there"
	got := Hello(want)

	if got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
