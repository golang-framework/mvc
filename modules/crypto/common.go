// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"hash"
)

type Common struct {

}

func newCommon() *Common {
	return &Common {

	}
}

func (m *Common) Engine(d ... interface{}) (interface{}, error) {
	if len(d) <= 1 || cast.ToString(d[1]) == "" {
		return nil, err.E(storage.KeyM33001)
	}

	var res hash.Hash

	switch d[0] {
	case storage.Md5:
		res = md5.New()
		return m.do(res, []byte(cast.ToString(d[1])))

	case storage.Sha1:
		res = sha1.New()
		return m.do(res, []byte(cast.ToString(d[1])))

	case storage.Sha256:
		res = sha256.New()
		return m.do(res, []byte(cast.ToString(d[1])))

	default:
		return nil, err.E(storage.KeyM33002)
	}
}

func (m *Common) do(res hash.Hash, p []byte) (interface{}, error) {
	res.Write(p)
	return hex.EncodeToString(res.Sum([]byte(""))), nil
}