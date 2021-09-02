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
	"testing"

	"github.com/tunabay/go-typeio"
)

func TestReadStringN(t *testing.T) {
	tcs := []struct {
		b string
		p byte
		n int
		s string
		e error
	}{
		{"", 0, 0, "", nil},
		{"", 0, 1, "", io.EOF},
		{"00000000", 0, 0, "", nil},
		{"00000000", 0, 1, "", nil},
		{"00616263", 0, 4, "", nil},
		{"00000000", 0, 5, "", io.ErrUnexpectedEOF},
		{"41424361626320202020", ' ', 10, "ABCabc", nil},
		{"e6bca2e5ad97204142430000", 0, 12, "漢字 ABC", nil},
	}
	for _, tc := range tcs {
		tag := fmt.Sprintf("%q, %02x, %d", tc.b, tc.p, tc.n)
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, err := typeio.ReadStringN(r, tc.n, tc.p)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%s: error expected: got %q", tag, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%s: unexpected type of error: got %q, want %q", tag, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%s: unexpected error: %s", tag, err)
		case tc.e == nil && got != tc.s:
			t.Errorf("%s: unexpected read: got %q, want %q", tag, got, tc.s)
		}
	}
}

func TestReadCString(t *testing.T) {
	tcs := []struct {
		b string
		n int
		s string
		e error
	}{
		{"", 0, "", io.EOF},
		{"00", 1, "", nil},
		{"41", 1, "A", nil},
		{"4100", 2, "A", nil},
		{"414243", 3, "ABC", nil},
		{"41424300", 4, "ABC", nil},
		{"4142432061626320", 8, "ABC abc ", nil},
		{"4142432061626320004142", 9, "ABC abc ", nil},
		{"e6bca2e5ad972041424300", 11, "漢字 ABC", nil},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		r := bytes.NewReader(b)
		got, n, err := typeio.ReadCString(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%s: error expected: got %q", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%s: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%s: unexpected error: %s", tc.b, err)
		case tc.e == nil && got != tc.s:
			t.Errorf("%s: unexpected read: got %q, want %q", tc.b, got, tc.s)
		case tc.e == nil && n != tc.n:
			t.Errorf("%s: unexpected read len: got %d, want %d", tc.b, n, tc.n)
		}
	}
}
