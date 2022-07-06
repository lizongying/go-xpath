# go-xpath

more simple xpath parser

## install

```
go get github.com/lizongying/go-xpath
```

## example

```
package main

import (
	"fmt"
	"github.com/lizongying/go-xpath/xpath"
)

func main() {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)
	fmt.Println(x)
}
```