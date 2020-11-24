# dsvreader

[![build-img]][build-url]
[![pkg-img]][pkg-url]
[![reportcard-img]][reportcard-url]
[![coverage-img]][coverage-url]

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
    panic("unexpected error: %s", err)
}
```

## Documentation

See [these docs][pkg-url].

## License

[MIT License](LICENSE).

[build-img]: https://github.com/cristalhq/dsvreader/workflows/build/badge.svg
[build-url]: https://github.com/cristalhq/dsvreader/actions
[pkg-img]: https://pkg.go.dev/badge/cristalhq/dsvreader
[pkg-url]: https://pkg.go.dev/github.com/cristalhq/dsvreader
[reportcard-img]: https://goreportcard.com/badge/cristalhq/dsvreader
[reportcard-url]: https://goreportcard.com/report/cristalhq/dsvreader
[coverage-img]: https://codecov.io/gh/cristalhq/dsvreader/branch/master/graph/badge.svg
[coverage-url]: https://codecov.io/gh/cristalhq/dsvreader
