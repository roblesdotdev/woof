package woof

import "testing"

func TestWoof(t *testing.T) {
	want := "Woof!"
	got := Woof()

	if want != got {
		t.Errorf("Want %s got %s", want, got)
	}
}
