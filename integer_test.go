// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"testing"

	"github.com/tunabay/go-typeio"
)

func TestReadUint8(t *testing.T) {
	tcs := []struct {
		b string
		v uint8
		e error
	}{
		{"00", 0, nil},
		{"0102", 1, nil},
		{"7f", 127, nil},
		{"80", 128, nil},
		{"fe", 254, nil},
		{"ff", 255, nil},
		{"aaff", 0xAA, nil},
		{"", 0, io.EOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint8(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %d", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %d, want %d", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint8(t *testing.T) {
	tcs := []struct {
		v uint8
		b string
	}{
		{0, "00"},
		{1, "01"},
		{127, "7f"},
		{128, "80"},
		{254, "fe"},
		{255, "ff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint8(w, tc.v); err != nil {
			t.Errorf("%d: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%d: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt8(t *testing.T) {
	tcs := []struct {
		b string
		v int8
		e error
	}{
		{"00", 0, nil},
		{"0102", 1, nil},
		{"7f", 127, nil},
		{"80", -128, nil},
		{"fe", -2, nil},
		{"ff", -1, nil},
		{"aaff", -86, nil},
		{"", 0, io.EOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt8(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %d", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %d, want %d", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt8(t *testing.T) {
	tcs := []struct {
		v int8
		b string
	}{
		{0, "00"},
		{1, "01"},
		{127, "7f"},
		{-128, "80"},
		{-2, "fe"},
		{-1, "ff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt8(w, tc.v); err != nil {
			t.Errorf("%d: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%d: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadUint16BE(t *testing.T) {
	tcs := []struct {
		b string
		v uint16
		e error
	}{
		{"0000", 0, nil},
		{"0001", 1, nil},
		{"1234", 0x1234, nil},
		{"ff00", 0xff00, nil},
		{"7fff", 0x7fff, nil},
		{"8000", 0x8000, nil},
		{"fffe", 0xfffe, nil},
		{"ffff", 0xffff, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint16BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint16BE(t *testing.T) {
	tcs := []struct {
		v uint16
		b string
	}{
		{0, "0000"},
		{1, "0001"},
		{0x1234, "1234"},
		{0xff00, "ff00"},
		{0x7fff, "7fff"},
		{0x8000, "8000"},
		{0xfffe, "fffe"},
		{0xffff, "ffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint16BE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadUint16LE(t *testing.T) {
	tcs := []struct {
		b string
		v uint16
		e error
	}{
		{"0000", 0, nil},
		{"0100", 1, nil},
		{"1234", 0x3412, nil},
		{"ff00", 0x00ff, nil},
		{"ff7f", 0x7fff, nil},
		{"0080", 0x8000, nil},
		{"feff", 0xfffe, nil},
		{"ffff", 0xffff, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint16LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint16LE(t *testing.T) {
	tcs := []struct {
		v uint16
		b string
	}{
		{0, "0000"},
		{1, "0100"},
		{0x1234, "3412"},
		{0xff00, "00ff"},
		{0x7fff, "ff7f"},
		{0x8000, "0080"},
		{0xfffe, "feff"},
		{0xffff, "ffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint16LE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt16BE(t *testing.T) {
	tcs := []struct {
		b string
		v int16
		e error
	}{
		{"0000", 0, nil},
		{"0001", 1, nil},
		{"1234", 0x1234, nil},
		{"ff00", -0x0100, nil},
		{"7fff", 0x7fff, nil},
		{"8000", -0x8000, nil},
		{"fffe", -2, nil},
		{"ffff", -1, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt16BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt16BE(t *testing.T) {
	tcs := []struct {
		v int16
		b string
	}{
		{0, "0000"},
		{1, "0001"},
		{0x1234, "1234"},
		{-0x0100, "ff00"},
		{0x7fff, "7fff"},
		{-0x8000, "8000"},
		{-2, "fffe"},
		{-1, "ffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt16BE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt16LE(t *testing.T) {
	tcs := []struct {
		b string
		v int16
		e error
	}{
		{"0000", 0, nil},
		{"0100", 1, nil},
		{"1234", 0x3412, nil},
		{"ff00", 0x00ff, nil},
		{"ff7f", 0x7fff, nil},
		{"0080", -0x8000, nil},
		{"feff", -2, nil},
		{"ffff", -1, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt16LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt16LE(t *testing.T) {
	tcs := []struct {
		v int16
		b string
	}{
		{0, "0000"},
		{1, "0100"},
		{0x1234, "3412"},
		{-0x0100, "00ff"},
		{0x7fff, "ff7f"},
		{-0x8000, "0080"},
		{-2, "feff"},
		{-1, "ffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt16LE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadUint32BE(t *testing.T) {
	tcs := []struct {
		b string
		v uint32
		e error
	}{
		{"00000000", 0, nil},
		{"00000001", 1, nil},
		{"12345678", 0x12345678, nil},
		{"ffcc3300", 0xffcc3300, nil},
		{"7fffffff", 0x7fffffff, nil},
		{"80000000", 0x80000000, nil},
		{"fffffffe", 0xfffffffe, nil},
		{"ffffffff", 0xffffffff, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint32BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint32BE(t *testing.T) {
	tcs := []struct {
		v uint32
		b string
	}{
		{0, "00000000"},
		{1, "00000001"},
		{0x12345678, "12345678"},
		{0xffcc3300, "ffcc3300"},
		{0x7fffffff, "7fffffff"},
		{0x80000000, "80000000"},
		{0xfffffffe, "fffffffe"},
		{0xffffffff, "ffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint32BE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadUint32LE(t *testing.T) {
	tcs := []struct {
		b string
		v uint32
		e error
	}{
		{"00000000", 0, nil},
		{"01000000", 1, nil},
		{"12345678", 0x78563412, nil},
		{"ffcc3300", 0x0033ccff, nil},
		{"ffffff7f", 0x7fffffff, nil},
		{"00000080", 0x80000000, nil},
		{"feffffff", 0xfffffffe, nil},
		{"ffffffff", 0xffffffff, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint32LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint32LE(t *testing.T) {
	tcs := []struct {
		v uint32
		b string
	}{
		{0, "00000000"},
		{1, "01000000"},
		{0x12345678, "78563412"},
		{0xffcc3300, "0033ccff"},
		{0x7fffffff, "ffffff7f"},
		{0x80000000, "00000080"},
		{0xfffffffe, "feffffff"},
		{0xffffffff, "ffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint32LE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt32BE(t *testing.T) {
	tcs := []struct {
		b string
		v int32
		e error
	}{
		{"00000000", 0, nil},
		{"00000001", 1, nil},
		{"12345678", 0x12345678, nil},
		{"ffcc3300", -0x0033cd00, nil},
		{"7fffffff", 0x7fffffff, nil},
		{"80000000", -0x80000000, nil},
		{"fffffffe", -2, nil},
		{"ffffffff", -1, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt32BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt32BE(t *testing.T) {
	tcs := []struct {
		v int32
		b string
	}{
		{0, "00000000"},
		{1, "00000001"},
		{0x12345678, "12345678"},
		{-0x0033cd00, "ffcc3300"},
		{0x7fffffff, "7fffffff"},
		{-0x80000000, "80000000"},
		{-2, "fffffffe"},
		{-1, "ffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt32BE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt32LE(t *testing.T) {
	tcs := []struct {
		b string
		v int32
		e error
	}{
		{"00000000", 0, nil},
		{"01000000", 1, nil},
		{"12345678", 0x78563412, nil},
		{"ffcc3300", 0x0033ccff, nil},
		{"ffffff7f", 0x7fffffff, nil},
		{"00000080", -0x80000000, nil},
		{"feffffff", -2, nil},
		{"ffffffff", -1, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt32LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt32LE(t *testing.T) {
	tcs := []struct {
		v int32
		b string
	}{
		{0, "00000000"},
		{1, "01000000"},
		{0x12345678, "78563412"},
		{-0x0033cd00, "0033ccff"},
		{0x7fffffff, "ffffff7f"},
		{-0x80000000, "00000080"},
		{-2, "feffffff"},
		{-1, "ffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt32LE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadUint64BE(t *testing.T) {
	tcs := []struct {
		b string
		v uint64
		e error
	}{
		{"0000000000000000", 0, nil},
		{"0000000000000001", 1, nil},
		{"1234567887654321", 0x1234567887654321, nil},
		{"ffeeddccbbaa9988", 0xffeeddccbbaa9988, nil},
		{"7fffffffffffffff", 0x7fffffffffffffff, nil},
		{"8000000000000000", 0x8000000000000000, nil},
		{"fffffffffffffffe", 0xfffffffffffffffe, nil},
		{"ffffffffffffffff", 0xffffffffffffffff, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint64BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint64BE(t *testing.T) {
	tcs := []struct {
		v uint64
		b string
	}{
		{0, "0000000000000000"},
		{1, "0000000000000001"},
		{0x1234567887654321, "1234567887654321"},
		{0xffeeddccbbaa9988, "ffeeddccbbaa9988"},
		{0x7fffffffffffffff, "7fffffffffffffff"},
		{0x8000000000000000, "8000000000000000"},
		{0xfffffffffffffffe, "fffffffffffffffe"},
		{0xffffffffffffffff, "ffffffffffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint64BE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadUint64LE(t *testing.T) {
	tcs := []struct {
		b string
		v uint64
		e error
	}{
		{"0000000000000000", 0, nil},
		{"0100000000000000", 1, nil},
		{"1234567887654321", 0x2143658778563412, nil},
		{"ffeeddccbbaa9988", 0x8899aabbccddeeff, nil},
		{"ffffffffffffff7f", 0x7fffffffffffffff, nil},
		{"0000000000000080", 0x8000000000000000, nil},
		{"feffffffffffffff", 0xfffffffffffffffe, nil},
		{"ffffffffffffffff", 0xffffffffffffffff, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUint64LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteUint64LE(t *testing.T) {
	tcs := []struct {
		v uint64
		b string
	}{
		{0, "0000000000000000"},
		{1, "0100000000000000"},
		{0x1234567887654321, "2143658778563412"},
		{0xffeeddccbbaa9988, "8899aabbccddeeff"},
		{0x7fffffffffffffff, "ffffffffffffff7f"},
		{0x8000000000000000, "0000000000000080"},
		{0xfffffffffffffffe, "feffffffffffffff"},
		{0xffffffffffffffff, "ffffffffffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteUint64LE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt64BE(t *testing.T) {
	tcs := []struct {
		b string
		v int64
		e error
	}{
		{"0000000000000000", 0, nil},
		{"0000000000000001", 1, nil},
		{"1234567887654321", 0x1234567887654321, nil},
		{"ffeeddccbbaa9988", -0x0011223344556678, nil},
		{"7fffffffffffffff", 0x7fffffffffffffff, nil},
		{"8000000000000000", -0x8000000000000000, nil},
		{"fffffffffffffffe", -2, nil},
		{"ffffffffffffffff", -1, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt64BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt64BE(t *testing.T) {
	tcs := []struct {
		v int64
		b string
	}{
		{0, "0000000000000000"},
		{1, "0000000000000001"},
		{0x1234567887654321, "1234567887654321"},
		{-0x0011223344556678, "ffeeddccbbaa9988"},
		{0x7fffffffffffffff, "7fffffffffffffff"},
		{-0x8000000000000000, "8000000000000000"},
		{-2, "fffffffffffffffe"},
		{-1, "ffffffffffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt64BE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}

func TestReadInt64LE(t *testing.T) {
	tcs := []struct {
		b string
		v int64
		e error
	}{
		{"0000000000000000", 0, nil},
		{"0100000000000000", 1, nil},
		{"1234567887654321", 0x2143658778563412, nil},
		{"ffeeddccbbaa9988", -0x7766554433221101, nil},
		{"ffffffffffffff7f", 0x7fffffffffffffff, nil},
		{"0000000000000080", -0x8000000000000000, nil},
		{"feffffffffffffff", -2, nil},
		{"ffffffffffffffff", -1, nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadInt64LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %x", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.v:
			t.Errorf("%q: unexpected read: got %x, want %x", tc.b, got, tc.v)
		}
	}
}

func TestWriteInt64LE(t *testing.T) {
	tcs := []struct {
		v int64
		b string
	}{
		{0, "0000000000000000"},
		{1, "0100000000000000"},
		{0x1234567887654321, "2143658778563412"},
		{-0x0011223344556678, "8899aabbccddeeff"},
		{0x7fffffffffffffff, "ffffffffffffff7f"},
		{-0x8000000000000000, "0000000000000080"},
		{-2, "feffffffffffffff"},
		{-1, "ffffffffffffffff"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteInt64LE(w, tc.v); err != nil {
			t.Errorf("%x: unexpected error: %s", tc.v, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%x: unexpected write: got %s, want %s", tc.v, got, tc.b)
		}
	}
}
