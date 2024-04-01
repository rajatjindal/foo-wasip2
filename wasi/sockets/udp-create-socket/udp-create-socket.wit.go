// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package udpcreatesocket represents the interface "wasi:sockets/udp-create-socket@0.2.0".
package udpcreatesocket

import (
	"github.com/rajatjindal/foo-wasip2/wasi/sockets/network"
	"github.com/rajatjindal/foo-wasip2/wasi/sockets/udp"
	"github.com/ydnar/wasm-tools-go/cm"
)

// ErrorCode represents the enum "wasi:sockets/network@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IPAddressFamily represents the enum "wasi:sockets/network@0.2.0#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// Network represents the resource "wasi:sockets/network@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// UDPSocket represents the resource "wasi:sockets/udp@0.2.0#udp-socket".
//
// See [udp.UDPSocket] for more information.
type UDPSocket = udp.UDPSocket

// CreateUDPSocket represents function "wasi:sockets/udp-create-socket@0.2.0#create-udp-socket".
//
// Create a new UDP socket.
//
// Similar to `socket(AF_INET or AF_INET6, SOCK_DGRAM, IPPROTO_UDP)` in POSIX.
// On IPv6 sockets, IPV6_V6ONLY is enabled by default and can't be configured otherwise.
//
// This function does not require a network capability handle. This is considered
// to be safe because
// at time of creation, the socket is not bound to any `network` yet. Up to the moment
// `bind` is called,
// the socket is effectively an in-memory configuration object, unable to communicate
// with the outside world.
//
// All sockets are non-blocking. Use the wasi-poll interface to block on asynchronous
// operations.
//
// # Typical errors
// - `not-supported`:     The specified `address-family` is not supported. (EAFNOSUPPORT)
// - `new-socket-limit`:  The new socket resource could not be created because of
// a system limit. (EMFILE, ENFILE)
//
// # References:
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/socket.html>
// - <https://man7.org/linux/man-pages/man2/socket.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-wsasocketw>
// - <https://man.freebsd.org/cgi/man.cgi?query=socket&sektion=2>
//
//	create-udp-socket: func(address-family: ip-address-family) -> result<own<udp-socket>,
//	error-code>
//
//go:nosplit
func CreateUDPSocket(addressFamily IPAddressFamily) cm.OKResult[UDPSocket, ErrorCode] {
	var result cm.OKResult[UDPSocket, ErrorCode]
	wasmimport_CreateUDPSocket(addressFamily, &result)
	return result
}

//go:wasmimport wasi:sockets/udp-create-socket@0.2.0 create-udp-socket
//go:noescape
func wasmimport_CreateUDPSocket(addressFamily IPAddressFamily, result *cm.OKResult[UDPSocket, ErrorCode])

type Interface interface {
	CreateUDPSocket(addressFamily IPAddressFamily) cm.OKResult[UDPSocket, ErrorCode]
}

var instance Interface
