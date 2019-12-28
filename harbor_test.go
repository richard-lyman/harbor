package harbor

import (
	"testing"
)

func TestStem(t *testing.T) {
	s := "Stem Source"
	d := "stem source"
	if out := Stem(s); out != d {
		t.Errorf("Stem(%q) = %q, want %q", s, out, d)
	}
}
