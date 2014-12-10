// +build js

//go:generate gopherjs build -m markdownfmt.go

package main

import (
	"github.com/gopherjs/gopherjs/js"

	"regexp"

	"github.com/shurcooL/markdownfmt/markdown"
)

var fmExp = fmRegexp{regexp.MustCompile(`^(?P<header>---\n[\S\s]*\n---\n)(?P<content>[\S\s]*)$`)}

type fmRegexp struct {
	*regexp.Regexp
}

func (r *fmRegexp) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}

		captures[name] = match[i]
	}
	return captures
}

func ProcessMarkdownText(text string) string {
	output, err := markdown.Process("", []byte(text), nil)
	if err != nil {
		println("ProcessMarkdown:", err.Error())
		return text
	}
	return string(output)
}

func ProcessMarkdown(text string) string {
	frontmatterMatches := fmExp.FindStringSubmatchMap(text)
	if len(frontmatterMatches) == 0 {
		return ProcessMarkdownText(text)
	}

	return frontmatterMatches["header"] + ProcessMarkdownText(frontmatterMatches["content"])
}

func main() {
	js.Module.Get("exports").Set("ProcessMarkdown", ProcessMarkdown)
}
