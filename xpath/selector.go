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

// One return a result
func (s *Selector) One(path string) (result *Result) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, path)
	if n == nil {
		return
	}
	str := htmlquery.InnerText(n)
	result = NewResult(str)
	return
}

// Many return a result array
func (s *Selector) Many(path string) (results []*Result) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, path)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		results = append(results, NewResult(str))
	}
	return
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

func (s *Selector) ManySelector(path string) (selectors []*Selector) {
	return s.FindNodeMany(path)
}

// FindNodeMany find nodes
func (s *Selector) FindNodeMany(path string) (selectors []*Selector) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, path)
	for _, n := range ns {
		selectors = append(selectors, &Selector{
			node: n,
		})
	}
	return
}

func (s *Selector) OneSelector(path string) (selector *Selector) {
	return s.FindNodeOne(path)
}

// FindNodeOne find node
func (s *Selector) FindNodeOne(path string) (selector *Selector) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, path)
	if n == nil {
		return
	}
	selector = &Selector{
		node: n,
	}
	return
}

// FindNodeOneOr find node
func (s *Selector) FindNodeOneOr(path string) (selector *Selector) {
	if s == nil {
		selector = &Selector{}
		return
	}
	if s.node == nil {
		selector = &Selector{}
		return
	}
	n := htmlquery.FindOne(s.node, path)
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
func (s *Selector) FindStrMany(path string) (list []string) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, path)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		str = strings.TrimSpace(str)
		list = append(list, str)
	}
	return
}

// FindStrOne find a string
func (s *Selector) FindStrOne(path string) (str string) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, path)
	if n == nil {
		return
	}
	str = htmlquery.InnerText(n)
	str = strings.TrimSpace(str)
	return
}

// FindStrOneOr find a string, will return a default string if you find nothing
func (s *Selector) FindStrOneOr(path string, or string) (str string) {
	if s == nil {
		str = or
		return
	}
	if s.node == nil {
		str = or
		return
	}
	n := htmlquery.FindOne(s.node, path)
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
func (s *Selector) FindIntMany(path string) (list []int) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	ns := htmlquery.Find(s.node, path)
	for _, n := range ns {
		str := htmlquery.InnerText(n)
		str = strings.TrimSpace(str)
		i, _ := strconv.Atoi(str)
		list = append(list, i)
	}
	return
}

// FindIntOne find int
func (s *Selector) FindIntOne(path string) (i int) {
	if s == nil {
		return
	}
	if s.node == nil {
		return
	}
	n := htmlquery.FindOne(s.node, path)
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
func (s *Selector) FindIntOneOr(path string, or int) (i int) {
	if s == nil {
		i = or
		return
	}
	if s.node == nil {
		i = or
		return
	}
	n := htmlquery.FindOne(s.node, path)
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

func (s *Selector) String() (str string) {
	return s.OutHtml(true)
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
