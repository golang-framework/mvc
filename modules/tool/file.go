// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"os"
)

type File struct {}

func (f *File) IsExists(pathname string) (bool, error) {
	var (
		ok bool = false
		err error
	)

	_, err = os.Stat(pathname)
	if err == nil {
		ok = true
		return ok, err
	}

	if os.IsNotExist(err) {
		ok = false
	} else {
		ok = true
	}

	return ok, err
}

func (f *File) Open(pathname string) (*os.File, error) {
	return os.Open(pathname)
}

func (f *File) Create(pathname string) (*os.File, error) {
	return os.Create(pathname)
}

func (f *File) Mkdir(dir string) error {
	if ok, _ := f.IsExists(dir); ok {
		return nil
	}

	return os.Mkdir(dir, os.ModePerm)
}

func (f *File) MkdirAll(dir string) error {
	if ok, _ := f.IsExists(dir); ok {
		return nil
	}

	return os.MkdirAll(dir, os.ModePerm)
}
