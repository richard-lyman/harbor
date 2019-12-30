// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package harbor

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"
)

func TestStem(t *testing.T) {
	t.Parallel()
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

func TestStemMap(t *testing.T) {
	t.Parallel()
	result, err := StemMap(strings.NewReader("read ready reads reading reader readiness"))
	if err != nil {
		t.Errorf("StemMap failed: %s", err)
	}
	if len(result) != 3 {
		t.Log(result)
		t.Fail()
	}
	if len(result["read"]) != 3 || result["read"][0] != "read" ||
		result["read"][1] != "reads" || result["read"][2] != "reading" {
		t.Fail()
	}
	if len(result["reader"]) != 1 || result["reader"][0] != "reader" {
		t.Fail()
	}
	if len(result["readi"]) != 2 || result["readi"][0] != "ready" || result["readi"][1] != "readiness" {
		t.Fail()
	}
}
