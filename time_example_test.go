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
	"time"

	"github.com/tunabay/go-typeio"
)

func ExampleReadUnixTimeUTC32BE() {
	b, _ := hex.DecodeString("4fb984124b50143b")
	r := bytes.NewReader(b)

	for {
		t, err := typeio.ReadUnixTimeUTC32BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(t)
	}

	// Output:
	// 2012-05-20 23:53:54 +0000 UTC
	// 2010-01-15 07:07:39 +0000 UTC
}

func ExampleReadUnixTime32BE() {
	b, _ := hex.DecodeString("4fb984124b50143b")
	r := bytes.NewReader(b)

	time.Local, _ = time.LoadLocation("Asia/Tokyo") // to fix the output
	for {
		t, err := typeio.ReadUnixTime32BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(t)
	}

	// Output:
	// 2012-05-21 08:53:54 +0900 JST
	// 2010-01-15 16:07:39 +0900 JST
}

func ExampleWriteUnixTime32BE() {
	w := new(bytes.Buffer)

	locJST, _ := time.LoadLocation("Asia/Tokyo")
	data := []time.Time{
		time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC),
		time.Date(2012, 5, 21, 8, 53, 54, 0, locJST),
		time.Date(2010, 1, 15, 7, 7, 39, 0, time.UTC),
	}
	for _, t := range data {
		if err := typeio.WriteUnixTime32BE(w, t); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 4fb984124fb984124b50143b
}

func ExampleReadUnixTimeUTC32LE() {
	b, _ := hex.DecodeString("1284b94f3b14504b")
	r := bytes.NewReader(b)

	for {
		t, err := typeio.ReadUnixTimeUTC32LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(t)
	}

	// Output:
	// 2012-05-20 23:53:54 +0000 UTC
	// 2010-01-15 07:07:39 +0000 UTC
}

func ExampleReadUnixTime32LE() {
	b, _ := hex.DecodeString("1284b94f3b14504b")
	r := bytes.NewReader(b)

	time.Local, _ = time.LoadLocation("Asia/Tokyo") // to fix the output
	for {
		t, err := typeio.ReadUnixTime32LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(t)
	}

	// Output:
	// 2012-05-21 08:53:54 +0900 JST
	// 2010-01-15 16:07:39 +0900 JST
}

func ExampleWriteUnixTime32LE() {
	w := new(bytes.Buffer)

	locJST, _ := time.LoadLocation("Asia/Tokyo")
	data := []time.Time{
		time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC),
		time.Date(2012, 5, 21, 8, 53, 54, 0, locJST),
		time.Date(2010, 1, 15, 7, 7, 39, 0, time.UTC),
	}
	for _, t := range data {
		if err := typeio.WriteUnixTime32LE(w, t); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 1284b94f1284b94f3b14504b
}
