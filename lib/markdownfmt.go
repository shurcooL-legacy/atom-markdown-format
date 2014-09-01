// +build js

package main

import (
	"github.com/gopherjs/gopherjs/js"

	"github.com/shurcooL/markdownfmt/markdown"
)

/*
Rebuild steps:

`gopherjs build -m markdownfmt.go`
*/

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
