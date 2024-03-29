# go-ext2mime

Package **ext2mime** provides tools inferring what MIME type is associated with a file extension, for the Go programming language.

For example:

* `.css` -> `text/css`
* `.epub` -> `application/epub+zip`,
* `.html` -> `text/html`,
* `.jpeg` -> `image/jpeg`
* `.js` -> `text/javascript`,
* `.json` -> `application/json`,
* `.md` -> `text/markdown`,
* `.png` -> `image/png`
* `.qoi` -> `image/qoi`
* etc etc etc.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-ext2mime

[![GoDoc](https://godoc.org/github.com/reiver/go-ext2mime?status.svg)](https://godoc.org/github.com/reiver/go-ext2mime)

## Example

Here is a simple example:

```golang
import "github.com/reiver/go-ext2mime"

// ...

fileExtension := path.Ext(filename)

mimeType := ext2mime.Get(fileExtension)
```

You can also add you own file-extension to MIME-type mapping with:
```golang
func init() {
	const fileExtension string = ".abc"
	const mimeType string = "application/abc"

	ext2mime.Set(fileExtension, mimeType)
}
```


## Import

To import package **ext2mime** use `import` code like the follownig:
```
import "github.com/reiver/go-ext2mime"
```

## Installation

To install package **ext2mime** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-ext2mime
```

## Author

Package **ext2mime** was written by [Charles Iliya Krempeaux](http://reiver.link)
