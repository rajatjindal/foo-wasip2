// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package stdout represents the interface "wasi:cli/stdout@0.2.0".
package stdout

import (
	"github.com/rajatjindal/foo-wasip2/wasi/io/streams"
)

// OutputStream represents the resource "wasi:io/streams@0.2.0#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// GetStdout represents function "wasi:cli/stdout@0.2.0#get-stdout".
//
//	get-stdout: func() -> output-stream
//
//go:nosplit
func GetStdout() OutputStream {
	return wasmimport_GetStdout()
}

//go:wasmimport wasi:cli/stdout@0.2.0 get-stdout
//go:noescape
func wasmimport_GetStdout() OutputStream

type Interface interface {
	GetStdout() OutputStream
}

var instance Interface
