// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package outgoinghandler represents the interface "wasi:http/outgoing-handler@0.2.0".
//
// This interface defines a handler of outgoing HTTP Requests. It should be
// imported by components which wish to make HTTP Requests.
package outgoinghandler

import (
	"github.com/rajatjindal/foo-wasip2/wasi/http/types"
	"github.com/ydnar/wasm-tools-go/cm"
)

// ErrorCode represents the variant "wasi:http/types@0.2.0#error-code".
//
// See [types.ErrorCode] for more information.
type ErrorCode = types.ErrorCode

// FutureIncomingResponse represents the resource "wasi:http/types@0.2.0#future-incoming-response".
//
// See [types.FutureIncomingResponse] for more information.
type FutureIncomingResponse = types.FutureIncomingResponse

// OutgoingRequest represents the resource "wasi:http/types@0.2.0#outgoing-request".
//
// See [types.OutgoingRequest] for more information.
type OutgoingRequest = types.OutgoingRequest

// RequestOptions represents the resource "wasi:http/types@0.2.0#request-options".
//
// See [types.RequestOptions] for more information.
type RequestOptions = types.RequestOptions

// Handle represents function "wasi:http/outgoing-handler@0.2.0#handle".
//
// This function is invoked with an outgoing HTTP Request, and it returns
// a resource `future-incoming-response` which represents an HTTP Response
// which may arrive in the future.
//
// The `options` argument accepts optional parameters for the HTTP
// protocol's transport layer.
//
// This function may return an error if the `outgoing-request` is invalid
// or not allowed to be made. Otherwise, protocol errors are reported
// through the `future-incoming-response`.
//
//	handle: func(request: outgoing-request, options: option<request-options>) -> result<future-incoming-response,
//	error-code>
//
//go:nosplit
func Handle(request OutgoingRequest, options cm.Option[RequestOptions]) cm.ErrResult[FutureIncomingResponse, ErrorCode] {
	var result cm.ErrResult[FutureIncomingResponse, ErrorCode]
	wasmimport_Handle(request, options, &result)
	return result
}

//go:wasmimport wasi:http/outgoing-handler@0.2.0 handle
//go:noescape
func wasmimport_Handle(request OutgoingRequest, options cm.Option[RequestOptions], result *cm.ErrResult[FutureIncomingResponse, ErrorCode])

type Interface interface {
	Handle(request OutgoingRequest, options cm.Option[RequestOptions]) cm.ErrResult[FutureIncomingResponse, ErrorCode]
}

var instance Interface
