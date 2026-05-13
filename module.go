//go:build wasip2

package main

import (
	modulewasm "go.jcbhmr.com.internal/mod-wasm/internal/wit-bindgen-go/jcbhmr/mod/module"
	"golang.org/x/mod/module"
)

func init() {
	modulewasm.Exports.VersionString = func(self modulewasm.Version) (result string) {
		selfGo := module.Version{Path: self.Path, Version: self.Version}
		return selfGo.String()
	}
}
