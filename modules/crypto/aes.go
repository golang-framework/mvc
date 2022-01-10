// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

type Aes struct {
	key []byte
}

func newAes() *Aes {
	return &Aes {

	}
}

func (m *Aes) Engine(d ... interface{}) (interface{}, error) {
	if count := len(d); count <= 0 {
		return nil, err.E(storage.KeyM33008)
	}

	if val := cast.ToString(d[0]); val == "" {
		return nil, err.E(storage.KeyM33008)
	}

	comm := newCommon()
	h, errCryptoCommonEngine := comm.Engine(storage.Md5, d[0])
	if errCryptoCommonEngine != nil {
		return nil, errCryptoCommonEngine
	}

	m.key = []byte(cast.ToString(h)[8:24])
	return m, nil
}

func (m *Aes) Encrypt(d interface{}) (interface{}, error) {
	encode, e := m.encrypt([]byte(cast.ToString(d)))
	if e != nil {
		return nil, e
	}

	return base64.StdEncoding.EncodeToString(encode), nil
}

func (m *Aes) Decrypt(cipherCode interface{}) (interface{}, error) {
	decode, e := base64.StdEncoding.DecodeString(cast.ToString(cipherCode))
	if e != nil {
		return nil, e
	}

	d, e := m.decrypt(decode)
	if e != nil {
		return nil, e
	}

	return string(d), nil
}

func (m *Aes) encrypt(d []byte) ([]byte, error) {
	block, e := aes.NewCipher(m.key)
	if e != nil {
		return nil, e
	}

	blockSize := block.BlockSize()
	d = m.pkcS5Padding(d, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, m.key[:blockSize])
	cryptoDsT := make([]byte, len(d))
	blockMode.CryptBlocks(cryptoDsT, d)

	return cryptoDsT, nil
}

func (m *Aes) decrypt(cryptoDsT []byte) ([]byte, error) {
	block, e := aes.NewCipher(m.key)
	if e != nil {
		return nil, e
	}

	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, m.key[:blockSize])
	d := make([]byte, len(cryptoDsT))
	blockMode.CryptBlocks(d, cryptoDsT)
	d = m.pkcS5UnPadding(d)

	return d, nil
}

func (m *Aes) pkcS5Padding(cipherTxT []byte, blockSize int) []byte {
	padding := blockSize - len(cipherTxT)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherTxT, padText ...)
}

func (m *Aes) pkcS5UnPadding(d []byte) []byte {
	aesLength := len(d)
	unpadding := int(d[aesLength-1])
	return d[:(aesLength-unpadding)]
}


