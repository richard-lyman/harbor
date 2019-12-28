// Copyright 2019 Richard Lyman. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

/*
An implementation of the [Porter2 Stemmer](https://snowballstem.org/algorithms/english/stemmer.html).

Port: synonym: harbor

Stem: synonym: arbor (More a suffix than a stem, but close enough.)

### Install harbor binary to use on a file

Assuming you already have Golang installed, run the following command to generate an executable named ```harbor``` in whatever directory you're in:

	GOBIN=`pwd` go get github.com/richard-lyman/harbor/cmd/harbor

### Use the harbor library in a Golang project

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
