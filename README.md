# go-xpath

more simple xpath parser

## document

For more usage, please refer to the test
[xpath_test](./test/xpath_test.go)

[document](https://pkg.go.dev/github.com/lizongying/go-xpath)

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