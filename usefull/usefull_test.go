package usefull

import (
	"testing"
)

func TestAddExclamation(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"Hello, world", "Hello, world!"},
		{"!!!", "!!!!"},
		{"", "!"},
	}

	for _, c := range cases {
		got := AddExclamation(c.in)
		if got != c.want {
			t.Errorf("AddExclamation(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
