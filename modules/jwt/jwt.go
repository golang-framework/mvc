// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	"encoding/base64"
	"strings"
	"time"

	"github.com/golang-framework/mvc/modules/crypto"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

type (
	M struct {
		cry   *crypto.M
		tools *tool.M

		sign    string
		headers *Headers
		payload *Payload
	}

	Headers struct {
		Typ interface{} // 声明类型
		Alg interface{} // 声明加密算法
	}

	Payload struct {
		Iss interface{}   // 签发者
		Sub interface{}   // 主题
		Aud interface{}   // 接受者
		Iat time.Time     // 生成签名时间
		Nbf time.Time     // 生效时间(定义在什么时间之前, JWT不可用, 需要晚于签发时间)
		Jti interface{}   // 编号(唯一身份标识, 识别一次行token, 避免重复攻击)
		Inf interface{}   // 自定义内容
		Exp time.Duration // 多少时间过期（时,分,秒）
	}
)

func New() *M {
	return &M{
		cry:   crypto.New(),
		tools: tool.New(),

		headers: &Headers{
			Typ: storage.JHeadersTyp,
			Alg: storage.Sha256,
		},
		payload: &Payload{
			Iss: strings.Join([]string{storage.Fw, storage.FwVersion}, ","),
			Aud: storage.JPayloadAud,
			Iat: time.Now(),
			Exp: time.Minute,
			Inf: nil,
		},
	}
}

func (m *M) Produce() (interface{}, error) {
	if errChkSignature := m.chkSignature(); errChkSignature != nil {
		return nil, errChkSignature
	}

	return m.generateCT(m.headers, m.payload)
}

func (m *M) Parse(c string) (*Headers, *Payload, error) {
	if errChkSignature := m.chkSignature(); errChkSignature != nil {
		return nil, nil, errChkSignature
	}

	return m.analysisCT(c)
}

func (m *M) Refresh(c string) (interface{}, error) {
	_, errJwTVerify := m.verify(c)
	if errJwTVerify != nil {
		return nil, errJwTVerify
	}

	_, payload, errJwTParse := m.Parse(c)
	if errJwTParse != nil {
		return nil, errJwTParse
	}

	if payload.Iat.Add(payload.Exp).Before(time.Now()) {
		return nil, err.E(storage.KeyM33019)
	}

	d, errJwTProduce := m.Produce()
	if errJwTProduce != nil {
		return nil, errJwTProduce
	}

	return d, nil
}

func (m *M) verify(c string) (int8, error) {
	headers, payload, errJwTParse := m.Parse(c)
	if errJwTParse != nil {
		return -1, errJwTParse
	}

	ciphers, errJwTGenerateCT := m.generateCT(headers, payload)
	if errJwTGenerateCT != nil {
		return -1, errJwTGenerateCT
	}

	if c != ciphers {
		return -1, err.E(storage.KeyM33018)
	}

	return 1, nil
}

func (m *M) analysisCT(c string) (*Headers, *Payload, error) {
	var (
		headers = &Headers{}
		payload = &Payload{}
	)

	switch m.headers.Alg {
	case storage.Md5, storage.Sha1, storage.Sha256:

		s := strings.Split(c, ".")
		if len(s) != 3 {
			return nil, nil, err.E(storage.KeyM33013)
		}

		m.cry.Mode = storage.Hmac
		m.cry.D = []interface{}{
			m.headers.Alg,
			m.sign,
			strings.Join([]string{s[0], s[1]}, storage.FwSeparate),
		}

		hmacEncode, errJwTCode := m.cry.Engine()
		if errJwTCode != nil {
			return nil, nil, errJwTCode
		}

		if s[2] != base64.StdEncoding.EncodeToString([]byte(cast.ToString(hmacEncode))) {
			return nil, nil, err.E(storage.KeyM33016)
		}

		if errJwTDecodeToHP := m.decodeToHP(s[0], s[1], headers, payload); errJwTDecodeToHP != nil {
			return nil, nil, errJwTDecodeToHP
		}

		break

	case storage.Aes:

		m.cry.Mode = storage.Aes
		m.cry.D = []interface{}{m.sign}

		aesEncode, errJwTEncode := m.cry.Engine()
		if errJwTEncode != nil {
			return nil, nil, errJwTEncode
		}

		aesDecode, errJwTDecode := aesEncode.(*crypto.Aes).Decrypt(c)
		if errJwTDecode != nil {
			return nil, nil, errJwTDecode
		}

		s := strings.Split(cast.ToString(aesDecode), storage.FwSeparate)
		if errJwTDecodeToHP := m.decodeToHP(s[0], s[1], headers, payload); errJwTDecodeToHP != nil {
			return nil, nil, errJwTDecodeToHP
		}

		break

	case storage.Rsa:

		break

	default:

		break
	}

	return headers, payload, nil
}

func (m *M) generateCT(headers *Headers, payload *Payload) (interface{}, error) {
	strHeaders, _ := m.tools.ToJson(headers)
	strPayload, _ := m.tools.ToJson(payload)

	bteHeaders := m.tools.Base64ToEncode([]byte(strHeaders))
	btePayload := m.tools.Base64ToEncode([]byte(strPayload))

	ciphers := strings.Join([]string{bteHeaders, btePayload}, storage.FwSeparate)

	var c interface{}

	switch m.headers.Alg {
	case storage.Md5, storage.Sha1, storage.Sha256:

		m.cry.Mode = storage.Hmac
		m.cry.D = []interface{}{m.headers.Alg, m.sign, ciphers}

		hmacEncode, errJwTHmacEngine := m.cry.Engine()
		if errJwTHmacEngine != nil {
			return nil, errJwTHmacEngine
		}

		c = strings.Join([]string{
			bteHeaders, btePayload,
			m.tools.Base64ToEncode([]byte(cast.ToString(hmacEncode))),
		}, ".")

		break

	case storage.Aes:

		m.cry.Mode = storage.Aes
		m.cry.D = []interface{}{m.sign}

		aesEngine, errJwTEngine := m.cry.Engine()
		if errJwTEngine != nil {
			return nil, errJwTEngine
		}

		aesEncode, errJwTEncode := aesEngine.(*crypto.Aes).Encrypt(ciphers)
		if errJwTEncode != nil {
			return nil, errJwTEncode
		}

		c = aesEncode

		break

	case storage.Rsa:

		return nil, err.E(storage.KeyM33012)
		break

	default:

		return nil, err.E(storage.KeyM33012)
		break
	}

	return c, nil
}

func (m *M) decodeToHP(cipherHeaders, cipherPayload string, headers *Headers, payload *Payload) error {
	decodeToHeaders, errDecodeToHeaders := m.tools.Base64ToDecode(cipherHeaders)
	if errDecodeToHeaders != nil {
		return err.E(storage.KeyM33014)
	}

	decodeToPayload, errDecodeToPayload := m.tools.Base64ToDecode(cipherPayload)
	if errDecodeToPayload != nil {
		return err.E(storage.KeyM33015)
	}

	_ = m.tools.ToStruct(decodeToHeaders, &headers)
	_ = m.tools.ToStruct(decodeToPayload, &payload)

	return nil
}

func (m *M) chkSignature() error {
	if strings.Trim(m.sign, " ") == "" {
		return err.E(storage.KeyM33017)
	}

	return nil
}

func (m *M) SetSignature(sign string) *M {
	m.sign = strings.Trim(sign, " ")
	return m
}

func (m *M) SetHeadersTyp(typ interface{}) *M {
	m.headers.Typ = typ
	return m
}

func (m *M) SetHeadersAlg(alg interface{}) *M {
	m.headers.Alg = alg
	return m
}

func (m *M) SetPayloadIss(iss interface{}) *M {
	m.payload.Iss = iss
	return m
}

func (m *M) SetPayloadSub(sub interface{}) *M {
	m.payload.Sub = sub
	return m
}

func (m *M) SetPayloadAud(aud interface{}) *M {
	m.payload.Aud = aud
	return m
}

func (m *M) SetPayloadInf(inf interface{}) *M {
	m.payload.Inf = inf
	return m
}

func (m *M) SetPayloadJti(jti interface{}) *M {
	m.payload.Jti = jti
	return m
}

func (m *M) SetPayloadIat(iat time.Time) *M {
	m.payload.Iat = iat
	return m
}

func (m *M) SetPayloadNbf(nbf time.Time) *M {
	m.payload.Nbf = nbf
	return m
}

func (m *M) SetPayloadExp(exp time.Duration) *M {
	m.payload.Exp = exp
	return m
}
