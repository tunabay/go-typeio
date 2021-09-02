// Copyright (c) 2021 Hirotsuna Mizuno. All rights reserved.
// Use of this source code is governed by the MIT license that can be found in
// the LICENSE file.

package typeio

import (
	"errors"
	"io"
	"net"
)

// ErrInvalidIP is the error thrown when an invalid IP address is specified.
var ErrInvalidIP = errors.New("invalid IP address")

// ReadIPv4 reads 4 bytes from r and returns them as an IPv4 address.
func ReadIPv4(r io.Reader) (net.IP, error) {
	b, err := readN(r, net.IPv4len)
	if err != nil {
		return nil, err
	}
	return net.IP(b), nil
}

// WriteIPv4 writes 4 bytes to w that represents the IPv4 address addr.
func WriteIPv4(w io.Writer, addr net.IP) error {
	ip4 := addr.To4()
	if ip4 == nil {
		return ErrInvalidIP
	}
	return write(w, ip4)
}

// ReadIPv6 reads 16 bytes from r and returns them as an IPv6 address.
func ReadIPv6(r io.Reader) (net.IP, error) {
	b, err := readN(r, net.IPv6len)
	if err != nil {
		return nil, err
	}
	return net.IP(b), nil
}

// WriteIPv6 writes 16 bytes to w that represents the IPv6 address addr. The
// addr can be an IPv4 address, and it will be written as an IPv4-mapped IPv6.
func WriteIPv6(w io.Writer, addr net.IP) error {
	ip16 := addr.To16()
	if ip16 == nil {
		return ErrInvalidIP
	}
	return write(w, ip16)
}
