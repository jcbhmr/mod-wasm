//go:build wasip2

package main

import (
	"go.bytecodealliance.org/cm"
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
	semverwit.Exports.Compare = func(v, w string) (result int32) {
		rGo := semver.Compare(v, w)
		return int32(rGo)
	}
	semverwit.Exports.IsValid = func(v string) (result bool) {
		return semver.IsValid(v)
	}
	semverwit.Exports.Major = func(v string) (result string) {
		return semver.Major(v)
	}
	semverwit.Exports.MajorMinor = func(v string) (result string) {
		return semver.MajorMinor(v)
	}
	semverwit.Exports.Max = func(v, w string) (result string) {
		return semver.Max(v, w)
	}
	semverwit.Exports.Prerelease = func(v string) (result string) {
		return semver.Prerelease(v)
	}
	semverwit.Exports.Sort = func(list cm.List[string]) (result cm.List[string]) {
		listGo := list.Slice()
		semver.Sort(listGo)
		return cm.ToList(listGo)
	}
}
