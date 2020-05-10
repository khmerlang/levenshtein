### Levenshtein Distance

[Levenshtein Distance](http://en.wikipedia.org/wiki/Levenshtein_distance) for calculate distance of two string in Golang.

#### Install

    go get github.com/khmerlang/levenshtein

#### Example

```go
package main

import (
  "fmt"
  "github.com/khmerlang/levenshtein"
)

func main() {
  s1 := "កកោស"
  s2 := "ក"
  fmt.Printf("ចំងាយរវាង %v និង %v គឺ %v\n",
    s1, s2, levenshtein.Distance(s1, s2))
  // -> ចំងាយរវាង កកោស និង ក គឺ 3
}

```
