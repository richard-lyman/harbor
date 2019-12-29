// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package harbor

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestStem(t *testing.T) {
	output, err := ioutil.ReadFile("test_data/output.txt")
	if err != nil {
		panic("Failed to read output file" + err.Error())
	}
	voc, err := ioutil.ReadFile("test_data/voc.txt")
	if err != nil {
		panic("Failed to read voc file" + err.Error())
	}
	outputWords := bytes.Fields(output)
	vocWords := bytes.Fields(voc)
	total := len(vocWords)
	failures := 0
	for i, word := range vocWords {
		s := Stem(word)
		expected := outputWords[i]
		if !bytes.Equal(s, expected) {
			failures += 1
			t.Errorf("Failed (%d): Stem(%s) != %s", i, s, expected)
		}
	}
	t.Logf("%.0f%% words passing (%d failures from %d words)", 100.0-(float64(failures)/float64(total))*100.0, failures, total)
}
