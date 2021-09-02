// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio_test

import (
	"fmt"
	"io"
	"math"
	"net"
	"time"

	"github.com/tunabay/go-typeio"
)

func Example_usage() {
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
	fmt.Println(u32)

	f64, _ := typeio.ReadFloat64BE(r)
	fmt.Println(f64)

	tm, _ := typeio.ReadUnixTimeUTC32BE(r)
	fmt.Println(tm)

	str, _, _ := typeio.ReadCString(r)
	fmt.Println(str)

	ipv4, _ := typeio.ReadIPv4(r)
	fmt.Println(ipv4)

	// Output:
	// 12345678
	// 3.141592653589793
	// 2012-05-20 23:53:54 +0000 UTC
	// null-terminated string.
	// 192.0.2.1
}
