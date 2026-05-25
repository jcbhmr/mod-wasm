//go:build wasip2

package main

import (
	modfilewit "go.jcbhmr.com.internal/mod-wasm/internal/jcbhmr/mod/modfile"
	"golang.org/x/mod/modfile"
)

func init() {
	modfilewit.Exports.GetGoVersionRE = func() {
		_ = modfile.GoVersionRE
	}
	modfilewit.Exports.GetToolchainRE = func() {
		_ = modfile.ToolchainRE
	}
	modfilewit.Exports.AutoQuote = func(s string) string {
		return modfile.AutoQuote(s)
	}
}
