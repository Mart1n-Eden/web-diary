package model

import (
	"os"
	"strings"
	"github.com/gomarkdown/markdown"
	_ "github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Article struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  []byte `json:"body"`
}

type PageData struct {
	Articles   []Article
	Pagination Pagination
}

type Pagination struct {
	Page     int
	PrevPage int
	NextPage int
	HasPrev  bool
	HasNext  bool
}

func (p *Article) Gluing() []byte {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	par := parser.NewWithExtensions(extensions)
	doc := par.Parse(p.Body)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	page, err := os.ReadFile("./view/article.html")
	if err != nil {
		panic(err)
	}

	page = []byte(strings.ReplaceAll(string(page), "{{.Title}}", p.Title))
	body := "<pre>" + string(markdown.Render(doc, renderer)) + "</pre>"
	page = []byte(strings.ReplaceAll(string(page), "<!-- article -->", body))

	return page
}