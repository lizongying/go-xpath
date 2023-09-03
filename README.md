# go-xpath

golang实现的简单xpath选择器，使用更简单。 simple xpath selector

[go-xpath](https://github.com/lizongying/go-xpath)

[document](https://pkg.go.dev/github.com/lizongying/go-xpath)

## Install

```
go get github.com/lizongying/go-xpath
```

## Usage

For more usage, please refer to the test
[selector_test](./xpath/selector_test.go)

```go
package main

import (
	"fmt"
	"github.com/lizongying/go-xpath/xpath"
)

func main() {
	html := `<html class="abc">....<div class="def">....</div><div class="gkl">123</div></html>`
	x, _ := xpath.NewSelectorFromStr(html)

	s := x.FindStrOne(`//div[@class="def"]/text()`)
	//....
	fmt.Println(s)

	i := x.FindIntOneOr(`//div[@class="gkl"]/text()`, 111)
	//123
	fmt.Println(i)

	i = x.FindIntOneOr(`//div[@class="mn"]/text()`, 111)
	//111
	fmt.Println(i)

	sl := x.FindStrMany(`//div/text()`)
	//[.... 123]
	fmt.Println(sl)

	s = x.FindNodeOne(`//html[@class="abc"]`).FindStrOne(`.//div[@class="gkl"]`)
	//123
	fmt.Println(s)
}

```