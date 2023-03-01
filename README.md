# go-xpath

simple xpath selector

[document](https://pkg.go.dev/github.com/lizongying/go-xpath)

## Install

```
go get github.com/lizongying/go-xpath
```

## Usage

For more usage, please refer to the test
[xpath_test](./test/xpath_test.go)

```go
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