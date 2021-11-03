// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	hmacCrypto "crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"hash"
)

type hmac struct {

}

func newHmac() *hmac {
	return &hmac {}
}

func (m *hmac) Engine(d ... interface{}) (interface{}, error) {
	if count := len(d); count <= 2 {
		return nil, err.E(storage.KeyM33009)
	}

	salt, keys := cast.ToString(d[1]), cast.ToString(d[2])
	if salt == "" || keys == "" {
		return nil, err.E(storage.KeyM33010)
	}

	var res hash.Hash

	switch d[0] {
	case storage.Md5:
		res = hmacCrypto.New(md5.New, []byte(salt))
		break

	case storage.Sha1:
		res = hmacCrypto.New(sha1.New, []byte(salt))
		break

	case storage.Sha256:
		res = hmacCrypto.New(sha256.New, []byte(salt))
		break

	default:
		return nil, err.E(storage.KeyM33011)
		break
	}

	return m.do(res, []byte(keys))
}

func (m *hmac) do(res hash.Hash, p []byte) (interface{}, error) {
	res.Write(p)
	return hex.EncodeToString(res.Sum([]byte(""))), nil
}
