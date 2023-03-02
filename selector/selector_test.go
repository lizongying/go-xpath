package selector

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// go test -v selector/*.go

// TestSelectorFromStr go test -v selector/*.go -run TestSelectorFromStr
func TestSelectorFromStr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	assert.Equal(t, x.GetNode().FirstChild.Data, "html")
}

// TestSelectorFromReader go test -v selector/*.go -run TestSelectorFromReader
func TestSelectorFromReader(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromReader(strings.NewReader(html))

	assert.Equal(t, x.GetNode().FirstChild.Data, "html")
}

// TestSelectorFromFile go test -v selector/*.go -run TestSelectorFromFile
func TestSelectorFromFile(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	file, _ := ioutil.TempFile(os.TempDir(), "")
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Log(err)
		}
	}(file.Name())
	_, _ = file.WriteString(html)
	x, _ := NewSelectorFromFile(file.Name())

	assert.Equal(t, x.GetNode().FirstChild.Data, "html")
}

// TestSelectorFindStrMany go test -v selector/*.go -run TestSelectorFindStrMany
func TestSelectorFindStrMany(t *testing.T) {
	html := []byte(`<html class="123">....<div class="789">....</div><div class="456">....</div></html>`)
	x, _ := NewSelectorFromBytes(html)

	li := x.FindStrMany(`//div[@class]/@class`)
	t.Log(li)
	assert.Equal(t, li, []string{"789", "456"})

	li = x.FindStrMany(`//div[@class]/@classs`)
	t.Log(li)
	assert.Equal(t, li, []string(nil))
}

// TestSelectorFindStrOne go test -v selector/*.go -run TestSelectorFindStrOne
func TestSelectorFindStrOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	str := x.FindStrOne(`//div[@class]/@class`)
	t.Log(str)
	assert.Equal(t, str, "789")

	str = x.FindStrOne(`//div[@class]/@classs`)
	t.Log(str)
	assert.Equal(t, str, "")
}

// TestSelectorFindStrOneOr go test -v selector/*.go -run TestSelectorFindStrOneOr
func TestSelectorFindStrOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	str := x.FindStrOneOr(`//div[@class]/@class`, "999")
	t.Log(str)
	assert.Equal(t, str, "789")

	str = x.FindStrOneOr(`//div[@class]/@classs`, "999")
	t.Log(str)
	assert.Equal(t, str, "999")
}

// TestSelectorFindIntMany go test -v selector/*.go -run TestSelectorFindIntMany
func TestSelectorFindIntMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	li := x.FindIntMany(`//div[@class]/@class`)
	t.Log(li)
	assert.Equal(t, li, []int{789, 456})

	li = x.FindIntMany(`//div[@class]/@classs`)
	t.Log(li)
	assert.Equal(t, li, []int(nil))
}

// TestSelectorFindIntOne go test -v selector/*.go -run TestSelectorFindIntOne
func TestSelectorFindIntOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	i := x.FindIntOne(`//div[@class]/@class`)
	t.Log(i)
	assert.Equal(t, i, 789)

	i = x.FindIntOne(`//div[@class]/@classs`)
	t.Log(i)
	assert.Equal(t, i, 0)
}

// TestSelectorFindIntOneOr go test -v selector/*.go -run TestSelectorFindIntOneOr
func TestSelectorFindIntOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	i := x.FindIntOneOr(`//div[@class]/@class`, 999)
	t.Log(i)
	assert.Equal(t, i, 789)

	i = x.FindIntOneOr(`//div[@class]/@classs`, 999)
	t.Log(i)
	assert.Equal(t, i, 999)
}

// TestSelectorFindNodeManyFindStrMany go test -v selector/*.go -run TestSelectorFindNodeManyFindStrMany
func TestSelectorFindNodeManyFindStrMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	for k, i := range x.FindNodeMany(`//div[@class]`) {
		li := i.FindStrMany(`./@class`)
		t.Log(li)
		if k == 0 {
			assert.Equal(t, li, []string{"789"})
		}
		if k == 1 {
			assert.Equal(t, li, []string{"456"})
		}
	}

	for _, i := range x.FindNodeMany(`//div[@classs]`) {
		li := i.FindStrMany(`./@class`)
		t.Log(li)
		assert.Equal(t, li, []string(nil))
	}
}

// TestSelectorFindNodeOneFindStrMany go test -v selector/*.go -run TestSelectorFindNodeOneFindStrMany
func TestSelectorFindNodeOneFindStrMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	li := x.FindNodeOne(`//div[@class]`).FindStrMany(`./@class`)
	t.Log(li)
	assert.Equal(t, li, []string{"789"})

	li = x.FindNodeOneOr(`//div[@classs]`).FindStrMany(`./@class`)
	t.Log(li)
	assert.Equal(t, li, []string(nil))
}

// TestSelectorOutHtml go test -v selector/*.go -run TestSelectorOutHtml
func TestSelectorOutHtml(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := NewSelectorFromStr(html)

	s := x.FindNodeOne(`//html[@class]`).OutHtml(true)
	t.Log(s)

	assert.Equal(t, s, "<html class=\"123\"><head></head><body>....<div class=\"789\">....</div><div class=\"456\">....</div></body></html>")

	s = x.FindNodeOne(`//html[@class]`).OutHtml(false)
	t.Log(s)

	assert.Equal(t, s, "<head></head><body>....<div class=\"789\">....</div><div class=\"456\">....</div></body>")
}
