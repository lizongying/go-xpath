package xpath

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

type Xpath struct {
	node *html.Node
}

// GetNode get node
func (x *Xpath) GetNode() (node *html.Node) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	node = x.node
	return
}

// FindNodeMany find nodes
func (x *Xpath) FindNodeMany(str string) (xs []*Xpath) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	ns := htmlquery.Find(x.node, str)
	for _, n := range ns {
		xs = append(xs, &Xpath{
			node: n,
		})
	}
	return
}

// FindNodeOne find node
func (x *Xpath) FindNodeOne(str string) (x1 *Xpath) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	n := htmlquery.FindOne(x.node, str)
	if n == nil {
		return
	}
	x1 = &Xpath{
		node: n,
	}
	return
}

// FindNodeOneOr find node
func (x *Xpath) FindNodeOneOr(str string) (x1 *Xpath) {
	if x == nil {
		x1 = &Xpath{}
		return
	}
	if x.node == nil {
		x1 = &Xpath{}
		return
	}
	n := htmlquery.FindOne(x.node, str)
	if n == nil {
		x1 = &Xpath{}
		return
	}
	x1 = &Xpath{
		node: n,
	}
	return
}

// FindStrMany find a string list
func (x *Xpath) FindStrMany(xpath string) (list []string) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	ns := htmlquery.Find(x.node, xpath)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		str = strings.TrimSpace(str)
		list = append(list, str)
	}
	return
}

// FindStrOne find a string
func (x *Xpath) FindStrOne(xpath string) (str string) {
	if x == nil {
		return
	}
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	n := htmlquery.FindOne(x.node, xpath)
	if n == nil {
		return
	}
	str = htmlquery.InnerText(n)
	str = strings.TrimSpace(str)
	return
}

// FindStrOneOr find a string, will return a default string if you find nothing
func (x *Xpath) FindStrOneOr(xpath string, or string) (str string) {
	if x == nil {
		str = or
		return
	}
	if x.node == nil {
		str = or
		return
	}
	n := htmlquery.FindOne(x.node, xpath)
	if n == nil {
		str = or
		return
	}
	str = htmlquery.InnerText(n)
	str = strings.TrimSpace(str)
	if str != "" {
		return
	}
	str = or
	return
}

// FindIntMany find int list
func (x *Xpath) FindIntMany(xpath string) (list []int) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	ns := htmlquery.Find(x.node, xpath)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		str = strings.TrimSpace(str)
		i, _ := strconv.Atoi(str)
		list = append(list, i)
	}
	return
}

// FindIntOne find int
func (x *Xpath) FindIntOne(xpath string) (i int) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	n := htmlquery.FindOne(x.node, xpath)
	if n == nil {
		return
	}
	str := htmlquery.InnerText(n)
	str = strings.TrimSpace(str)
	if str != "" {
		i, _ = strconv.Atoi(str)
		return
	}
	return
}

// FindIntOneOr find  int, will return a default int if you find nothing
func (x *Xpath) FindIntOneOr(xpath string, or int) (i int) {
	if x == nil {
		i = or
		return
	}
	if x.node == nil {
		i = or
		return
	}
	n := htmlquery.FindOne(x.node, xpath)
	if n == nil {
		i = or
		return
	}
	str := htmlquery.InnerText(n)
	str = strings.TrimSpace(str)
	if str != "" {
		i, _ = strconv.Atoi(str)
		return
	}
	i = or
	return
}

// NewXpathFromStr xpath init
func NewXpathFromStr(str string) (xpath *Xpath, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.Parse(strings.NewReader(str))
	if err != nil {
		return
	}
	xpath = &Xpath{
		node: node,
	}
	return
}

// NewXpathFromBytes xpath init
func NewXpathFromBytes(b []byte) (xpath *Xpath, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.Parse(bytes.NewReader(b))
	if err != nil {
		return
	}
	xpath = &Xpath{
		node: node,
	}
	return
}

// NewXpathFromReader xpath init
func NewXpathFromReader(i io.Reader) (xpath *Xpath, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.Parse(i)
	if err != nil {
		return
	}
	xpath = &Xpath{
		node: node,
	}
	return
}

// NewXpathFromFile xpath init
func NewXpathFromFile(f string) (xpath *Xpath, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.LoadDoc(f)
	if err != nil {
		return
	}
	xpath = &Xpath{
		node: node,
	}
	return
}

// OutHtml return html
func (x *Xpath) OutHtml(self bool) (str string) {
	if x == nil {
		return
	}
	if x.node == nil {
		return
	}
	str = htmlquery.OutputHTML(x.node, self)
	return
}
