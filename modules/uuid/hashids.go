// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid

import (
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/speps/go-hashids/v2"
)

const (
	minLength int = 10
)

type HashIds struct {
	tools *tool.M
	hid *hashids.HashID
}

func newHashIds() *HashIds {
	return &HashIds {
		tools: tool.New(),
	}
}

func (m *HashIds) Generate(d ... interface{}) (interface{}, error) {
	hd := hashids.NewData()

	if len(d) > 0 && d[0] != "" {
		hd.Salt = d[0].(string)
	} else {
		hd.Salt = m.tools.RandomStr(10)
	}

	if len(d) > 1 && d[1].(int) > 0 {
		hd.MinLength = d[1].(int)
	} else {
		hd.MinLength = minLength
	}

	var err error
	m.hid, err = hashids.NewWithData(hd)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (m *HashIds) EncodeInt64(nums []int64) (string, error) {
	return m.hid.EncodeInt64(nums)
}

func (m *HashIds) DecodeInt64(h string) ([]int64, error) {
	return m.hid.DecodeInt64WithError(h)
}

func (m *HashIds) EncodeHex(d string) (string, error) {
	return m.hid.EncodeHex(d)
}

func (m *HashIds) DecodeHex(d string) (string, error) {
	return m.hid.DecodeHex(d)
}