// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
An implementation of the [Porter2 Stemmer](https://snowballstem.org/algorithms/english/stemmer.html).

Port: synonym: harbor

Arbor: synonym: stem (More a suffix than a stem, but close enough.)

Install harbor binary to use on a file

Assuming you already have Golang installed, run the following command to generate an executable named ```harbor``` in whatever directory you're in:

	GOBIN=`pwd` go get github.com/richard-lyman/harbor/cmd/harbor

For information on how to use the harbor command run:

	./harbor -h

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

Use the harbor library in a Golang project

Assuming you're using Golang modules, the following is an example of using the harbor library in a Golang project:

	package main

	import (
		"fmt"
		"github.com/richard-lyman/harbor"
	)

	func main() {
		fmt.Println(harbor.Stem("Stem Source"))
	}

*/
package harbor
