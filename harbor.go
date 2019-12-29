// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package harbor

import (
	"bytes"
	"regexp"
)

var c = regexp.MustCompile(`[^aeiou]`)
var v = regexp.MustCompile(`[aeiouy]`)
var cs = regexp.MustCompile(c.String() + `[^aeiouy]*`)
var vs = regexp.MustCompile(v.String() + `[aeiou]*`)

var C = cs.String()
var V = vs.String()
var mgr0 = regexp.MustCompile(`^(` + C + `)?` + V + C)
var meq1 = regexp.MustCompile(`^(` + C + `)?` + V + C + `(` + V + `)?$`)
var mgr1 = regexp.MustCompile(`^(` + C + `)?` + V + C + V + C)
var vowelInStem = regexp.MustCompile(`^(` + C + `)?` + v.String())

var re1a1 = regexp.MustCompile(`^(.+?)(ss|i)es$`)
var re1a2 = regexp.MustCompile(`^(.+?)([^s])s$`)
var re1b1 = regexp.MustCompile(`^(.+?)eed$`)
var re1b2 = regexp.MustCompile(`^(.+?)(ed|ing)$`)
var re1b21 = regexp.MustCompile(`(at|bl|iz)$`)
var re1b22 = regexp.MustCompile(`^` + C + v.String() + `[^aeiouwxy]$`)
var re1c = regexp.MustCompile(`^(.+?)y$`)
var re2 = regexp.MustCompile(`^(.+?)(ational|tional|enci|anci|izer|bli|alli|entli|eli|ousli|ization|ation|ator|alism|iveness|fulness|ousness|aliti|iviti|biliti|logi)$`)
var re3 = regexp.MustCompile(`^(.+?)(icate|ative|alize|iciti|ical|ful|ness)$`)
var re41 = regexp.MustCompile(`^(.+?)(al|ance|ence|er|ic|able|ible|ant|ement|ment|ent|ou|ism|ate|iti|ous|ive|ize)$`)
var re42 = regexp.MustCompile(`^(.+?)(s|t)(ion)$`)
var re5 = regexp.MustCompile(`^(.+?)e$`)
var re51 = regexp.MustCompile(`^` + C + v.String() + `[^aeiouwxy]$`)
var re52 = regexp.MustCompile(`ll$`)

func removeLast(bs []byte) []byte         { return bs[:len(bs)-1] }
func addToEnd(bs []byte, s string) []byte { return append(bs, []byte(s)...) }

func Stem(bs []byte) []byte {
	bs = bytes.TrimSpace(bs)
	if len(bs) < 3 {
		return bs
	}
	if bs[0] == "y"[0] {
		bs[0] = "Y"[0]
	}
	if re1a1.Match(bs) {
		bs = bs[:len(bs)-2]
	}
	if re1a2.Match(bs) {
		bs = removeLast(bs)
	}
	if re1b1.Match(bs) {
		if subs := re1b1.FindSubmatch(bs); mgr0.Match(subs[1]) {
			bs = removeLast(bs)
		}
	} else if re1b2.Match(bs) {
		if subs := re1b2.FindSubmatch(bs); vowelInStem.Match(subs[1]) {
			bs = subs[1]
			if re1b21.Match(bs) {
				bs = addToEnd(bs, "e")
			} else if lastTwo := bs[len(bs)-2:]; len(lastTwo) == 2 && lastTwo[0] == lastTwo[1] && !bytes.ContainsAny(lastTwo, `aeiouylsz`) {
				bs = removeLast(bs)
			} else if re1b22.Match(bs) {
				bs = addToEnd(bs, "e")
			}
		}
	}
	if re1c.Match(bs) {
		if subs := re1c.FindSubmatch(bs); vowelInStem.Match(subs[1]) {
			bs = addToEnd(subs[1], "i")
		}
	}
	if re2.Match(bs) {
		if subs := re2.FindSubmatch(bs); mgr0.Match(subs[1]) {
			bs = addToEnd(subs[1], list2[string(subs[2])])
		}
	}
	if re3.Match(bs) {
		if subs := re3.FindSubmatch(bs); mgr0.Match(subs[1]) {
			bs = addToEnd(subs[1], list3[string(subs[2])])
		}
	}
	if re41.Match(bs) {
		if subs := re41.FindSubmatch(bs); mgr1.Match(subs[1]) {
			bs = subs[1]
		}
	} else if re42.Match(bs) {
		subs := re42.FindSubmatch(bs)
		if stem := append(subs[1], subs[2]...); mgr1.Match(stem) {
			bs = stem
		}
	}
	if re5.Match(bs) {
		if subs := re5.FindSubmatch(bs); mgr1.Match(subs[1]) || (meq1.Match(subs[1]) && !re51.Match(subs[1])) {
			bs = subs[1]
		}
	}
	if re52.Match(bs) && mgr1.Match(bs) {
		bs = removeLast(bs)
	}
	if bs[0] == "Y"[0] {
		bs[0] = "y"[0]
	}
	return bs
}

var list2 = map[string]string{
	"ational": "ate",
	"tional":  "tion",
	"enci":    "ence",
	"anci":    "ance",
	"izer":    "ize",
	"bli":     "ble",
	"alli":    "al",
	"entli":   "ent",
	"eli":     "e",
	"ousli":   "ous",
	"ization": "ize",
	"ation":   "ate",
	"ator":    "ate",
	"alism":   "al",
	"iveness": "ive",
	"fulness": "ful",
	"ousness": "ous",
	"aliti":   "al",
	"iviti":   "ive",
	"biliti":  "ble",
	"logi":    "log",
}

var list3 = map[string]string{
	"icate": "ic",
	"ative": "",
	"alize": "al",
	"iciti": "ic",
	"ical":  "ic",
	"ful":   "",
	"ness":  "",
}
