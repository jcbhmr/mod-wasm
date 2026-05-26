//go:build wasip2

package main

import (
	"errors"
	"io"
	"os"

	"go.bytecodealliance.org/cm"
	zipwit "go.jcbhmr.com.internal/mod-wasm/internal/jcbhmr/mod/zip"
	"golang.org/x/mod/module"
	"golang.org/x/mod/zip"
)

type pathFile string

func (f pathFile) Path() string {
	return string(f)
}

func (f pathFile) Lstat() (os.FileInfo, error) {
	return os.Lstat(string(f))
}

func (f pathFile) Open() (io.ReadCloser, error) {
	return os.Open(string(f))
}

func init() {
	zipwit.Exports.GetMaxZipFile = func() (result int64) {
		return zip.MaxZipFile
	}
	zipwit.Exports.GetMaxGoMod = func() (result int64) {
		return zip.MaxGoMod
	}
	zipwit.Exports.GetMaxLicense = func() (result int64) {
		return zip.MaxLICENSE
	}
	zipwit.Exports.Create = func(w string, m zipwit.Version, files cm.List[zipwit.File]) (result cm.Result[zipwit.Error, struct{}, zipwit.Error]) {
		wGo, err := os.Open(w)
		if err != nil {
			result.SetErr(zipwit.Error(err.Error()))
			return
		}
		mGo := module.Version{Path: m.Path, Version: m.Version}
		filesGo := make([]zip.File, files.Len())
		for i, f := range files.Slice() {
			filesGo[i] = pathFile(f)
		}
		errGo := zip.Create(wGo, mGo, filesGo)
		if errGo != nil {
			result.SetErr(zipwit.Error(errGo.Error()))
			return
		}
		result.SetOK(struct{}{})
		return
	}
	zipwit.Exports.CreateFromDir = func(w string, m zipwit.Version, dir string) (result cm.Result[zipwit.Error, struct{}, zipwit.Error]) {
		wGo, err := os.Open(w)
		if err != nil {
			result.SetErr(zipwit.Error(err.Error()))
		}
		mGo := module.Version{Path: m.Path, Version: m.Version}
		errGo := zip.CreateFromDir(wGo, mGo, dir)
		if errGo != nil {
			result.SetErr(zipwit.Error(errGo.Error()))
			return
		}
		result.SetOK(struct{}{})
		return
	}
	zipwit.Exports.CreateFromVCS = func(w string, m zipwit.Version, repoRoot, revision, subdir string) (result cm.Result[zipwit.Error, struct{}, zipwit.Error]) {
		wGo, err := os.Open(w)
		if err != nil {
			result.SetErr(zipwit.Error(err.Error()))
		}
		mGo := module.Version{Path: m.Path, Version: m.Version}
		errGo := zip.CreateFromVCS(wGo, mGo, repoRoot, revision, subdir)
		if errGo != nil {
			result.SetErr(zipwit.Error(errGo.Error()))
			return
		}
		result.SetOK(struct{}{})
		return
	}
	zipwit.Exports.CheckDir = func(dir string) (result cm.Tuple[zipwit.CheckedFiles, cm.Option[zipwit.Error]]) {
		cfGo, errGo := zip.CheckDir(dir)
		cfOmitted := make([]zipwit.FileError, len(cfGo.Omitted))
		for i, v := range cfGo.Omitted {
			var err cm.Option[zipwit.Error]
			if v.Err != nil {
				err = cm.Some(zipwit.Error(v.Err.Error()))
			} else {
				err = cm.None[zipwit.Error]()
			}
			cfOmitted[i] = zipwit.FileError{Path: v.Path, Err: err}
		}
		cfInvalid := make([]zipwit.FileError, len(cfGo.Invalid))
		for i, v := range cfGo.Invalid {
			var err cm.Option[zipwit.Error]
			if v.Err != nil {
				err = cm.Some(zipwit.Error(v.Err.Error()))
			} else {
				err = cm.None[zipwit.Error]()
			}
			cfInvalid[i] = zipwit.FileError{Path: v.Path, Err: err}
		}
		var cfSizeError cm.Option[zipwit.Error]
		if cfGo.SizeError != nil {
			cfSizeError = cm.Some(zipwit.Error(cfGo.SizeError.Error()))
		} else {
			cfSizeError = cm.None[zipwit.Error]()
		}
		cf := zipwit.CheckedFiles{
			Valid:     cm.ToList(cfGo.Valid),
			Omitted:   cm.ToList(cfOmitted),
			Invalid:   cm.ToList(cfInvalid),
			SizeError: cfSizeError,
		}
		var f1 cm.Option[zipwit.Error]
		if errGo != nil {
			f1 = cm.Some(zipwit.Error(errGo.Error()))
		} else {
			f1 = cm.None[zipwit.Error]()
		}
		return cm.Tuple[zipwit.CheckedFiles, cm.Option[zipwit.Error]]{F0: cf, F1: f1}
	}
	zipwit.Exports.CheckFiles = func(files cm.List[zipwit.File]) (result cm.Tuple[zipwit.CheckedFiles, cm.Option[zipwit.Error]]) {
		filesGo := make([]zip.File, files.Len())
		for i, f := range files.Slice() {
			filesGo[i] = pathFile(f)
		}
		cfGo, errGo := zip.CheckFiles(filesGo)
		cfOmitted := make([]zipwit.FileError, len(cfGo.Omitted))
		for i, v := range cfGo.Omitted {
			var err cm.Option[zipwit.Error]
			if v.Err != nil {
				err = cm.Some(zipwit.Error(v.Err.Error()))
			} else {
				err = cm.None[zipwit.Error]()
			}
			cfOmitted[i] = zipwit.FileError{Path: v.Path, Err: err}
		}
		cfInvalid := make([]zipwit.FileError, len(cfGo.Invalid))
		for i, v := range cfGo.Invalid {
			var err cm.Option[zipwit.Error]
			if v.Err != nil {
				err = cm.Some(zipwit.Error(v.Err.Error()))
			} else {
				err = cm.None[zipwit.Error]()
			}
			cfInvalid[i] = zipwit.FileError{Path: v.Path, Err: err}
		}
		var cfSizeError cm.Option[zipwit.Error]
		if cfGo.SizeError != nil {
			cfSizeError = cm.Some(zipwit.Error(cfGo.SizeError.Error()))
		} else {
			cfSizeError = cm.None[zipwit.Error]()
		}
		cf := zipwit.CheckedFiles{
			Valid:     cm.ToList(cfGo.Valid),
			Omitted:   cm.ToList(cfOmitted),
			Invalid:   cm.ToList(cfInvalid),
			SizeError: cfSizeError,
		}
		var f1 cm.Option[zipwit.Error]
		if errGo != nil {
			f1 = cm.Some(zipwit.Error(errGo.Error()))
		} else {
			f1 = cm.None[zipwit.Error]()
		}
		return cm.Tuple[zipwit.CheckedFiles, cm.Option[zipwit.Error]]{F0: cf, F1: f1}
	}
	zipwit.Exports.CheckZip = func(m zipwit.Version, zipFile string) (result cm.Tuple[zipwit.CheckedFiles, cm.Option[zipwit.Error]]) {
		mGo := module.Version{Path: m.Path, Version: m.Version}
		cfGo, errGo := zip.CheckZip(mGo, zipFile)
		cfOmitted := make([]zipwit.FileError, len(cfGo.Omitted))
		for i, v := range cfGo.Omitted {
			var err cm.Option[zipwit.Error]
			if v.Err != nil {
				err = cm.Some(zipwit.Error(v.Err.Error()))
			} else {
				err = cm.None[zipwit.Error]()
			}
			cfOmitted[i] = zipwit.FileError{Path: v.Path, Err: err}
		}
		cfInvalid := make([]zipwit.FileError, len(cfGo.Invalid))
		for i, v := range cfGo.Invalid {
			var err cm.Option[zipwit.Error]
			if v.Err != nil {
				err = cm.Some(zipwit.Error(v.Err.Error()))
			} else {
				err = cm.None[zipwit.Error]()
			}
			cfInvalid[i] = zipwit.FileError{Path: v.Path, Err: err}
		}
		var cfSizeError cm.Option[zipwit.Error]
		if cfGo.SizeError != nil {
			cfSizeError = cm.Some(zipwit.Error(cfGo.SizeError.Error()))
		} else {
			cfSizeError = cm.None[zipwit.Error]()
		}
		cf := zipwit.CheckedFiles{
			Valid:     cm.ToList(cfGo.Valid),
			Omitted:   cm.ToList(cfOmitted),
			Invalid:   cm.ToList(cfInvalid),
			SizeError: cfSizeError,
		}
		var f1 cm.Option[zipwit.Error]
		if errGo != nil {
			f1 = cm.Some(zipwit.Error(errGo.Error()))
		} else {
			f1 = cm.None[zipwit.Error]()
		}
		return cm.Tuple[zipwit.CheckedFiles, cm.Option[zipwit.Error]]{F0: cf, F1: f1}
	}
	zipwit.Exports.CheckedFilesErr = func(self zipwit.CheckedFiles) (result cm.Option[zipwit.Error]) {
		omittedGo := make([]zip.FileError, self.Omitted.Len())
		for i, v := range self.Omitted.Slice() {
			var errGo error
			if err := v.Err.Some(); err != nil {
				errGo = errors.New(string(*err))
			}
			omittedGo[i] = zip.FileError{
				Path: v.Path,
				Err:  errGo,
			}
		}
		invalidGo := make([]zip.FileError, self.Invalid.Len())
		for i, v := range self.Omitted.Slice() {
			var errGo error
			if err := v.Err.Some(); err != nil {
				errGo = errors.New(string(*err))
			}
			invalidGo[i] = zip.FileError{
				Path: v.Path,
				Err:  errGo,
			}
		}
		var sizeErrorGo error
		if sizeError := self.SizeError.Some(); sizeError != nil {
			sizeErrorGo = errors.New(string(*sizeError))
		}
		selfGo := zip.CheckedFiles{
			Valid:     self.Valid.Slice(),
			Omitted:   omittedGo,
			Invalid:   invalidGo,
			SizeError: sizeErrorGo,
		}
		errGo := selfGo.Err()
		if errGo != nil {
			return cm.Some(zipwit.Error(errGo.Error()))
		} else {
			return cm.None[zipwit.Error]()
		}
	}
	zipwit.Exports.FileErrorError = func(self zipwit.FileError) (result string) {
		var errGo error
		if err := self.Err.Some(); err != nil {
			errGo = errors.New(string(*err))
		}
		selfGo := zip.FileError{
			Path: self.Path,
			Err:  errGo,
		}
		return selfGo.Error()
	}
	zipwit.Exports.FileErrorUnwrap = func(self zipwit.FileError) (result cm.Option[zipwit.Error]) {
		var errGo error
		if err := self.Err.Some(); err != nil {
			errGo = errors.New(string(*err))
		}
		selfGo := zip.FileError{
			Path: self.Path,
			Err:  errGo,
		}
		rGo := selfGo.Unwrap()
		if rGo != nil {
			return cm.Some(zipwit.Error(rGo.Error()))
		} else {
			return cm.None[zipwit.Error]()
		}
	}
	zipwit.Exports.FileErrorListError = func(self zipwit.FileErrorList) (result string) {
		fileErrorsGo := make([]zip.FileError, self.Len())
		for i, v := range self.Slice() {
			var errGo error
			if err := v.Err.Some(); err != nil {
				errGo = errors.New(string(*err))
			}
			fileErrorsGo[i] = zip.FileError{
				Path: v.Path,
				Err:  errGo,
			}
		}
		selfGo := zip.FileErrorList(fileErrorsGo)
		return selfGo.Error()
	}
	zipwit.Exports.UnrecognizedVCSErrorError = func(self zipwit.UnrecognizedVCSError) (result string) {
		selfGo := zip.UnrecognizedVCSError{RepoRoot: self.RepoRoot}
		return selfGo.Error()
	}
}
