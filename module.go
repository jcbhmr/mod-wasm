//go:build wasip2

package main

import (
	"errors"
	"time"

	"go.bytecodealliance.org/cm"
	modulewit "go.jcbhmr.com.internal/mod-wasm/internal/jcbhmr/mod/module"
	"golang.org/x/mod/module"
)

func init() {
	modulewit.Exports.PseudoVersionTimestampFormat = func(t int64) (result string) {
		tGo := time.Unix(t, 0)
		return tGo.Format(module.PseudoVersionTimestampFormat)
	}
	modulewit.Exports.CanonicalVersion = func(v string) (result string) {
		return module.CanonicalVersion(v)
	}
	modulewit.Exports.Check = func(path, version string) (result cm.Option[modulewit.Error]) {
		rGo := module.Check(path, version)
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	modulewit.Exports.CheckFilePath = func(path string) (result cm.Option[modulewit.Error]) {
		rGo := module.CheckFilePath(path)
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	modulewit.Exports.CheckImportPath = func(path string) (result cm.Option[modulewit.Error]) {
		rGo := module.CheckImportPath(path)
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	modulewit.Exports.CheckPath = func(path string) (result cm.Option[modulewit.Error]) {
		rGo := module.CheckPath(path)
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	modulewit.Exports.CheckPathMajor = func(v, pathMajor string) (result cm.Option[modulewit.Error]) {
		rGo := module.CheckPathMajor(v, pathMajor)
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	modulewit.Exports.EscapePath = func(path string) (result cm.Result[string, string, modulewit.Error]) {
		rGo, errGo := module.EscapePath(path)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo)
		return
	}
	modulewit.Exports.EscapeVersion = func(v string) (result cm.Result[string, string, modulewit.Error]) {
		rGo, errGo := module.EscapeVersion(v)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo)
		return
	}
	modulewit.Exports.IsPseudoVersion = func(v string) (result bool) {
		return module.IsPseudoVersion(v)
	}
	modulewit.Exports.IsZeroPseudoVersion = func(v string) (result bool) {
		return module.IsZeroPseudoVersion(v)
	}
	modulewit.Exports.MatchPathMajor = func(v, pathMajor string) (result bool) {
		return module.MatchPathMajor(v, pathMajor)
	}
	modulewit.Exports.MatchPrefixPatterns = func(globs, target string) (result bool) {
		return module.MatchPrefixPatterns(globs, target)
	}
	modulewit.Exports.PathMajorPrefix = func(pathMajor string) (result string) {
		return module.PathMajorPrefix(pathMajor)
	}
	modulewit.Exports.PseudoVersion = func(major, older string, t int64, rev string) (result string) {
		tGo := time.Unix(t, 0)
		return module.PseudoVersion(major, older, tGo, rev)
	}
	modulewit.Exports.PseudoVersionBase = func(v string) (result cm.Result[string, string, modulewit.Error]) {
		rGo, errGo := module.PseudoVersionBase(v)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo)
		return
	}
	modulewit.Exports.PseudoVersionRev = func(v string) (result cm.Result[string, string, modulewit.Error]) {
		rGo, errGo := module.PseudoVersionRev(v)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo)
		return
	}
	modulewit.Exports.PseudoVersionTime = func(v string) (result cm.Result[string, int64, modulewit.Error]) {
		rGo, errGo := module.PseudoVersionTime(v)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo.Unix())
		return
	}
	modulewit.Exports.Sort = func(list cm.List[modulewit.Version]) (result cm.List[modulewit.Version]) {
		listGo := make([]module.Version, list.Len())
		for i, v := range list.Slice() {
			listGo[i] = module.Version{Path: v.Path, Version: v.Version}
		}
		module.Sort(listGo)
		sorted := make([]modulewit.Version, len(listGo))
		for i, v := range listGo {
			sorted[i] = modulewit.Version{Path: v.Path, Version: v.Version}
		}
		return cm.ToList(sorted)
	}
	modulewit.Exports.SplitPathVersion = func(path string) (result cm.Option[[2]string]) {
		prefix, pathMajor, ok := module.SplitPathVersion(path)
		if !ok {
			return cm.None[[2]string]()
		}
		return cm.Some([2]string{prefix, pathMajor})
	}
	modulewit.Exports.UnescapePath = func(escaped string) (result cm.Result[string, string, modulewit.Error]) {
		rGo, errGo := module.UnescapePath(escaped)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo)
		return
	}
	modulewit.Exports.UnescapeVersion = func(escaped string) (result cm.Result[string, string, modulewit.Error]) {
		rGo, errGo := module.UnescapeVersion(escaped)
		if errGo != nil {
			result.SetErr(modulewit.Error(errGo.Error()))
			return
		}
		result.SetOK(rGo)
		return
	}
	modulewit.Exports.VersionError = func(v modulewit.Version, err modulewit.Error) (result modulewit.Error) {
		vGo := module.Version{Path: v.Path, Version: v.Version}
		errGo := errors.New(string(err))
		rGo := module.VersionError(vGo, errGo)
		return modulewit.Error(rGo.Error())
	}
	modulewit.Exports.ZeroPseudoVersion = func(major string) (result string) {
		return module.ZeroPseudoVersion(major)
	}
	invalidPathError := resourceMap[*module.InvalidPathError]{}
	modulewit.Exports.InvalidPathError.Constructor = func() (result modulewit.InvalidPathError) {
		selfGo := &module.InvalidPathError{}
		return modulewit.InvalidPathError(invalidPathError.Add(selfGo))
	}
	modulewit.Exports.InvalidPathError.GetKind = func(self cm.Rep) (result string) {
		selfGo := invalidPathError.Get(self)
		return selfGo.Kind
	}
	modulewit.Exports.InvalidPathError.SetKind = func(self cm.Rep, v string) {
		selfGo := invalidPathError.Get(self)
		selfGo.Kind = v
	}
	modulewit.Exports.InvalidPathError.GetPath = func(self cm.Rep) (result string) {
		selfGo := invalidPathError.Get(self)
		return selfGo.Path
	}
	modulewit.Exports.InvalidPathError.SetPath = func(self cm.Rep, v string) {
		selfGo := invalidPathError.Get(self)
		selfGo.Path = v
	}
	modulewit.Exports.InvalidPathError.GetErr = func(self cm.Rep) (result cm.Option[modulewit.Error]) {
		selfGo := invalidPathError.Get(self)
		if selfGo.Err != nil {
			return cm.Some(modulewit.Error(selfGo.Err.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	modulewit.Exports.InvalidPathError.Error = func(self cm.Rep) (result string) {
		selfGo := invalidPathError.Get(self)
		return selfGo.Error()
	}
	modulewit.Exports.InvalidPathError.Unwrap = func(self cm.Rep) (result cm.Option[modulewit.Error]) {
		selfGo := invalidPathError.Get(self)
		rGo := selfGo.Unwrap()
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		} else {
			return cm.None[modulewit.Error]()
		}
	}
	invalidVersionError := resourceMap[*module.InvalidVersionError]{}
	modulewit.Exports.InvalidVersionError.ExportConstructor = func() (result modulewit.InvalidVersionError) {
		selfGo := &module.InvalidVersionError{}
		return modulewit.InvalidVersionError(invalidVersionError.Add(selfGo))
	}
	modulewit.Exports.InvalidVersionError.GetVersion = func(self cm.Rep) (result string) {
		selfGo := invalidVersionError.Get(self)
		return selfGo.Version
	}
	modulewit.Exports.InvalidVersionError.SetVersion = func(self cm.Rep, v string) {
		selfGo := invalidVersionError.Get(self)
		selfGo.Version = v
	}
	modulewit.Exports.InvalidVersionError.GetPseudo = func(self cm.Rep) (result bool) {
		selfGo := invalidVersionError.Get(self)
		return selfGo.Pseudo
	}
	modulewit.Exports.InvalidVersionError.SetPseudo = func(self cm.Rep, v bool) {
		selfGo := invalidVersionError.Get(self)
		selfGo.Pseudo = v
	}
	modulewit.Exports.InvalidVersionError.GetErr = func(self cm.Rep) (result cm.Option[modulewit.Error]) {
		selfGo := invalidVersionError.Get(self)
		if selfGo.Err != nil {
			return cm.Some(modulewit.Error(selfGo.Err.Error()))
		}
		return cm.None[modulewit.Error]()
	}
	modulewit.Exports.InvalidVersionError.Error = func(self cm.Rep) (result string) {
		selfGo := invalidVersionError.Get(self)
		return selfGo.Error()
	}
	modulewit.Exports.InvalidVersionError.Unwrap = func(self cm.Rep) (result cm.Option[modulewit.Error]) {
		selfGo := invalidVersionError.Get(self)
		rGo := selfGo.Unwrap()
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		}
		return cm.None[modulewit.Error]()
	}
	moduleError := resourceMap[*module.ModuleError]{}
	modulewit.Exports.ModuleError.ExportConstructor_ = func() (result modulewit.ModuleError) {
		selfGo := &module.ModuleError{}
		return modulewit.ModuleError(moduleError.Add(selfGo))
	}
	modulewit.Exports.ModuleError.GetPath = func(self cm.Rep) (result string) {
		selfGo := moduleError.Get(self)
		return selfGo.Path
	}
	modulewit.Exports.ModuleError.SetPath = func(self cm.Rep, v string) {
		selfGo := moduleError.Get(self)
		selfGo.Path = v
	}
	modulewit.Exports.ModuleError.GetVersion = func(self cm.Rep) (result string) {
		selfGo := moduleError.Get(self)
		return selfGo.Version
	}
	modulewit.Exports.ModuleError.SetVersion = func(self cm.Rep, v string) {
		selfGo := moduleError.Get(self)
		selfGo.Version = v
	}
	modulewit.Exports.ModuleError.GetErr = func(self cm.Rep) (result cm.Option[modulewit.Error]) {
		selfGo := moduleError.Get(self)
		if selfGo.Err != nil {
			return cm.Some(modulewit.Error(selfGo.Err.Error()))
		}
		return cm.None[modulewit.Error]()
	}
	modulewit.Exports.ModuleError.Error = func(self cm.Rep) (result string) {
		selfGo := moduleError.Get(self)
		return selfGo.Error()
	}
	modulewit.Exports.ModuleError.Unwrap = func(self cm.Rep) (result cm.Option[modulewit.Error]) {
		selfGo := moduleError.Get(self)
		rGo := selfGo.Unwrap()
		if rGo != nil {
			return cm.Some(modulewit.Error(rGo.Error()))
		}
		return cm.None[modulewit.Error]()
	}
	modulewit.Exports.VersionString = func(self modulewit.Version) (result string) {
		selfGo := module.Version{Path: self.Path, Version: self.Version}
		return selfGo.String()
	}
}
