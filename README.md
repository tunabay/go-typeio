# go-typeio

[![Go Reference](https://pkg.go.dev/badge/github.com/tunabay/go-typeio.svg)](https://pkg.go.dev/github.com/tunabay/go-typeio)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)

## Overview

Package typeio provides handy functions for reading and writing primitive data
types over io.Reader and io.Writer.

## Usage

```
import "github.com/tunabay/go-typeio"

func main() {
	r, w := io.Pipe()

	go func() {
		_ = typeio.WriteUint32BE(w, 12345678)
		_ = typeio.WriteFloat64BE(w, math.Pi)
		_ = typeio.WriteUnixTime32BE(w, time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC))
		_, _ = w.Write([]byte("null-terminated string.\x00"))
		_ = typeio.WriteIPv4(w, net.ParseIP("192.0.2.1"))
		w.Close()
	}()

	u32, _ := typeio.ReadUint32BE(r)
	fmt.Println(u32) // 12345678

	f64, _ := typeio.ReadFloat64BE(r)
	fmt.Println(f64) // 3.141592653589793

	tm, _ := typeio.ReadUnixTimeUTC32BE(r)
	fmt.Println(tm) // 2012-05-20 23:53:54 +0000 UTC

	str, _, _ := typeio.ReadCString(r)
	fmt.Println(str) // null-terminated string.

	ipv4, _ := typeio.ReadIPv4(r)
	fmt.Println(ipv4) // 192.0.2.1
}
```
[Run in Go Playground](https://play.golang.org/p/N2uNvQxOegz)

## Documentation and examples:

- Read the [documentation](https://pkg.go.dev/github.com/tunabay/go-typeio).

## License

go-typeio is available under the MIT license. See the [LICENSE](LICENSE) file for more information.
