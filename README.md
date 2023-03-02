# go-xpath

简单的xpath选择器 simple xpath selector

[document](https://pkg.go.dev/github.com/lizongying/go-xpath)

## Install

```
go get github.com/lizongying/go-xpath
```

## Usage

For more usage, please refer to the test
[selector_test](./selector/selector_test.go)

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-xpath/selector"
)

func main() {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := selector.NewSelectorFromStr(html)
	fmt.Println(x)
}
```