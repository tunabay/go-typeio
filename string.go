// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

// ReadStringN reads exactly n bytes from r and returns it as a string with the
// first pad and the rest removed. Generally, it is expected to specify 0 (null
// character) or ' ' (space) as pad.
func ReadStringN(r io.Reader, n int, pad byte) (string, error) {
	b, err := readN(r, n)
	if err != nil {
		return "", err
	}
	if p := bytes.IndexByte(b, pad); 0 <= p {
		b = b[:p]
	}
	return string(b), nil
}

// ReadCString keeps reading from r until it encounters a null character
// represented by 0x00 or io.EOF, and returns the part before the null character
// as a string. The second return value is the total number of bytes read
// including the null character. Be careful of overruns when using ReadCString
// for unpredictable inputs.
func ReadCString(r io.Reader) (string, int, error) {
	var b []byte
	var t int
	buf := []byte{0}
	for {
		n, err := r.Read(buf)
		if n == 1 {
			t++
			if buf[0] == 0 {
				break
			}
			b = append(b, buf[0])
		}
		if err != nil {
			if errors.Is(err, io.EOF) {
				if 0 < t {
					return string(b), t, nil
				}
				return string(b), t, io.EOF
			}
			return string(b), t, fmt.Errorf("read failure: %w", err)
		}
	}
	return string(b), t, nil
}
