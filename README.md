# harbor

[![GoDoc](https://godoc.org/github.com/richard-lyman/harbor?status.svg)](https://godoc.org/github.com/richard-lyman/harbor)

An implementation of the [Porter2 Stemmer](https://snowballstem.org/algorithms/english/stemmer.html).

Port: synonym: harbor

Arbor: synonym: stem (More a suffix than a stem, but close enough.)

### Install harbor binary to use on a file

Assuming you already have Golang installed, run the following command to generate an executable named ```harbor``` in whatever directory you're in:

```
GOBIN=`pwd` go get github.com/richard-lyman/harbor/cmd/harbor
```

For information on how to use the harbor command run:

```
./harbor -h
```

### Use the harbor library in a Golang project

Assuming you're using Golang modules, the following is an example of using the harbor library in a Golang project:

```
package main

import (
	"fmt"
	"github.com/richard-lyman/harbor"
)

func main() {
	fmt.Printf("%s\n", harbor.Stem([]byte("complicating")))
}
```
