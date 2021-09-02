// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio

import (
	"encoding/binary"
	"io"
)

// ReadUint8 reads 1 byte from r and returns it as a uint8 value.
func ReadUint8(r io.Reader) (uint8, error) {
	b, err := readN(r, 1)
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

// WriteUint8 writes 1 byte to w that represents the value v of uint8.
func WriteUint8(w io.Writer, v uint8) error {
	return write(w, []byte{v})
}

// ReadInt8 reads 1 byte from r and returns it as an int8 value.
func ReadInt8(r io.Reader) (int8, error) {
	b, err := readN(r, 1)
	if err != nil {
		return 0, err
	}
	return int8(b[0]), nil
}

// WriteInt8 writes 1 byte to w that represents the value v of int8.
func WriteInt8(w io.Writer, v int8) error {
	return write(w, []byte{uint8(v)})
}

// ReadUint16BE reads 2 bytes in big-endian byte order from r and returns them
// as a uint16 value.
func ReadUint16BE(r io.Reader) (uint16, error) {
	b, err := readN(r, 2)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint16(b), nil
}

// WriteUint16BE writes 2 bytes to w that represent the value v of uint16 in
// big-endian byte order.
func WriteUint16BE(w io.Writer, v uint16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, v)
	return write(w, b)
}

// ReadUint16LE reads 2 bytes in little-endian byte order from r and returns
// them as a uint16 value.
func ReadUint16LE(r io.Reader) (uint16, error) {
	b, err := readN(r, 2)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint16(b), nil
}

// WriteUint16LE writes 2 bytes to w that represent the value v of uint16 in
// little-endian byte order.
func WriteUint16LE(w io.Writer, v uint16) error {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return write(w, b)
}

// ReadInt16BE reads 2 bytes in big-endian byte order from r and returns them as
// an int16 value.
func ReadInt16BE(r io.Reader) (int16, error) {
	b, err := readN(r, 2)
	if err != nil {
		return 0, err
	}
	return int16(binary.BigEndian.Uint16(b)), nil
}

// WriteInt16BE writes 2 bytes to w that represent the value v of int16 in
// big-endian byte order.
func WriteInt16BE(w io.Writer, v int16) error {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(v))
	return write(w, b)
}

// ReadInt16LE reads 2 bytes in little-endian byte order from r and returns them
// as an int16 value.
func ReadInt16LE(r io.Reader) (int16, error) {
	b, err := readN(r, 2)
	if err != nil {
		return 0, err
	}
	return int16(binary.LittleEndian.Uint16(b)), nil
}

// WriteInt16LE writes 2 bytes to w that represent the value v of int16 in
// little-endian byte order.
func WriteInt16LE(w io.Writer, v int16) error {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, uint16(v))
	return write(w, b)
}

// ReadUint32BE reads 4 bytes in big-endian byte order from r and returns them
// as a uint32 value.
func ReadUint32BE(r io.Reader) (uint32, error) {
	b, err := readN(r, 4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(b), nil
}

// WriteUint32BE writes 4 bytes to w that represent the value v of uint32 in
// big-endian byte order.
func WriteUint32BE(w io.Writer, v uint32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return write(w, b)
}

// ReadUint32LE reads 4 bytes in little-endian byte order from r and returns
// them as a uint32 value.
func ReadUint32LE(r io.Reader) (uint32, error) {
	b, err := readN(r, 4)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint32(b), nil
}

// WriteUint32LE writes 4 bytes to w that represent the value v of uint32 in
// little-endian byte order.
func WriteUint32LE(w io.Writer, v uint32) error {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return write(w, b)
}

// ReadInt32BE reads 4 bytes in big-endian byte order from r and returns them as
// an int32 value.
func ReadInt32BE(r io.Reader) (int32, error) {
	b, err := readN(r, 4)
	if err != nil {
		return 0, err
	}
	return int32(binary.BigEndian.Uint32(b)), nil
}

// WriteInt32BE writes 4 bytes to w that represent the value v of int32 in
// big-endian byte order.
func WriteInt32BE(w io.Writer, v int32) error {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(v))
	return write(w, b)
}

// ReadInt32LE reads 4 bytes in little-endian byte order from r and returns them
// as an int32 value.
func ReadInt32LE(r io.Reader) (int32, error) {
	b, err := readN(r, 4)
	if err != nil {
		return 0, err
	}
	return int32(binary.LittleEndian.Uint32(b)), nil
}

// WriteInt32LE writes 4 bytes to w that represent the value v of int32 in
// little-endian byte order.
func WriteInt32LE(w io.Writer, v int32) error {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, uint32(v))
	return write(w, b)
}

// ReadUint64BE reads 8 bytes in big-endian byte order from r and returns them
// as a uint64 value.
func ReadUint64BE(r io.Reader) (uint64, error) {
	b, err := readN(r, 8)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint64(b), nil
}

// WriteUint64BE writes 8 bytes to w that represent the value v of uint64 in
// big-endian byte order.
func WriteUint64BE(w io.Writer, v uint64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return write(w, b)
}

// ReadUint64LE reads 8 bytes in little-endian byte order from r and returns
// them as a uint64 value.
func ReadUint64LE(r io.Reader) (uint64, error) {
	b, err := readN(r, 8)
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b), nil
}

// WriteUint64LE writes 8 bytes to w that represent the value v of uint64 in
// little-endian byte order.
func WriteUint64LE(w io.Writer, v uint64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, v)
	return write(w, b)
}

// ReadInt64BE reads 8 bytes in big-endian byte order from r and returns them as
// an int64 value.
func ReadInt64BE(r io.Reader) (int64, error) {
	b, err := readN(r, 8)
	if err != nil {
		return 0, err
	}
	return int64(binary.BigEndian.Uint64(b)), nil
}

// WriteInt64BE writes 8 bytes to w that represent the value v of int64 in
// big-endian byte order.
func WriteInt64BE(w io.Writer, v int64) error {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return write(w, b)
}

// ReadInt64LE reads 8 bytes in little-endian byte order from r and returns them
// as an int64 value.
func ReadInt64LE(r io.Reader) (int64, error) {
	b, err := readN(r, 8)
	if err != nil {
		return 0, err
	}
	return int64(binary.LittleEndian.Uint64(b)), nil
}

// WriteInt64LE writes 8 bytes to w that represent the value v of int64 in
// little-endian byte order.
func WriteInt64LE(w io.Writer, v int64) error {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(v))
	return write(w, b)
}
