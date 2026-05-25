//go:build wasip2

package main

import (
	semverwit "go.jcbhmr.com.internal/mod-wasm/internal/jcbhmr/mod/semver"
	"golang.org/x/mod/semver"
)

func init() {
	semverwit.Exports.Build = func(v string) (result string) {
		return semver.Build(v)
	}
	semverwit.Exports.Canonical = func(v string) (result string) {
		return semver.Canonical(v)
	}
}
