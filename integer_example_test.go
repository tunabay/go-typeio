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

func ExampleReadUint8() {
	b, _ := hex.DecodeString("F07F80")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint8(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 240
	// 127
	// 128
}

func ExampleWriteUint8() {
	w := new(bytes.Buffer)

	data := []uint8{240, 127, 128}
	for _, v := range data {
		if err := typeio.WriteUint8(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// f07f80
}

func ExampleReadInt8() {
	b, _ := hex.DecodeString("F07F80")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt8(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// -16
	// 127
	// -128
}

func ExampleWriteInt8() {
	w := new(bytes.Buffer)

	data := []int8{-16, 127, -128}
	for _, v := range data {
		if err := typeio.WriteInt8(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// f07f80
}

func ExampleReadUint16BE() {
	b, _ := hex.DecodeString("000180000200")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint16BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 1
	// 32768
	// 512
}

func ExampleWriteUint16BE() {
	w := new(bytes.Buffer)

	data := []uint16{1, 32768, 512}
	for _, v := range data {
		if err := typeio.WriteUint16BE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 000180000200
}

func ExampleReadUint16LE() {
	b, _ := hex.DecodeString("000180000200")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint16LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 256
	// 128
	// 2
}

func ExampleWriteUint16LE() {
	w := new(bytes.Buffer)

	data := []uint16{256, 128, 2}
	for _, v := range data {
		if err := typeio.WriteUint16LE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 000180000200
}

func ExampleReadInt16BE() {
	b, _ := hex.DecodeString("00018000fffe")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt16BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 1
	// -32768
	// -2
}

func ExampleWriteInt16BE() {
	w := new(bytes.Buffer)

	data := []int16{1, -32768, -2}
	for _, v := range data {
		if err := typeio.WriteInt16BE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 00018000fffe
}

func ExampleReadInt16LE() {
	b, _ := hex.DecodeString("00010080feff")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt16LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Println(v)
	}

	// Output:
	// 256
	// -32768
	// -2
}

func ExampleWriteInt16LE() {
	w := new(bytes.Buffer)

	data := []int16{256, -32768, -2}
	for _, v := range data {
		if err := typeio.WriteInt16LE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 00010080feff
}

func ExampleReadUint32BE() {
	b, _ := hex.DecodeString("01020304fffffffe")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint32BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 1020304
	// fffffffe
}

func ExampleWriteUint32BE() {
	w := new(bytes.Buffer)

	data := []uint32{0x01020304, 0xfffffffe}
	for _, v := range data {
		if err := typeio.WriteUint32BE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 01020304fffffffe
}

func ExampleReadUint32LE() {
	b, _ := hex.DecodeString("01020304feffffff")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint32LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 4030201
	// fffffffe
}

func ExampleWriteUint32LE() {
	w := new(bytes.Buffer)

	data := []uint32{0x01020304, 0xfffffffe}
	for _, v := range data {
		if err := typeio.WriteUint32LE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 04030201feffffff
}

func ExampleReadInt32BE() {
	b, _ := hex.DecodeString("01020304fffffffe")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt32BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 1020304
	// -2
}

func ExampleWriteInt32BE() {
	w := new(bytes.Buffer)

	data := []int32{0x01020304, -2}
	for _, v := range data {
		if err := typeio.WriteInt32BE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 01020304fffffffe
}

func ExampleReadInt32LE() {
	b, _ := hex.DecodeString("01020304feffffff")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt32LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 4030201
	// -2
}

func ExampleWriteInt32LE() {
	w := new(bytes.Buffer)

	data := []int32{0x01020304, -2}
	for _, v := range data {
		if err := typeio.WriteInt32LE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 04030201feffffff
}

func ExampleReadUint64BE() {
	b, _ := hex.DecodeString("0102030405060708fffffffffffffffe")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint64BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 102030405060708
	// fffffffffffffffe
}

func ExampleWriteUint64BE() {
	w := new(bytes.Buffer)

	data := []uint64{0x0102030405060708, 0xfffffffffffffffe}
	for _, v := range data {
		if err := typeio.WriteUint64BE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 0102030405060708fffffffffffffffe
}

func ExampleReadUint64LE() {
	b, _ := hex.DecodeString("0102030405060708feffffffffffffff")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadUint64LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 807060504030201
	// fffffffffffffffe
}

func ExampleWriteUint64LE() {
	w := new(bytes.Buffer)

	data := []uint64{0x0102030405060708, 0xfffffffffffffffe}
	for _, v := range data {
		if err := typeio.WriteUint64LE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 0807060504030201feffffffffffffff
}

func ExampleReadInt64BE() {
	b, _ := hex.DecodeString("0102030405060708fffffffffffffffe")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt64BE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 102030405060708
	// -2
}

func ExampleWriteInt64BE() {
	w := new(bytes.Buffer)

	data := []int64{0x0102030405060708, -2}
	for _, v := range data {
		if err := typeio.WriteInt64BE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 0102030405060708fffffffffffffffe
}

func ExampleReadInt64LE() {
	b, _ := hex.DecodeString("0102030405060708feffffffffffffff")
	r := bytes.NewReader(b)

	for {
		v, err := typeio.ReadInt64LE(r)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		fmt.Printf("%x\n", v)
	}

	// Output:
	// 807060504030201
	// -2
}

func ExampleWriteInt64LE() {
	w := new(bytes.Buffer)

	data := []int64{0x0102030405060708, -2}
	for _, v := range data {
		if err := typeio.WriteInt64LE(w, v); err != nil {
			panic(err)
		}
	}
	fmt.Println(hex.EncodeToString(w.Bytes()))

	// Output:
	// 0807060504030201feffffffffffffff
}
