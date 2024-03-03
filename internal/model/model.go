package model

import (
	"fmt"
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

	page := []byte(fmt.Sprintf("<!DOCTYPE html><html lang=\"en\"> <head> <meta charset=\"UTF-8\"> <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"> <title>Заголовок вашей страницы</title> </head> <body> <header> <h1>%s</h1> </header> </body> </html>", p.Title))
	page = append(page, markdown.Render(doc, renderer)...)

	return page
}