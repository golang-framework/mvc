// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid

import "github.com/golang-framework/mvc/storage"

type (
	uuid interface {
		Generate(d ... interface{}) (interface{}, error)
	}

	M struct {
		Mode interface{}
		D []interface{}
	}
)

var (
	_ uuid = &common{}
	_ uuid = &snowflake{}
	_ uuid = &HashIds{}
)

func New() *M {
	return &M {

	}
}

func (m *M) Generate() (interface{}, error) {
	var res uuid

	switch m.Mode {
	case storage.HashIds:

		res = newHashIds()
		break

	case storage.SnowFlake:

		res = newSnowFlake()
		break

	default:

		res = newCommon()
		break
	}

	return res.Generate(m.D ...)
}
