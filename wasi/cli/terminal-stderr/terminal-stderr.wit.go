// Code generated by wit-bindgen-go. DO NOT EDIT.

//go:build !wasip1

// Package terminalstderr represents the interface "wasi:cli/terminal-stderr@0.2.0".
//
// An interface providing an optional `terminal-output` for stderr as a
// link-time authority.
package terminalstderr

import (
	terminaloutput "github.com/rajatjindal/foo-wasip2/wasi/cli/terminal-output"
	"github.com/ydnar/wasm-tools-go/cm"
)

// TerminalOutput represents the resource "wasi:cli/terminal-output@0.2.0#terminal-output".
//
// See [terminaloutput.TerminalOutput] for more information.
type TerminalOutput = terminaloutput.TerminalOutput

// GetTerminalStderr represents function "wasi:cli/terminal-stderr@0.2.0#get-terminal-stderr".
//
// If stderr is connected to a terminal, return a `terminal-output` handle
// allowing further interaction with it.
//
//	get-terminal-stderr: func() -> option<terminal-output>
//
//go:nosplit
func GetTerminalStderr() cm.Option[TerminalOutput] {
	var result cm.Option[TerminalOutput]
	wasmimport_GetTerminalStderr(&result)
	return result
}

//go:wasmimport wasi:cli/terminal-stderr@0.2.0 get-terminal-stderr
//go:noescape
func wasmimport_GetTerminalStderr(result *cm.Option[TerminalOutput])

type Interface interface {
	GetTerminalStderr() cm.Option[TerminalOutput]
}

var instance Interface
