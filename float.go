// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio

import (
	"encoding/binary"
	"io"
	"math"
)

// ReadFloat32BE reads 4 bytes in big-endian byte order from r and returns them
// as an float32 as defined in IEEE 754.
func ReadFloat32BE(r io.Reader) (float32, error) {
	b, err := readN(r, 4)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.BigEndian.Uint32(b)), nil
}

// WriteFloat32BE writes 4 bytes to w that represent the IEEE 754 float32 value
// v in big-endian byte order.
func WriteFloat32BE(w io.Writer, v float32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, math.Float32bits(v))
	return write(w, b)
}

// ReadFloat32LE reads 4 bytes in big-endian byte order from r and returns them
// as an float32 as defined in IEEE 754.
func ReadFloat32LE(r io.Reader) (float32, error) {
	b, err := readN(r, 4)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(binary.LittleEndian.Uint32(b)), nil
}

// WriteFloat32LE writes 4 bytes to w that represent the IEEE 754 float32 value
// v in big-endian byte order.
func WriteFloat32LE(w io.Writer, v float32) error {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, math.Float32bits(v))
	return write(w, b)
}

// ReadFloat64BE reads 8 bytes in big-endian byte order from r and returns them
// as an float64 as defined in IEEE 754.
func ReadFloat64BE(r io.Reader) (float64, error) {
	b, err := readN(r, 8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.BigEndian.Uint64(b)), nil
}

// WriteFloat64BE writes 8 bytes to w that represent the IEEE 754 float64 value
// v in big-endian byte order.
func WriteFloat64BE(w io.Writer, v float64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, math.Float64bits(v))
	return write(w, b)
}

// ReadFloat64LE reads 8 bytes in big-endian byte order from r and returns them
// as an float64 as defined in IEEE 754.
func ReadFloat64LE(r io.Reader) (float64, error) {
	b, err := readN(r, 8)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(binary.LittleEndian.Uint64(b)), nil
}

// WriteFloat64LE writes 8 bytes to w that represent the IEEE 754 float64 value
// v in big-endian byte order.
func WriteFloat64LE(w io.Writer, v float64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, math.Float64bits(v))
	return write(w, b)
}
