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

	"github.com/tunabay/go-typeio"
)

func ExampleReadStringN() {
	b, _ := hex.DecodeString("48656c6c6f00e4b896e7958c")
	r := bytes.NewReader(b)

	for {
		s, err := typeio.ReadStringN(r, 6, 0)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(s)
	}

	// Output:
	// Hello
	// 世界
}

func ExampleReadCString() {
	b, _ := hex.DecodeString("48656c6c6f00e4b896e7958c")
	r := bytes.NewReader(b)

	for {
		s, n, err := typeio.ReadCString(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(n, s)
	}

	// Output:
	// 6 Hello
	// 6 世界
}
