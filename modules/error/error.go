// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package error

import (
	"errors"
	"fmt"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"runtime"
)

type M struct {
	EMsg *storage.E
}

func (m *M) Initialized() {
	if m.EMsg != nil {
		storage.SetError(m.EMsg)
	}
}

func E(k string, content ... interface{}) error {
	return Err(storage.ErrPrefix, k, content ...)
}

func Err(pfx, k string, content ... interface{}) error {
	return errors.New(storage.GetError(pfx, k, content ...))
}

func Num(pfx, k string, d ... int) string {
	start, limit := 1, 5

	if len(d) > 0 && d[0] > 0 {
		start = d[0]
	}

	if len(d) > 1 && d[1] > 0 {
		limit = d[1]
	}

	sg := tool.New()
	return sg.SourceGrab(storage.GetError(pfx, k), start, limit)
}

func Position() interface{} {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}


