// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"github.com/golang-framework/mvc/storage"
	err "github.com/golang-framework/mvc/modules/error"
)

type (
	crypto interface {
		Engine(d ... interface{}) (interface{}, error)
	}

	M struct {
		Mode interface{}
		D []interface{}
	}
)

var (
	_ crypto = &Common{}
)

func New() *M {
	return &M {}
}

func (m *M) Engine() (interface{}, error) {
	switch m.Mode {
	case storage.Common:
		var res crypto = newCommon()
		return res.Engine(m.D ...)

	default:
		return nil, err.E(storage.KeyM33003)
	}
}

