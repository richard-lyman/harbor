// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

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
