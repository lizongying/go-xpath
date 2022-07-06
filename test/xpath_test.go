package test

import (
	"github.com/lizongying/go-xpath/xpath"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

// go test -v test/xpath_test.go

// TestXpathFromStrParse go test -v test/xpath_test.go -run TestXpathFromStrParse
func TestXpathFromStrParse(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	assert.Equal(t, x.Node.FirstChild.Data, "html")
}

// TestXpathFromReaderParse go test -v test/xpath_test.go -run TestXpathFromReaderParse
func TestXpathFromReaderParse(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromReader(strings.NewReader(html))

	assert.Equal(t, x.Node.FirstChild.Data, "html")
}

// TestXpathFromFile go test -v test/xpath_test.go -run TestXpathFromFile
func TestXpathFromFile(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	file, _ := ioutil.TempFile(os.TempDir(), "")
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			t.Log(err)
		}
	}(file.Name())
	_, _ = file.WriteString(html)
	x, _ := xpath.NewXpathFromFile(file.Name())

	assert.Equal(t, x.Node.FirstChild.Data, "html")
}

// TestXpathFindStrMany go test -v test/xpath_test.go -run TestXpathFindStrMany
func TestXpathFindStrMany(t *testing.T) {
	html := []byte(`<html class="123">....<div class="789">....</div><div class="456">....</div></html>`)
	x, _ := xpath.NewXpathFromBytes(html)

	li := x.FindStrMany(`//div[@class]/@class`)
	t.Log(li)
	assert.Equal(t, li, []string{"789", "456"})

	li = x.FindStrMany(`//div[@class]/@classs`)
	t.Log(li)
	assert.Equal(t, li, []string(nil))
}

// TestXpathFindStrOne go test -v test/xpath_test.go -run TestXpathFindStrOne
func TestXpathFindStrOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	str := x.FindStrOne(`//div[@class]/@class`)
	t.Log(str)
	assert.Equal(t, str, "789")

	str = x.FindStrOne(`//div[@class]/@classs`)
	t.Log(str)
	assert.Equal(t, str, "")
}

// TestXpathFindStrOneOr go test -v test/xpath_test.go -run TestXpathFindStrOneOr
func TestXpathFindStrOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	str := x.FindStrOneOr(`//div[@class]/@class`, "999")
	t.Log(str)
	assert.Equal(t, str, "789")

	str = x.FindStrOneOr(`//div[@class]/@classs`, "999")
	t.Log(str)
	assert.Equal(t, str, "999")
}

// TestXpathFindIntMany go test -v test/xpath_test.go -run TestXpathFindIntMany
func TestXpathFindIntMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	li := x.FindIntMany(`//div[@class]/@class`)
	t.Log(li)
	assert.Equal(t, li, []int{789, 456})

	li = x.FindIntMany(`//div[@class]/@classs`)
	t.Log(li)
	assert.Equal(t, li, []int(nil))
}

// TestXpathFindIntOne go test -v test/xpath_test.go -run TestXpathFindIntOne
func TestXpathFindIntOne(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	i := x.FindIntOne(`//div[@class]/@class`)
	t.Log(i)
	assert.Equal(t, i, 789)

	i = x.FindIntOne(`//div[@class]/@classs`)
	t.Log(i)
	assert.Equal(t, i, 0)
}

// TestXpathFindIntOneOr go test -v test/xpath_test.go -run TestXpathFindIntOneOr
func TestXpathFindIntOneOr(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	i := x.FindIntOneOr(`//div[@class]/@class`, 999)
	t.Log(i)
	assert.Equal(t, i, 789)

	i = x.FindIntOneOr(`//div[@class]/@classs`, 999)
	t.Log(i)
	assert.Equal(t, i, 999)
}

// TestXpathFindNodeManyFindStrMany go test -v test/xpath_test.go -run TestXpathFindNodeManyFindStrMany
func TestXpathFindNodeManyFindStrMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

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

// TestXpathFindNodeOneFindStrMany go test -v test/xpath_test.go -run TestXpathFindNodeOneFindStrMany
func TestXpathFindNodeOneFindStrMany(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	li := x.FindNodeOne(`//div[@class]`).FindStrMany(`./@class`)
	t.Log(li)
	assert.Equal(t, li, []string{"789"})

	li = x.FindNodeOneOr(`//div[@classs]`).FindStrMany(`./@class`)
	t.Log(li)
	assert.Equal(t, li, []string(nil))
}

// TestXpathOutHtml go test -v test/xpath_test.go -run TestXpathOutHtml
func TestXpathOutHtml(t *testing.T) {
	html := `<html class="123">....<div class="789">....</div><div class="456">....</div></html>`
	x, _ := xpath.NewXpathFromStr(html)

	s := x.FindNodeOne(`//html[@class]`).OutHtml(true)
	t.Log(s)

	assert.Equal(t, s, "<html class=\"123\"><head></head><body>....<div class=\"789\">....</div><div class=\"456\">....</div></body></html>")

	s = x.FindNodeOne(`//html[@class]`).OutHtml(false)
	t.Log(s)

	assert.Equal(t, s, "<head></head><body>....<div class=\"789\">....</div><div class=\"456\">....</div></body>")
}
