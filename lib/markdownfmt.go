// +build js

//go:generate gopherjs build -m markdownfmt.go

package main

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/shurcooL/markdownfmt/markdown"
)

func ProcessMarkdown(text string) string {
	output, err := markdown.Process("", []byte(text), nil)
	if err != nil {
		println("ProcessMarkdown:", err.Error())
		return text
	}
	return string(output)
}

func main() {
	js.Module.Get("exports").Set("ProcessMarkdown", ProcessMarkdown)
}
