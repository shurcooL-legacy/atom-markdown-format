Markdown Format Package
=======================

It runs `markdownfmt` on save, which formats your Markdown text to look like the HTML.

See [installation instructions](#installation).

![Markdown Format Demo](https://github.com/shurcooL/atom-markdown-format/blob/master/Demo.gif?raw=true)

Installation
------------

In Settings View, point to the path of the `markdownfmt` binary, which you can get at [github.com/shurcooL/markdownfmt](https://github.com/shurcooL/markdownfmt).

### On Mac OS X

Using [Homebrew](http://brew.sh/):

```bash
brew install go mercurial
mkdir ~/gocode
```

Add to `~/.bash_profile` or `~/.profile`:

```bash
export GOPATH="$HOME/gocode"
export PATH="$PATH:$GOPATH/bin"
```

Get `markdownfmt`:

```bash
go get -u github.com/shurcooL/markdownfmt
```
