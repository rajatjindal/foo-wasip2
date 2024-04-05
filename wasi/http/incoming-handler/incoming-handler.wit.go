// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package incominghandler represents the interface "wasi:http/incoming-handler@0.2.0".
//
// This interface defines a handler of incoming HTTP Requests. It should
// be exported by components which can respond to HTTP Requests.
package incominghandler

import (
	"github.com/rajatjindal/foo-wasip2/wasi/http/types"
)

// IncomingRequest represents the resource "wasi:http/types@0.2.0#incoming-request".
//
// See [types.IncomingRequest] for more information.
type IncomingRequest = types.IncomingRequest

// ResponseOutparam represents the resource "wasi:http/types@0.2.0#response-outparam".
//
// See [types.ResponseOutparam] for more information.
type ResponseOutparam = types.ResponseOutparam

// Handle represents function "wasi:http/incoming-handler@0.2.0#handle".
//
// This function is invoked with an incoming HTTP Request, and a resource
// `response-outparam` which provides the capability to reply with an HTTP
// Response. The response is sent by calling the `response-outparam.set`
// method, which allows execution to continue after the response has been
// sent. This enables both streaming to the response body, and performing other
// work.
//
// The implementor of this function must write a response to the
// `response-outparam` before returning, or else the caller will respond
// with an error on its behalf.
//
//	handle: func(request: incoming-request, response-out: response-outparam)
//
//go:nosplit
func Handle(request IncomingRequest, responseOut ResponseOutparam) {
	wasmimport_Handle(request, responseOut)
}

//go:wasmimport wasi:http/incoming-handler@0.2.0 handle
//go:noescape
func wasmimport_Handle(request IncomingRequest, responseOut ResponseOutparam)

type Interface interface {
	Handle(request IncomingRequest, responseOut ResponseOutparam)
}

var instance Interface
