# jwt

[![Build Status][travis-img]][travis-url]
[![GoDoc][doc-img]][doc-url]
[![Go Report Card][reportcard-img]][reportcard-url]
[![Go Report Card][coverage-img]][coverage-url]

Fast reader for delimiter-separated data

## Features

* Supports CSV, TSV, PSV.
* Supports user defined delimiter.
* Dependency-free.
* Optimized for speed.
* Based on [Aliaksandr Valialkin's TSVReader](https://github.com/valyala/tsvreader)

## Install

```
go get github.com/cristalhq/dsvreader
```

## Example

```go
bs := bytes.NewBufferString(
  `foo\t42\n
  bar\t123\n`)

r := dsvreader.NewTSV(bs)
for r.Next() {
  col1 := r.String()
  col2 := r.Int()
  fmt.Printf("col1=%s, col2=%d\n", col1, col2)
}

if err := r.Error(); err != nil {
  fmt.Printf("unexpected error: %s", err)
}
```

## Documentation

See [these docs](https://godoc.org/github.com/cristalhq/dsvreader).

## License

[MIT License](LICENSE).

[travis-img]: https://travis-ci.org/cristalhq/dsvreader.svg?branch=master
[travis-url]: https://travis-ci.org/cristalhq/dsvreader
[doc-img]: https://godoc.org/github.com/cristalhq/dsvreader?status.svg
[doc-url]: https://godoc.org/github.com/cristalhq/dsvreader
[reportcard-img]: https://goreportcard.com/badge/cristalhq/dsvreader
[reportcard-url]: https://goreportcard.com/report/cristalhq/dsvreader
[coverage-img]: https://coveralls.io/repos/github/cristalhq/dsvreader/badge.svg?branch=master
[coverage-url]: https://coveralls.io/github/cristalhq/dsvreader?branch=master
