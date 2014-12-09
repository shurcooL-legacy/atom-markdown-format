// +build js

//go:generate gopherjs build -m markdownfmt.go

package main

import (
	"github.com/gopherjs/gopherjs/js"

	"gopkg.in/yaml.v2"

	"github.com/shurcooL/markdownfmt/markdown"
)

func ProcessMarkdown(text string) string {
	err := yaml.Unmarshal([]byte(text), &map[string]interface{}{})
	if err == nil {
		return text
	}

	output := []byte("")
	output, err = markdown.Process("", []byte(text), nil)
	if err != nil {
		println("ProcessMarkdown:", err.Error())
		return text
	}

	return string(output)
}

func main() {
	js.Module.Get("exports").Set("ProcessMarkdown", ProcessMarkdown)
}
