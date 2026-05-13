//go:build wasip2

package main

import (
	modfilewasm "go.jcbhmr.com.internal/mod-wasm/internal/wit-bindgen-go/jcbhmr/mod/modfile"
	"golang.org/x/mod/modfile"
)

func init() {
	modfilewasm.Exports.GetGoVersionRE = func() {
		_ = modfile.GoVersionRE
	}
	modfilewasm.Exports.GetToolchainRE = func() {
		_ = modfile.ToolchainRE
	}
	modfilewasm.Exports.AutoQuote = func(s string) string {
		return modfile.AutoQuote(s)
	}
}
