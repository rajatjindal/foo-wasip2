// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package udp represents the interface "wasi:sockets/udp@0.2.0".
package udp

import (
	"github.com/rajatjindal/foo-wasip2/wasi/io/poll"
	"github.com/rajatjindal/foo-wasip2/wasi/sockets/network"
	"github.com/ydnar/wasm-tools-go/cm"
)

// ErrorCode represents the enum "wasi:sockets/network@0.2.0#error-code".
//
// See [network.ErrorCode] for more information.
type ErrorCode = network.ErrorCode

// IncomingDatagram represents the record "wasi:sockets/udp@0.2.0#incoming-datagram".
//
// A received datagram.
//
//	record incoming-datagram {
//		data: list<u8>,
//		remote-address: ip-socket-address,
//	}
type IncomingDatagram struct {
	// The payload.
	//
	// Theoretical max size: ~64 KiB. In practice, typically less than 1500 bytes.
	Data cm.List[uint8]

	// The source address.
	//
	// This field is guaranteed to match the remote address the stream was initialized
	// with, if any.
	//
	// Equivalent to the `src_addr` out parameter of `recvfrom`.
	RemoteAddress IPSocketAddress
}

// IncomingDatagramStream represents the resource "wasi:sockets/udp@0.2.0#incoming-datagram-stream".
//
//	resource incoming-datagram-stream
type IncomingDatagramStream cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingDatagramStream) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [resource-drop]incoming-datagram-stream
//go:noescape
func (self IncomingDatagramStream) wasmimport_ResourceDrop()

// Receive represents method "receive".
//
// Receive messages on the socket.
//
// This function attempts to receive up to `max-results` datagrams on the socket without
// blocking.
// The returned list may contain fewer elements than requested, but never more.
//
// This function returns successfully with an empty list when either:
// - `max-results` is 0, or:
// - `max-results` is greater than 0, but no results are immediately available.
// This function never returns `error(would-block)`.
//
// # Typical errors
// - `remote-unreachable`: The remote address is not reachable. (ECONNRESET, ENETRESET
// on Windows, EHOSTUNREACH, EHOSTDOWN, ENETUNREACH, ENETDOWN, ENONET)
// - `connection-refused`: The connection was refused. (ECONNREFUSED)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/recvfrom.html>
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/recvmsg.html>
// - <https://man7.org/linux/man-pages/man2/recv.2.html>
// - <https://man7.org/linux/man-pages/man2/recvmmsg.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-recv>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-recvfrom>
// - <https://learn.microsoft.com/en-us/previous-versions/windows/desktop/legacy/ms741687(v=vs.85)>
// - <https://man.freebsd.org/cgi/man.cgi?query=recv&sektion=2>
//
//	receive: func(max-results: u64) -> result<list<incoming-datagram>, error-code>
//
//go:nosplit
func (self IncomingDatagramStream) Receive(maxResults uint64) cm.OKResult[cm.List[IncomingDatagram], ErrorCode] {
	var result cm.OKResult[cm.List[IncomingDatagram], ErrorCode]
	self.wasmimport_Receive(maxResults, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]incoming-datagram-stream.receive
//go:noescape
func (self IncomingDatagramStream) wasmimport_Receive(maxResults uint64, result *cm.OKResult[cm.List[IncomingDatagram], ErrorCode])

// Subscribe represents method "subscribe".
//
// Create a `pollable` which will resolve once the stream is ready to receive again.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self IncomingDatagramStream) Subscribe() Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]incoming-datagram-stream.subscribe
//go:noescape
func (self IncomingDatagramStream) wasmimport_Subscribe() Pollable

// IPAddressFamily represents the enum "wasi:sockets/network@0.2.0#ip-address-family".
//
// See [network.IPAddressFamily] for more information.
type IPAddressFamily = network.IPAddressFamily

// IPSocketAddress represents the variant "wasi:sockets/network@0.2.0#ip-socket-address".
//
// See [network.IPSocketAddress] for more information.
type IPSocketAddress = network.IPSocketAddress

// Network represents the resource "wasi:sockets/network@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// OutgoingDatagram represents the record "wasi:sockets/udp@0.2.0#outgoing-datagram".
//
// A datagram to be sent out.
//
//	record outgoing-datagram {
//		data: list<u8>,
//		remote-address: option<ip-socket-address>,
//	}
type OutgoingDatagram struct {
	// The payload.
	Data cm.List[uint8]

	// The destination address.
	//
	// The requirements on this field depend on how the stream was initialized:
	// - with a remote address: this field must be None or match the stream's remote address
	// exactly.
	// - without a remote address: this field is required.
	//
	// If this value is None, the send operation is equivalent to `send` in POSIX. Otherwise
	// it is equivalent to `sendto`.
	RemoteAddress cm.Option[IPSocketAddress]
}

// OutgoingDatagramStream represents the resource "wasi:sockets/udp@0.2.0#outgoing-datagram-stream".
//
//	resource outgoing-datagram-stream
type OutgoingDatagramStream cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingDatagramStream) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [resource-drop]outgoing-datagram-stream
//go:noescape
func (self OutgoingDatagramStream) wasmimport_ResourceDrop()

// CheckSend represents method "check-send".
//
// Check readiness for sending. This function never blocks.
//
// Returns the number of datagrams permitted for the next call to `send`,
// or an error. Calling `send` with more datagrams than this function has
// permitted will trap.
//
// When this function returns ok(0), the `subscribe` pollable will
// become ready when this function will report at least ok(1), or an
// error.
//
// Never returns `would-block`.
//
//	check-send: func() -> result<u64, error-code>
//
//go:nosplit
func (self OutgoingDatagramStream) CheckSend() cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_CheckSend(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]outgoing-datagram-stream.check-send
//go:noescape
func (self OutgoingDatagramStream) wasmimport_CheckSend(result *cm.OKResult[uint64, ErrorCode])

// Send represents method "send".
//
// Send messages on the socket.
//
// This function attempts to send all provided `datagrams` on the socket without blocking
// and
// returns how many messages were actually sent (or queued for sending). This function
// never
// returns `error(would-block)`. If none of the datagrams were able to be sent, `ok(0)`
// is returned.
//
// This function semantically behaves the same as iterating the `datagrams` list and
// sequentially
// sending each individual datagram until either the end of the list has been reached
// or the first error occurred.
// If at least one datagram has been sent successfully, this function never returns
// an error.
//
// If the input list is empty, the function returns `ok(0)`.
//
// Each call to `send` must be permitted by a preceding `check-send`. Implementations
// must trap if
// either `check-send` was not called or `datagrams` contains more items than `check-send`
// permitted.
//
// # Typical errors
// - `invalid-argument`:        The `remote-address` has the wrong address family.
// (EAFNOSUPPORT)
// - `invalid-argument`:        The IP address in `remote-address` is set to INADDR_ANY
// (`0.0.0.0` / `::`). (EDESTADDRREQ, EADDRNOTAVAIL)
// - `invalid-argument`:        The port in `remote-address` is set to 0. (EDESTADDRREQ,
// EADDRNOTAVAIL)
// - `invalid-argument`:        The socket is in "connected" mode and `remote-address`
// is `some` value that does not match the address passed to `stream`. (EISCONN)
// - `invalid-argument`:        The socket is not "connected" and no value for `remote-address`
// was provided. (EDESTADDRREQ)
// - `remote-unreachable`:      The remote address is not reachable. (ECONNRESET,
// ENETRESET on Windows, EHOSTUNREACH, EHOSTDOWN, ENETUNREACH, ENETDOWN, ENONET)
// - `connection-refused`:      The connection was refused. (ECONNREFUSED)
// - `datagram-too-large`:      The datagram is too large. (EMSGSIZE)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/sendto.html>
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/sendmsg.html>
// - <https://man7.org/linux/man-pages/man2/send.2.html>
// - <https://man7.org/linux/man-pages/man2/sendmmsg.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-send>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-sendto>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-wsasendmsg>
// - <https://man.freebsd.org/cgi/man.cgi?query=send&sektion=2>
//
//	send: func(datagrams: list<outgoing-datagram>) -> result<u64, error-code>
//
//go:nosplit
func (self OutgoingDatagramStream) Send(datagrams cm.List[OutgoingDatagram]) cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_Send(datagrams, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]outgoing-datagram-stream.send
//go:noescape
func (self OutgoingDatagramStream) wasmimport_Send(datagrams cm.List[OutgoingDatagram], result *cm.OKResult[uint64, ErrorCode])

// Subscribe represents method "subscribe".
//
// Create a `pollable` which will resolve once the stream is ready to send again.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self OutgoingDatagramStream) Subscribe() Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]outgoing-datagram-stream.subscribe
//go:noescape
func (self OutgoingDatagramStream) wasmimport_Subscribe() Pollable

// Pollable represents the resource "wasi:io/poll@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// UDPSocket represents the resource "wasi:sockets/udp@0.2.0#udp-socket".
//
// A UDP socket handle.
//
//	resource udp-socket
type UDPSocket cm.Resource

// ResourceDrop represents the Canonical ABI function "resource-drop".
//
// Drops a resource handle.
//
//go:nosplit
func (self UDPSocket) ResourceDrop() {
	self.wasmimport_ResourceDrop()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [resource-drop]udp-socket
//go:noescape
func (self UDPSocket) wasmimport_ResourceDrop()

// AddressFamily represents method "address-family".
//
// Whether this is a IPv4 or IPv6 socket.
//
// Equivalent to the SO_DOMAIN socket option.
//
//	address-family: func() -> ip-address-family
//
//go:nosplit
func (self UDPSocket) AddressFamily() IPAddressFamily {
	return self.wasmimport_AddressFamily()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.address-family
//go:noescape
func (self UDPSocket) wasmimport_AddressFamily() IPAddressFamily

// FinishBind represents method "finish-bind".
//
//	finish-bind: func() -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) FinishBind() cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_FinishBind(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.finish-bind
//go:noescape
func (self UDPSocket) wasmimport_FinishBind(result *cm.ErrResult[struct{}, ErrorCode])

// LocalAddress represents method "local-address".
//
// Get the current bound address.
//
// POSIX mentions:
// > If the socket has not been bound to a local name, the value
// > stored in the object pointed to by `address` is unspecified.
//
// WASI is stricter and requires `local-address` to return `invalid-state` when the
// socket hasn't been bound yet.
//
// # Typical errors
// - `invalid-state`: The socket is not bound to any local address.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getsockname.html>
// - <https://man7.org/linux/man-pages/man2/getsockname.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-getsockname>
// - <https://man.freebsd.org/cgi/man.cgi?getsockname>
//
//	local-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self UDPSocket) LocalAddress() cm.OKResult[IPSocketAddress, ErrorCode] {
	var result cm.OKResult[IPSocketAddress, ErrorCode]
	self.wasmimport_LocalAddress(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.local-address
//go:noescape
func (self UDPSocket) wasmimport_LocalAddress(result *cm.OKResult[IPSocketAddress, ErrorCode])

// ReceiveBufferSize represents method "receive-buffer-size".
//
// The kernel buffer space reserved for sends/receives on this socket.
//
// If the provided value is 0, an `invalid-argument` error is returned.
// Any other value will never cause an error, but it might be silently clamped and/or
// rounded.
// I.e. after setting a value, reading the same setting back may return a different
// value.
//
// Equivalent to the SO_RCVBUF and SO_SNDBUF socket options.
//
// # Typical errors
// - `invalid-argument`:     (set) The provided value was 0.
//
//	receive-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) ReceiveBufferSize() cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_ReceiveBufferSize(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.receive-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_ReceiveBufferSize(result *cm.OKResult[uint64, ErrorCode])

// RemoteAddress represents method "remote-address".
//
// Get the address the socket is currently streaming to.
//
// # Typical errors
// - `invalid-state`: The socket is not streaming to a specific remote address. (ENOTCONN)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/getpeername.html>
// - <https://man7.org/linux/man-pages/man2/getpeername.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-getpeername>
// - <https://man.freebsd.org/cgi/man.cgi?query=getpeername&sektion=2&n=1>
//
//	remote-address: func() -> result<ip-socket-address, error-code>
//
//go:nosplit
func (self UDPSocket) RemoteAddress() cm.OKResult[IPSocketAddress, ErrorCode] {
	var result cm.OKResult[IPSocketAddress, ErrorCode]
	self.wasmimport_RemoteAddress(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.remote-address
//go:noescape
func (self UDPSocket) wasmimport_RemoteAddress(result *cm.OKResult[IPSocketAddress, ErrorCode])

// SendBufferSize represents method "send-buffer-size".
//
//	send-buffer-size: func() -> result<u64, error-code>
//
//go:nosplit
func (self UDPSocket) SendBufferSize() cm.OKResult[uint64, ErrorCode] {
	var result cm.OKResult[uint64, ErrorCode]
	self.wasmimport_SendBufferSize(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.send-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_SendBufferSize(result *cm.OKResult[uint64, ErrorCode])

// SetReceiveBufferSize represents method "set-receive-buffer-size".
//
//	set-receive-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetReceiveBufferSize(value uint64) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetReceiveBufferSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.set-receive-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_SetReceiveBufferSize(value uint64, result *cm.ErrResult[struct{}, ErrorCode])

// SetSendBufferSize represents method "set-send-buffer-size".
//
//	set-send-buffer-size: func(value: u64) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetSendBufferSize(value uint64) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetSendBufferSize(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.set-send-buffer-size
//go:noescape
func (self UDPSocket) wasmimport_SetSendBufferSize(value uint64, result *cm.ErrResult[struct{}, ErrorCode])

// SetUnicastHopLimit represents method "set-unicast-hop-limit".
//
//	set-unicast-hop-limit: func(value: u8) -> result<_, error-code>
//
//go:nosplit
func (self UDPSocket) SetUnicastHopLimit(value uint8) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_SetUnicastHopLimit(value, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.set-unicast-hop-limit
//go:noescape
func (self UDPSocket) wasmimport_SetUnicastHopLimit(value uint8, result *cm.ErrResult[struct{}, ErrorCode])

// StartBind represents method "start-bind".
//
// Bind the socket to a specific network on the provided IP address and port.
//
// If the IP address is zero (`0.0.0.0` in IPv4, `::` in IPv6), it is left to the
// implementation to decide which
// network interface(s) to bind to.
// If the port is zero, the socket will be bound to a random free port.
//
// # Typical errors
// - `invalid-argument`:          The `local-address` has the wrong address family.
// (EAFNOSUPPORT, EFAULT on Windows)
// - `invalid-state`:             The socket is already bound. (EINVAL)
// - `address-in-use`:            No ephemeral ports available. (EADDRINUSE, ENOBUFS
// on Windows)
// - `address-in-use`:            Address is already in use. (EADDRINUSE)
// - `address-not-bindable`:      `local-address` is not an address that the `network`
// can bind to. (EADDRNOTAVAIL)
// - `not-in-progress`:           A `bind` operation is not in progress.
// - `would-block`:               Can't finish the operation, it is still in progress.
// (EWOULDBLOCK, EAGAIN)
//
// # Implementors note
// Unlike in POSIX, in WASI the bind operation is async. This enables
// interactive WASI hosts to inject permission prompts. Runtimes that
// don't want to make use of this ability can simply call the native
// `bind` as part of either `start-bind` or `finish-bind`.
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/bind.html>
// - <https://man7.org/linux/man-pages/man2/bind.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock/nf-winsock-bind>
// - <https://man.freebsd.org/cgi/man.cgi?query=bind&sektion=2&format=html>
//
//	start-bind: func(network: borrow<network>, local-address: ip-socket-address) ->
//	result<_, error-code>
//
//go:nosplit
func (self UDPSocket) StartBind(network_ Network, localAddress IPSocketAddress) cm.ErrResult[struct{}, ErrorCode] {
	var result cm.ErrResult[struct{}, ErrorCode]
	self.wasmimport_StartBind(network_, localAddress, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.start-bind
//go:noescape
func (self UDPSocket) wasmimport_StartBind(network_ Network, localAddress IPSocketAddress, result *cm.ErrResult[struct{}, ErrorCode])

// Stream represents method "stream".
//
// Set up inbound & outbound communication channels, optionally to a specific peer.
//
// This function only changes the local socket configuration and does not generate
// any network traffic.
// On success, the `remote-address` of the socket is updated. The `local-address`
// may be updated as well,
// based on the best network path to `remote-address`.
//
// When a `remote-address` is provided, the returned streams are limited to communicating
// with that specific peer:
// - `send` can only be used to send to this destination.
// - `receive` will only return datagrams sent from the provided `remote-address`.
//
// This method may be called multiple times on the same socket to change its association,
// but
// only the most recently returned pair of streams will be operational. Implementations
// may trap if
// the streams returned by a previous invocation haven't been dropped yet before calling
// `stream` again.
//
// The POSIX equivalent in pseudo-code is:
//
//	if (was previously connected) {
//	connect(s, AF_UNSPEC)
//	}
//	if (remote_address is Some) {
//	connect(s, remote_address)
//	}
//
// Unlike in POSIX, the socket must already be explicitly bound.
//
// # Typical errors
// - `invalid-argument`:          The `remote-address` has the wrong address family.
// (EAFNOSUPPORT)
// - `invalid-argument`:          The IP address in `remote-address` is set to INADDR_ANY
// (`0.0.0.0` / `::`). (EDESTADDRREQ, EADDRNOTAVAIL)
// - `invalid-argument`:          The port in `remote-address` is set to 0. (EDESTADDRREQ,
// EADDRNOTAVAIL)
// - `invalid-state`:             The socket is not bound.
// - `address-in-use`:            Tried to perform an implicit bind, but there were
// no ephemeral ports available. (EADDRINUSE, EADDRNOTAVAIL on Linux, EAGAIN on BSD)
// - `remote-unreachable`:        The remote address is not reachable. (ECONNRESET,
// ENETRESET, EHOSTUNREACH, EHOSTDOWN, ENETUNREACH, ENETDOWN, ENONET)
// - `connection-refused`:        The connection was refused. (ECONNREFUSED)
//
// # References
// - <https://pubs.opengroup.org/onlinepubs/9699919799/functions/connect.html>
// - <https://man7.org/linux/man-pages/man2/connect.2.html>
// - <https://learn.microsoft.com/en-us/windows/win32/api/winsock2/nf-winsock2-connect>
// - <https://man.freebsd.org/cgi/man.cgi?connect>
//
//	stream: func(remote-address: option<ip-socket-address>) -> result<tuple<incoming-datagram-stream,
//	outgoing-datagram-stream>, error-code>
//
//go:nosplit
func (self UDPSocket) Stream(remoteAddress cm.Option[IPSocketAddress]) cm.OKResult[cm.Tuple[IncomingDatagramStream, OutgoingDatagramStream], ErrorCode] {
	var result cm.OKResult[cm.Tuple[IncomingDatagramStream, OutgoingDatagramStream], ErrorCode]
	self.wasmimport_Stream(remoteAddress, &result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.stream
//go:noescape
func (self UDPSocket) wasmimport_Stream(remoteAddress cm.Option[IPSocketAddress], result *cm.OKResult[cm.Tuple[IncomingDatagramStream, OutgoingDatagramStream], ErrorCode])

// Subscribe represents method "subscribe".
//
// Create a `pollable` which will resolve once the socket is ready for I/O.
//
// Note: this function is here for WASI Preview2 only.
// It's planned to be removed when `future` is natively supported in Preview3.
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self UDPSocket) Subscribe() Pollable {
	return self.wasmimport_Subscribe()
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.subscribe
//go:noescape
func (self UDPSocket) wasmimport_Subscribe() Pollable

// UnicastHopLimit represents method "unicast-hop-limit".
//
// Equivalent to the IP_TTL & IPV6_UNICAST_HOPS socket options.
//
// If the provided value is 0, an `invalid-argument` error is returned.
//
// # Typical errors
// - `invalid-argument`:     (set) The TTL value must be 1 or higher.
//
//	unicast-hop-limit: func() -> result<u8, error-code>
//
//go:nosplit
func (self UDPSocket) UnicastHopLimit() cm.OKResult[uint8, ErrorCode] {
	var result cm.OKResult[uint8, ErrorCode]
	self.wasmimport_UnicastHopLimit(&result)
	return result
}

//go:wasmimport wasi:sockets/udp@0.2.0 [method]udp-socket.unicast-hop-limit
//go:noescape
func (self UDPSocket) wasmimport_UnicastHopLimit(result *cm.OKResult[uint8, ErrorCode])

type Interface interface {
}

var instance Interface
