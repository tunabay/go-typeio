// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"math"
	"strconv"
	"testing"

	"github.com/tunabay/go-typeio"
)

func f32eq(x, y float32) bool { return x == y || (math.IsNaN(float64(x)) && math.IsNaN(float64(y))) }
func f64eq(x, y float64) bool { return x == y || (math.IsNaN(x) && math.IsNaN(y)) }
func f32s(v float32) string   { return strconv.FormatFloat(float64(v), 'g', -1, 32) }
func f64s(v float64) string   { return strconv.FormatFloat(v, 'g', -1, 64) }

func TestReadFloat32BE(t *testing.T) {
	tcs := []struct {
		b string
		v float32
		e error
	}{
		{"00000000", 0, nil},
		{"40490fdb", math.Pi, nil},
		{"c0490fdb", -math.Pi, nil},
		{"50ca871c", math.E * 10000000000, nil},
		{"d0ca871c", -math.E * 10000000000, nil},
		{"7f7fffff", math.MaxFloat32, nil},
		{"ff7fffff", -math.MaxFloat32, nil},
		{"00000001", math.SmallestNonzeroFloat32, nil},
		{"80000001", -math.SmallestNonzeroFloat32, nil},
		{"7f800000", float32(math.Inf(+1)), nil},
		{"ff800000", float32(math.Inf(-1)), nil},
		{"7fc00000", float32(math.NaN()), nil},
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
		got, err := typeio.ReadFloat32BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %s", tc.b, f32s(got))
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !f32eq(got, tc.v):
			t.Errorf("%q: unexpected read: got %s, want %s", tc.b, f32s(got), f32s(tc.v))
		}
	}
}

func TestWriteFloat32BE(t *testing.T) {
	tcs := []struct {
		v float32
		b string
	}{
		{0, "00000000"},
		{math.Pi, "40490fdb"},
		{-math.Pi, "c0490fdb"},
		{math.E * 10000000000, "50ca871c"},
		{-math.E * 10000000000, "d0ca871c"},
		{math.MaxFloat32, "7f7fffff"},
		{-math.MaxFloat32, "ff7fffff"},
		{math.SmallestNonzeroFloat32, "00000001"},
		{-math.SmallestNonzeroFloat32, "80000001"},
		{float32(math.Inf(+1)), "7f800000"},
		{float32(math.Inf(-1)), "ff800000"},
		{float32(math.NaN()), "7fc00000"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteFloat32BE(w, tc.v); err != nil {
			t.Errorf("%s: unexpected error: %s", f32s(tc.v), err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%s: unexpected write: got %s, want %s", f32s(tc.v), got, tc.b)
		}
	}
}

func TestReadFloat32LE(t *testing.T) {
	tcs := []struct {
		b string
		v float32
		e error
	}{
		{"00000000", 0, nil},
		{"db0f4940", math.Pi, nil},
		{"db0f49c0", -math.Pi, nil},
		{"1c87ca50", math.E * 10000000000, nil},
		{"1c87cad0", -math.E * 10000000000, nil},
		{"ffff7f7f", math.MaxFloat32, nil},
		{"ffff7fff", -math.MaxFloat32, nil},
		{"01000000", math.SmallestNonzeroFloat32, nil},
		{"01000080", -math.SmallestNonzeroFloat32, nil},
		{"0000807f", float32(math.Inf(+1)), nil},
		{"000080ff", float32(math.Inf(-1)), nil},
		{"0000c07f", float32(math.NaN()), nil},
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
		got, err := typeio.ReadFloat32LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %s", tc.b, f32s(got))
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !f32eq(got, tc.v):
			t.Errorf("%q: unexpected read: got %s, want %s", tc.b, f32s(got), f32s(tc.v))
		}
	}
}

func TestWriteFloat32LE(t *testing.T) {
	tcs := []struct {
		v float32
		b string
	}{
		{0, "00000000"},
		{math.Pi, "db0f4940"},
		{-math.Pi, "db0f49c0"},
		{math.E * 10000000000, "1c87ca50"},
		{-math.E * 10000000000, "1c87cad0"},
		{math.MaxFloat32, "ffff7f7f"},
		{-math.MaxFloat32, "ffff7fff"},
		{math.SmallestNonzeroFloat32, "01000000"},
		{-math.SmallestNonzeroFloat32, "01000080"},
		{float32(math.Inf(+1)), "0000807f"},
		{float32(math.Inf(-1)), "000080ff"},
		{float32(math.NaN()), "0000c07f"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteFloat32LE(w, tc.v); err != nil {
			t.Errorf("%s: unexpected error: %s", f32s(tc.v), err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%s: unexpected write: got %s, want %s", f32s(tc.v), got, tc.b)
		}
	}
}

func TestReadFloat64BE(t *testing.T) {
	tcs := []struct {
		b string
		v float64
		e error
	}{
		{"0000000000000000", 0, nil},
		{"400921fb54442d18", math.Pi, nil},
		{"c00921fb54442d18", -math.Pi, nil},
		{"421950e38fb25ca0", math.E * 10000000000, nil},
		{"c21950e38fb25ca0", -math.E * 10000000000, nil},
		{"7fefffffffffffff", math.MaxFloat64, nil},
		{"ffefffffffffffff", -math.MaxFloat64, nil},
		{"0000000000000001", math.SmallestNonzeroFloat64, nil},
		{"8000000000000001", -math.SmallestNonzeroFloat64, nil},
		{"7ff0000000000000", math.Inf(+1), nil},
		{"fff0000000000000", math.Inf(-1), nil},
		{"7ff8000000000001", math.NaN(), nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadFloat64BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %s", tc.b, f64s(got))
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !f64eq(got, tc.v):
			t.Errorf("%q: unexpected read: got %s, want %s", tc.b, f64s(got), f64s(tc.v))
		}
	}
}

func TestWriteFloat64BE(t *testing.T) {
	tcs := []struct {
		v float64
		b string
	}{
		{0, "0000000000000000"},
		{math.Pi, "400921fb54442d18"},
		{-math.Pi, "c00921fb54442d18"},
		{math.E * 10000000000, "421950e38fb25ca0"},
		{-math.E * 10000000000, "c21950e38fb25ca0"},
		{math.MaxFloat64, "7fefffffffffffff"},
		{-math.MaxFloat64, "ffefffffffffffff"},
		{math.SmallestNonzeroFloat64, "0000000000000001"},
		{-math.SmallestNonzeroFloat64, "8000000000000001"},
		{math.Inf(+1), "7ff0000000000000"},
		{math.Inf(-1), "fff0000000000000"},
		{math.NaN(), "7ff8000000000001"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteFloat64BE(w, tc.v); err != nil {
			t.Errorf("%s: unexpected error: %s", f64s(tc.v), err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%s: unexpected write: got %s, want %s", f64s(tc.v), got, tc.b)
		}
	}
}

func TestReadFloat64LE(t *testing.T) {
	tcs := []struct {
		b string
		v float64
		e error
	}{
		{"0000000000000000", 0, nil},
		{"182d4454fb210940", math.Pi, nil},
		{"182d4454fb2109c0", -math.Pi, nil},
		{"a05cb28fe3501942", math.E * 10000000000, nil},
		{"a05cb28fe35019c2", -math.E * 10000000000, nil},
		{"ffffffffffffef7f", math.MaxFloat64, nil},
		{"ffffffffffffefff", -math.MaxFloat64, nil},
		{"0100000000000000", math.SmallestNonzeroFloat64, nil},
		{"0100000000000080", -math.SmallestNonzeroFloat64, nil},
		{"000000000000f07f", math.Inf(+1), nil},
		{"000000000000f0ff", math.Inf(-1), nil},
		{"010000000000f87f", math.NaN(), nil},
		{"", 0, io.EOF},
		{"ff", 0, io.ErrUnexpectedEOF},
		{"ffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffff", 0, io.ErrUnexpectedEOF},
		{"ffffffffffffff", 0, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadFloat64LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %s", tc.b, f64s(got))
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !f64eq(got, tc.v):
			t.Errorf("%q: unexpected read: got %s, want %s", tc.b, f64s(got), f64s(tc.v))
		}
	}
}

func TestWriteFloat64LE(t *testing.T) {
	tcs := []struct {
		v float64
		b string
	}{
		{0, "0000000000000000"},
		{math.Pi, "182d4454fb210940"},
		{-math.Pi, "182d4454fb2109c0"},
		{math.E * 10000000000, "a05cb28fe3501942"},
		{-math.E * 10000000000, "a05cb28fe35019c2"},
		{math.MaxFloat64, "ffffffffffffef7f"},
		{-math.MaxFloat64, "ffffffffffffefff"},
		{math.SmallestNonzeroFloat64, "0100000000000000"},
		{-math.SmallestNonzeroFloat64, "0100000000000080"},
		{math.Inf(+1), "000000000000f07f"},
		{math.Inf(-1), "000000000000f0ff"},
		{math.NaN(), "010000000000f87f"},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		if err := typeio.WriteFloat64LE(w, tc.v); err != nil {
			t.Errorf("%s: unexpected error: %s", f64s(tc.v), err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%s: unexpected write: got %s, want %s", f64s(tc.v), got, tc.b)
		}
	}
}
