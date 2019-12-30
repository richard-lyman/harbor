// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"github.com/richard-lyman/harbor"
	"io"
	"os"
	"text/template"
	"unicode"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stdout, `

The harbor command has very few options.

 1. To process STDIN, set the last argument to a dash ("-"):
     ./harbor -
  As another example:
     cat file.txt | ./harbor -

 2. To process a file, pass it as the last argument:
     ./harbor file.txt
  You can pass as many files as you want, each will be processed serially in the sequence given.

 3. To modify the output format, pass it as a string value assigned to the format flag ("-f"):
     ./harbor -f '{{range .}}{{printf "%%s\n" .Stem}}{{end}}' file.txt
  The type to be formated is a slice of structs, where each struct has three fields: Pos, Word, and Stem.
  The Pos field stores the byte position in the input where the word was found. Word and Stem are what they sound like.
  For more information on valid values for this flag, see https://golang.org/pkg/text/template/

 4. To modify the output format using a built-in formatter, pass one of the following values to the format flag:
  - "default": This is the format used when the format flag is left unspecified:
      ./harbor -f "default" file.txt
  - "plain": This provides each Stem on it's own line:
      ./harbor -f "plain" file.txt
  - "compact": This provides each Stem separated by a space:
      ./harbor -f "compact" file.txt
  - "csv": Stem and Word use the string %%q format: https://golang.org/pkg/fmt/:
      ./harbor -f "csv" file.txt
  - "json": This uses a template func 'inner', which is available in any format:
      ./harbor -f "json" file.txt


`)
		os.Exit(0)
	}
}

const DEFAULT = "default"
const PLAIN = "plain"
const COMPACT = "compact"
const CSV = "csv"
const JSON = "json"
const defaultFormat = `RESULT:

{{range $i, $t := .}}{{printf "%d: %s\n" $i $t}}{{end}}
`
const plainFormat = `{{range $i, $t := .}}{{inner $i "\n"}}{{printf "%s" $t.Stem}}{{end}}`
const compactFormat = `{{range $i, $t := .}}{{inner $i " "}}{{printf "%s" $t.Stem}}{{end}}`
const csvFormat = `INDEX,POS,WORD,STEM
{{range $i, $t := .}}{{inner $i "\n"}}{{printf "%d,%d,%q,%q" $i $t.Pos $t.Word $t.Stem}}{{end}}`
const jsonFormat = `[
{{range $i, $t := .}}
{{- inner $i ","}}{{printf "{\"index\":%d,\"pos\":%d,\"word\":%q,\"stem\":%q}" $i $t.Pos $t.Word $t.Stem}}
{{end}}]`

var format = flag.String("f", DEFAULT, "")

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("\n\nFAILURE: Failed to provide required options. Please provide files to process or signal to process STDIN")
		flag.Usage()
		os.Exit(1)
	}
	switch *format {
	case DEFAULT:
		*format = defaultFormat
	case PLAIN:
		*format = plainFormat
	case COMPACT:
		*format = compactFormat
	case CSV:
		*format = csvFormat
	case JSON:
		*format = jsonFormat
	}
	args := flag.Args()
	if len(args) == 1 && args[0] == "-" {
		output(process(os.Stdin))
	} else if len(args) > 0 {
		rs := []io.Reader{}
		for _, arg := range args {
			r, err := os.Open(arg)
			if err != nil {
				panic(fmt.Sprintf("Failed to open '%s': %s", arg, err))
			}
			rs = append(rs, r)
		}
		output(process(io.MultiReader(rs...)))
	}
}

type triplet struct {
	Pos  int
	Word []byte
	Stem []byte
}

func (t triplet) String() string {
	return fmt.Sprintf("@%d %s %s", t.Pos, t.Word, t.Stem)
}

func process(inr io.Reader) []triplet {
	in := bufio.NewReader(inr)
	result := []triplet{}
	var tmp bytes.Buffer
	pos := 0
	for {
		r, n, err := in.ReadRune()
		if n > 0 {
			pos = pos + n
		}
		if n > 0 && unicode.IsLetter(r) {
			tmp.WriteRune(r)
		} else { // Either we read something that wasn't a letter, or we read nothing
			if tmp.Len() > 0 {
				tmptmp := make([]byte, tmp.Len())
				copy(tmptmp, tmp.Bytes())
				result = append(result, triplet{pos - len(tmptmp) - 1, tmptmp, harbor.Stem(tmptmp)})
				tmp.Reset()
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil && err != io.EOF {
			fmt.Printf("Failed to read rune: %s\n", err)
			os.Exit(1)
		}
	}
	return result
}

func output(ts []triplet) {
	if err := template.Must(template.New("").Funcs(template.FuncMap{
		"inner": func(i int, s string) string {
			if i == 0 {
				return ""
			}
			return s
		},
	}).Parse(*format)).Execute(os.Stdout, ts); err != nil {
		fmt.Printf("Failed to execute given template format: %s", err)
		os.Exit(1)
	}
}
