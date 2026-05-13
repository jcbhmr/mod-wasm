//go:build wasip2

package main

import (
	semverwasm "go.jcbhmr.com.internal/mod-wasm/internal/wit-bindgen-go/jcbhmr/mod/semver"
	"golang.org/x/mod/semver"
)

func init() {
	semverwasm.Exports.Build = func(v string) (result string) {
		return semver.Build(v)
	}
	semverwasm.Exports.Canonical = func(v string) (result string) {
		return semver.Canonical(v)
	}
}
