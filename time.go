// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio

import (
	"io"
	"time"
)

// ReadUnixTimeUTC32BE reads 4 bytes in big-endian byte order from r, interprets
// it as a UNIX time, the number of seconds elapsed since Jan 1, 1970 UTC, and
// returns the UTC time it represents.
// Note that this data type has the well-known Y2038 problem.
func ReadUnixTimeUTC32BE(r io.Reader) (time.Time, error) {
	t, err := ReadUint32BE(r)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(t), 0).UTC(), nil
}

// ReadUnixTime32BE is identical to ReadUnixTimeUTC32BE except that it returns
// the local time rather than UTC.
// Note that this data type has the well-known Y2038 problem.
func ReadUnixTime32BE(r io.Reader) (time.Time, error) {
	t, err := ReadUint32BE(r)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(t), 0).Local(), nil
}

// ReadUnixTimeUTC32LE reads 4 bytes in little-endian byte order from r,
// interprets it as a UNIX time, the number of seconds elapsed since Jan 1, 1970
// UTC, and returns the UTC time it represents.
// Note that this data type has the well-known Y2038 problem.
func ReadUnixTimeUTC32LE(r io.Reader) (time.Time, error) {
	t, err := ReadUint32LE(r)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(t), 0).UTC(), nil
}

// ReadUnixTime32LE is identical to ReadUnixTimeUTC32LE except that it returns
// the local time rather than UTC.
// Note that this data type has the well-known Y2038 problem.
func ReadUnixTime32LE(r io.Reader) (time.Time, error) {
	t, err := ReadUint32LE(r)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(int64(t), 0).Local(), nil
}

// WriteUnixTime32BE writes 4 bytes to w that represent the UNIX time for t, the
// number of seconds elapsed since Jan 1, 1970 UTC, in big-endian byte order.
// The written bytes do not depend on the location associated with t. The
// appropriate UTC value for t is always used, regardless of the associated
// location.
// Note that this data type has the well-known Y2038 problem. Time values before
// the 1970 epoch time or afte the Y2038 are not written correctly.
func WriteUnixTime32BE(w io.Writer, t time.Time) error {
	return WriteUint32BE(w, uint32(t.Unix()))
}

// WriteUnixTime32LE writes 4 bytes to w that represent the UNIX time for t, the
// number of seconds elapsed since Jan 1, 1970 UTC, in little-endian byte order.
// The written bytes do not depend on the location associated with t. The
// appropriate UTC value for t is always used, regardless of the associated
// location.
// Note that this data type has the well-known Y2038 problem. Time values before
// the 1970 epoch time or afte the Y2038 are not written correctly.
func WriteUnixTime32LE(w io.Writer, t time.Time) error {
	return WriteUint32LE(w, uint32(t.Unix()))
}
