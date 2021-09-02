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
	"time"

	"github.com/tunabay/go-typeio"
)

// loceq compares two time.Location x and y.
func loceq(x, y *time.Location) bool {
	xt := time.Date(2006, 1, 2, 15, 4, 5, 6, x)
	yt := time.Date(2006, 1, 2, 15, 4, 5, 6, y)
	return xt.Equal(yt)
}

func TestReadUnixTimeUTC32BE(t *testing.T) {
	locJST, _ := time.LoadLocation("Asia/Tokyo")
	tcs := []struct {
		b string
		t time.Time
		e error
	}{
		{"00000000", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), nil},
		{"4fb98412", time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC), nil},
		{"4fb98412", time.Date(2012, 5, 21, 8, 53, 54, 0, locJST), nil},
		{"c0", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8ff", time.Time{}, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUnixTimeUTC32BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %v", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !got.Equal(tc.t):
			t.Errorf("%q: unexpected read: got %v, want %v", tc.b, got, tc.t)
		case tc.e == nil && !loceq(got.Location(), time.UTC):
			t.Errorf("%q: unexpected loc: got %s, want %s", tc.b, got.Location(), time.UTC)
		}
	}
}

func TestReadUnixTime32BE(t *testing.T) {
	locJST, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = locJST
	tcs := []struct {
		b string
		t time.Time
		e error
	}{
		{"00000000", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), nil},
		{"4fb98412", time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC), nil},
		{"4fb98412", time.Date(2012, 5, 21, 8, 53, 54, 0, locJST), nil},
		{"c0", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8ff", time.Time{}, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUnixTime32BE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %v", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !got.Equal(tc.t):
			t.Errorf("%q: unexpected read: got %v, want %v", tc.b, got, tc.t)
		case tc.e == nil && !loceq(got.Location(), locJST):
			t.Errorf("%q: unexpected loc: got %s, want %s", tc.b, got.Location(), locJST)
		}
	}
}

func TestWriteUnixTime32BE(t *testing.T) {
	locJST, _ := time.LoadLocation("Asia/Tokyo")
	tcs := []struct {
		t time.Time
		b string
		e error
	}{
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), "00000000", nil},
		{time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC), "4fb98412", nil},
		{time.Date(2012, 5, 21, 8, 53, 54, 0, locJST), "4fb98412", nil},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		err := typeio.WriteUnixTime32BE(w, tc.t)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%v: error expected.", tc.t)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%v: unexpected type of error: got %q, want %q", tc.t, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%v: unexpected error: %s", tc.t, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%v: unexpected write: got %s, want %s", tc.t, got, tc.b)
		}
	}
}

func TestReadUnixTimeUTC32LE(t *testing.T) {
	locJST, _ := time.LoadLocation("Asia/Tokyo")
	tcs := []struct {
		b string
		t time.Time
		e error
	}{
		{"00000000", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), nil},
		{"1284b94f", time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC), nil},
		{"1284b94f", time.Date(2012, 5, 21, 8, 53, 54, 0, locJST), nil},
		{"c0", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8ff", time.Time{}, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUnixTimeUTC32LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %v", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !got.Equal(tc.t):
			t.Errorf("%q: unexpected read: got %v, want %v", tc.b, got, tc.t)
		case tc.e == nil && !loceq(got.Location(), time.UTC):
			t.Errorf("%q: unexpected loc: got %s, want %s", tc.b, got.Location(), time.UTC)
		}
	}
}

func TestReadUnixTime32LE(t *testing.T) {
	locJST, _ := time.LoadLocation("Asia/Tokyo")
	time.Local = locJST
	tcs := []struct {
		b string
		t time.Time
		e error
	}{
		{"00000000", time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), nil},
		{"1284b94f", time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC), nil},
		{"1284b94f", time.Date(2012, 5, 21, 8, 53, 54, 0, locJST), nil},
		{"c0", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8", time.Time{}, io.ErrUnexpectedEOF},
		{"c0a8ff", time.Time{}, io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadUnixTime32LE(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %v", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !got.Equal(tc.t):
			t.Errorf("%q: unexpected read: got %v, want %v", tc.b, got, tc.t)
		case tc.e == nil && !loceq(got.Location(), locJST):
			t.Errorf("%q: unexpected loc: got %s, want %s", tc.b, got.Location(), locJST)
		}
	}
}

func TestWriteUnixTime32LE(t *testing.T) {
	locJST, _ := time.LoadLocation("Asia/Tokyo")
	tcs := []struct {
		t time.Time
		b string
		e error
	}{
		{time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC), "00000000", nil},
		{time.Date(2012, 5, 20, 23, 53, 54, 0, time.UTC), "1284b94f", nil},
		{time.Date(2012, 5, 21, 8, 53, 54, 0, locJST), "1284b94f", nil},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		err := typeio.WriteUnixTime32LE(w, tc.t)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%v: error expected.", tc.t)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%v: unexpected type of error: got %q, want %q", tc.t, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%v: unexpected error: %s", tc.t, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%v: unexpected write: got %s, want %s", tc.t, got, tc.b)
		}
	}
}
