// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package instancenetwork represents the interface "wasi:sockets/instance-network@0.2.0".
//
// This interface provides a value-export of the default network handle..
package instancenetwork

import (
	"github.com/rajatjindal/foo-wasip2/wasi/sockets/network"
)

// Network represents the resource "wasi:sockets/network@0.2.0#network".
//
// See [network.Network] for more information.
type Network = network.Network

// InstanceNetwork represents function "wasi:sockets/instance-network@0.2.0#instance-network".
//
// Get a handle to the default network.
//
//	instance-network: func() -> network
//
//go:nosplit
func InstanceNetwork() Network {
	return wasmimport_InstanceNetwork()
}

//go:wasmimport wasi:sockets/instance-network@0.2.0 instance-network
//go:noescape
func wasmimport_InstanceNetwork() Network

type Interface interface {
	InstanceNetwork() Network
}

var instance Interface
