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
