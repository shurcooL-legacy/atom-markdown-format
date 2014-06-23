// +build ignore

package main

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/russross/blackfriday"
	"github.com/shurcooL/markdownfmt/markdown"
)

/*
Rebuild steps:

`gopherjs build -m markdownfmt.go`
*/

func ProcessMarkdown(text string) string {
	// GitHub Flavored Markdown-like extensions.
	extensions := 0
	extensions |= blackfriday.EXTENSION_NO_INTRA_EMPHASIS
	extensions |= blackfriday.EXTENSION_TABLES
	extensions |= blackfriday.EXTENSION_FENCED_CODE
	extensions |= blackfriday.EXTENSION_AUTOLINK
	extensions |= blackfriday.EXTENSION_STRIKETHROUGH
	extensions |= blackfriday.EXTENSION_SPACE_HEADERS
	//extensions |= blackfriday.EXTENSION_HARD_LINE_BREAK

	output := blackfriday.Markdown([]byte(text), markdown.NewRenderer(), extensions)
	return string(output)
}

func main() {
	js.Module.Get("exports").Set("ProcessMarkdown", ProcessMarkdown)
}
