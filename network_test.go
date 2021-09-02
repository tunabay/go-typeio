// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio_test

import (
	"bytes"
	"encoding/hex"
	"errors"
	"io"
	"net"
	"testing"

	"github.com/tunabay/go-typeio"
)

func TestReadIPv4(t *testing.T) {
	tcs := []struct {
		b, a string
		e    error
	}{
		{"00000000", "0.0.0.0", nil},
		{"c0a8ff00", "192.168.255.0", nil},
		{"ffffffff", "255.255.255.255", nil},
		{"", "", io.EOF},
		{"c0", "", io.ErrUnexpectedEOF},
		{"c0a8", "", io.ErrUnexpectedEOF},
		{"c0a8ff", "", io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		want := net.ParseIP(tc.a)
		r := bytes.NewReader(b)
		got, err := typeio.ReadIPv4(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %s", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !got.Equal(want):
			t.Errorf("%q: unexpected read: got %s, want %s", tc.b, got, want)
		}
	}
}

func TestWriteIPv4(t *testing.T) {
	tcs := []struct {
		a net.IP
		b string
		e error
	}{
		{net.ParseIP("0.0.0.0"), "00000000", nil},
		{net.ParseIP("192.168.255.0"), "c0a8ff00", nil},
		{net.ParseIP("::ffff:10.255.255.255"), "0affffff", nil},
		{net.ParseIP("fd12::3456:7890"), "", typeio.ErrInvalidIP},
		{net.IP([]byte{0, 0, 0, 0, 0}), "", typeio.ErrInvalidIP},
		{net.IP([]byte{0, 0, 0}), "", typeio.ErrInvalidIP},
		{net.IP([]byte{0}), "", typeio.ErrInvalidIP},
		{nil, "", typeio.ErrInvalidIP},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		err := typeio.WriteIPv4(w, tc.a)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%s: error expected.", tc.a)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%s: unexpected type of error: got %q, want %q", tc.a, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%s: unexpected error: %s", tc.a, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%s: unexpected write: got %s, want %s", tc.a, got, tc.b)
		}
	}
}

func TestReadIPv6(t *testing.T) {
	tcs := []struct {
		b, a string
		e    error
	}{
		{"00000000000000000000000000000000", "::", nil},
		{"00000000000000000000000000000001", "::1", nil},
		{
			"ffffffffffffffffffffffffffffffff",
			"ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff",
			nil,
		},
		{"20010db8000000000000000001020304", "2001:db8::102:304", nil},
		{"", "", io.EOF},
		{"20", "", io.ErrUnexpectedEOF},
		{"20010db80000000000000000010203", "", io.ErrUnexpectedEOF},
	}
	for _, tc := range tcs {
		b, err := hex.DecodeString(tc.b)
		if err != nil {
			t.Errorf("%q: invalid test data: %s", tc.b, err)
			continue
		}
		want := net.ParseIP(tc.a)
		r := bytes.NewReader(b)
		got, err := typeio.ReadIPv6(r)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%q: error expected: got %s", tc.b, got)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%q: unexpected type of error: got %q, want %q", tc.b, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%q: unexpected error: %s", tc.b, err)
		case tc.e == nil && !got.Equal(want):
			t.Errorf("%q: unexpected read: got %s, want %s", tc.b, got, want)
		}
	}
}

func TestWriteIPv6(t *testing.T) {
	tcs := []struct {
		a net.IP
		b string
		e error
	}{
		{net.ParseIP("::"), "00000000000000000000000000000000", nil},
		{net.ParseIP("::1"), "00000000000000000000000000000001", nil},
		{
			net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"),
			"ffffffffffffffffffffffffffffffff",
			nil,
		},
		{
			net.ParseIP("192.168.255.0"),
			"00000000000000000000ffffc0a8ff00",
			nil,
		},
		{
			net.ParseIP("fd12::3456:7890"),
			"fd120000000000000000000034567890",
			nil,
		},
		{net.IP(bytes.Repeat([]byte{0}, 15)), "", typeio.ErrInvalidIP},
		{net.IP(bytes.Repeat([]byte{0}, 14)), "", typeio.ErrInvalidIP},
		{
			net.IP(bytes.Repeat([]byte{0}, 4)),
			"00000000000000000000ffff00000000",
			nil,
		},
		{net.IP([]byte{0, 0, 0}), "", typeio.ErrInvalidIP},
		{net.IP([]byte{0}), "", typeio.ErrInvalidIP},
		{nil, "", typeio.ErrInvalidIP},
	}
	for _, tc := range tcs {
		w := new(bytes.Buffer)
		err := typeio.WriteIPv6(w, tc.a)
		switch {
		case tc.e != nil && err == nil:
			t.Errorf("%s: error expected.", tc.a)
		case tc.e != nil && !errors.Is(err, tc.e):
			t.Errorf("%s: unexpected type of error: got %q, want %q", tc.a, err, tc.e)
		case tc.e == nil && err != nil:
			t.Errorf("%s: unexpected error: %s", tc.a, err)
		}
		if got := hex.EncodeToString(w.Bytes()); got != tc.b {
			t.Errorf("%s: unexpected write: got %s, want %s", tc.a, got, tc.b)
		}
	}
}
