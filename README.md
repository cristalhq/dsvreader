# dsvreader

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![version-img]][version-url]

Fast reader for delimiter-separated data in Go.

## Features

* Supports CSV, TSV, PSV.
* Supports user defined delimiter.
* Dependency-free.
* Optimized for speed.
* Based on [Aliaksandr Valialkin's TSVReader](https://github.com/valyala/tsvreader)

See [docs][pkg-url].

## Install

Go version 1.17+

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
    panic("unexpected error: %s", err)
}
```

See examples: [example_test.go](example_test.go).

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/dsvreader/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/dsvreader/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/dsvreader
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/dsvreader
[version-img]: https://img.shields.io/github/v/release/cristalhq/dsvreader
[version-url]: https://github.com/cristalhq/dsvreader/releases
