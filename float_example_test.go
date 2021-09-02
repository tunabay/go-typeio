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

func ExampleReadFloat32BE() {
	b, _ := hex.DecodeString("40490fdbbdaaaaab")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadFloat32BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 3.1415927
	// -0.083333336
}

func ExampleWriteFloat32BE() {
	w := new(bytes.Buffer)

	data := []float32{
		3.141592653589793,
		-.083333333333333,
	}
	for _, t := range data {
		if err := typeio.WriteFloat32BE(w, t); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 40490fdbbdaaaaab
}

func ExampleReadFloat32LE() {
	b, _ := hex.DecodeString("db0f4940abaaaabd")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadFloat32LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 3.1415927
	// -0.083333336
}

func ExampleWriteFloat32LE() {
	w := new(bytes.Buffer)

	data := []float32{
		3.141592653589793,
		-.083333333333333,
	}
	for _, t := range data {
		if err := typeio.WriteFloat32LE(w, t); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// db0f4940abaaaabd
}

func ExampleReadFloat64BE() {
	b, _ := hex.DecodeString("400921fb54442d18bfb5555555555555")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadFloat64BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 3.141592653589793
	// -0.08333333333333333
}

func ExampleWriteFloat64BE() {
	w := new(bytes.Buffer)

	data := []float64{
		3.14159265358979323846,
		-.08333333333333333333,
	}
	for _, t := range data {
		if err := typeio.WriteFloat64BE(w, t); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 400921fb54442d18bfb5555555555555
}

func ExampleReadFloat64LE() {
	b, _ := hex.DecodeString("182d4454fb210940555555555555b5bf")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadFloat64LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 3.141592653589793
	// -0.08333333333333333
}

func ExampleWriteFloat64LE() {
	w := new(bytes.Buffer)

	data := []float64{
		3.14159265358979323846,
		-.08333333333333333333,
	}
	for _, t := range data {
		if err := typeio.WriteFloat64LE(w, t); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 182d4454fb210940555555555555b5bf
}
