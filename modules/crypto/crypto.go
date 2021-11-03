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
	_ crypto = &common{}
	_ crypto = &hmac{}
	_ crypto = &Aes{}
)

func New() *M {
	return &M {}
}

func (m *M) Engine() (interface{}, error) {
	var res crypto

	switch m.Mode {
	case storage.Common:
		res = newCommon()
		break

	case storage.Hmac:
		res = newHmac()
		break

	case storage.Aes:
		res = newAes()
		break

	default:
		return nil, err.E(storage.KeyM33003)
		break
	}

	return res.Engine(m.D ...)
}

