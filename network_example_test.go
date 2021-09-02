// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/tunabay/go-typeio"
)

func ExampleReadIPv4() {
	b, _ := hex.DecodeString("c0000201cb00710fc63364ff")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadIPv4(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 192.0.2.1
	// 203.0.113.15
	// 198.51.100.255
}

func ExampleWriteIPv4() {
	w := new(bytes.Buffer)

	data := []net.IP{
		net.ParseIP("192.0.2.1"),
		net.ParseIP("203.0.113.15"),
		net.ParseIP("198.51.100.255"),
	}
	for _, v := range data {
		if err := typeio.WriteIPv4(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// c0000201cb00710fc63364ff
}

func ExampleReadIPv6() {
	b1 := "20010db8000000000000000012345678"
	b2 := "00000000000000000000ffffc0000280"
	b, _ := hex.DecodeString(b1 + b2)
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadIPv6(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 2001:db8::1234:5678
	// 192.0.2.128
}

func ExampleWriteIPv6() {
	w := new(bytes.Buffer)

	data := []net.IP{
		net.ParseIP("2001:db8::1234:5678"),
		net.ParseIP("::ffff:192.0.2.128"), // IPv4-mapped
	}
	for _, v := range data {
		if err := typeio.WriteIPv6(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 20010db800000000000000001234567800000000000000000000ffffc0000280
}
