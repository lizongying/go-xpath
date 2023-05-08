package xpath

import (
	"bytes"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io"
	"strconv"
	"strings"
)

type Selector struct {
	node *html.Node
}

// GetNode get node
func (s *Selector) GetNode() (node *html.Node) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	node = s.node
	return
}

// FindNodeMany find nodes
func (s *Selector) FindNodeMany(xpath string) (selectors []*Selector) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, xpath)
	for _, n := range ns {
		selectors = append(selectors, &Selector{
			node: n,
		})
	}
	return
}

// FindNodeOne find node
func (s *Selector) FindNodeOne(xpath string) (selector *Selector) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, xpath)
	if n == nil {
		return
	}
	selector = &Selector{
		node: n,
	}
	return
}

// FindNodeOneOr find node
func (s *Selector) FindNodeOneOr(xpath string) (selector *Selector) {
	if s == nil {
		selector = &Selector{}
		return
	}
	if s.node == nil {
		selector = &Selector{}
		return
	}
	n := htmlquery.FindOne(s.node, xpath)
	if n == nil {
		selector = &Selector{}
		return
	}
	selector = &Selector{
		node: n,
	}
	return
}

// FindStrMany find a string list
func (s *Selector) FindStrMany(xpath string) (list []string) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, xpath)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		str = strings.TrimSpace(str)
		list = append(list, str)
	}
	return
}

// FindStrOne find a string
func (s *Selector) FindStrOne(xpath string) (str string) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, xpath)
	if n == nil {
		return
	}
	str = htmlquery.InnerText(n)
	str = strings.TrimSpace(str)
	return
}

// FindStrOneOr find a string, will return a default string if you find nothing
func (s *Selector) FindStrOneOr(xpath string, or string) (str string) {
	if s == nil {
		str = or
		return
	}
	if s.node == nil {
		str = or
		return
	}
	n := htmlquery.FindOne(s.node, xpath)
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
func (s *Selector) FindIntMany(xpath string) (list []int) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, xpath)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		str = strings.TrimSpace(str)
		i, _ := strconv.Atoi(str)
		list = append(list, i)
	}
	return
}

// FindIntOne find int
func (s *Selector) FindIntOne(xpath string) (i int) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, xpath)
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
func (s *Selector) FindIntOneOr(xpath string, or int) (i int) {
	if s == nil {
		i = or
		return
	}
	if s.node == nil {
		i = or
		return
	}
	n := htmlquery.FindOne(s.node, xpath)
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

// NewSelectorFromStr selector from str
func NewSelectorFromStr(s string) (selector *Selector, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.Parse(strings.NewReader(s))
	if err != nil {
		return
	}
	selector = &Selector{
		node: node,
	}
	return
}

// NewSelectorFromBytes selector from bytes
func NewSelectorFromBytes(b []byte) (selector *Selector, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.Parse(bytes.NewReader(b))
	if err != nil {
		return
	}
	selector = &Selector{
		node: node,
	}
	return
}

// NewSelectorFromReader selector from reader
func NewSelectorFromReader(i io.Reader) (selector *Selector, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.Parse(i)
	if err != nil {
		return
	}
	selector = &Selector{
		node: node,
	}
	return
}

// NewSelectorFromFile selector from file
func NewSelectorFromFile(f string) (selector *Selector, err error) {
	htmlquery.DisableSelectorCache = true
	node, err := htmlquery.LoadDoc(f)
	if err != nil {
		return
	}
	selector = &Selector{
		node: node,
	}
	return
}

// OutHtml return html
func (s *Selector) OutHtml(self bool) (str string) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	str = htmlquery.OutputHTML(s.node, self)
	return
}
