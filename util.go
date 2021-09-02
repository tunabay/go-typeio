// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio

import (
	"errors"
	"fmt"
	"io"
)

func readN(r io.Reader, n int) ([]byte, error) {
	b := make([]byte, n)
	if _, err := io.ReadFull(r, b); err != nil {
		if errors.Is(err, io.EOF) {
			return nil, io.EOF
		}
		return nil, fmt.Errorf("read failure: %w", err)
	}
	return b, nil
}

func write(w io.Writer, b []byte) error {
	if _, err := w.Write(b); err != nil {
		return fmt.Errorf("write failure: %w", err)
	}
	return nil
}
