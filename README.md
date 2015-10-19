### Deprecation Notice

Unfortunately, I've made the difficult decision to deprecate this package. My initial plan was to make `markdownfmt` essentially perfect, the way `gofmt` is, so that it could be always run-on-save in every possible situation, without needing for configuration or to ever turn it off.

Along the way, I learned there are many obstacles that prevent from that being possible with Markdown (without doing an exorbitant amount of work). For one, there are ambiguities in the original spec. CommonMark tries to fix that, but it is too complex IMO. People have their personal preferences that they want to preserve because Markdown allows multiple ways to do the same thing. Some blog engines try to put YAML headers in files with `.md` extension, so parsing a Markdown file cannot be done reliably without "guessing" if it has a YAML header or not.

I developed this package in my free time. With my very limited resources, I can't develop and maintain it further. I don't want to offer a package that provides a sub-par user experience. After 2 weeks (on November 1, 2015), my plan is to remove the package from [Atom package directory](https://atom.io/packages/markdown-format) and move this repository to <https://github.com/shurcooL-legacy>.

Markdown Format Package
=======================

Formats your Markdown text (to look like the rendered HTML) on save.

![Markdown Format Demo](https://github.com/shurcooL/atom-markdown-format/blob/master/Demo.gif?raw=true)
