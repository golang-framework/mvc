// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	err "github.com/golang-framework/mvc/modules/error"
	mJwT "github.com/golang-framework/mvc/modules/jwt"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/storage"
	"strings"
	"time"
)

func NewJwT(name string) *Component {
	key := strings.Join([]string{"JwT", name}, ".")
	arr := property.Instance.Get(key, nil).(map[string]interface{})

	componentJwT := New()

	if arr["sig"] != "" {
		componentJwT.Signature = arr["sig"].(string)
	}

	if arr["typ"] != "" {
		componentJwT.Typ = arr["typ"]
	}

	if arr["alg"] != "" {
		componentJwT.Alg = arr["alg"]
	}

	if arr["iss"] != "" {
		componentJwT.Iss = arr["iss"]
	}

	if arr["sub"] != "" {
		componentJwT.Sub = arr["sub"]
	}

	if arr["aud"] != "" {
		componentJwT.Aud = arr["aud"]
	}

	return componentJwT
}

type Component struct {
	Signature string
	Typ interface{} 	// 声明类型
	Alg interface{} 	// 声明加密算法
	Iss interface{} 	// 签发者
	Sub interface{} 	// 主题
	Aud interface{} 	// 接受者
	Iat time.Time 		// 生成签名时间
	Nbf time.Time 		// 生效时间(定义在什么时间之前, JWT不可用, 需要晚于签发时间)
	Jti interface{} 	// 编号(唯一身份标识, 识别一次行token, 避免重复攻击)
	Inf interface{} 	// 自定义内容
	Exp time.Duration 	// 多少时间过期（时,分,秒）

	jwt *mJwT.M
}

func New() *Component {
	return &Component {
		Typ: storage.JHeadersTyp,
		Iat: time.Now(),
		Exp: time.Minute,

		jwt: mJwT.New(),
	}
}

func (c *Component) Produce() (interface{}, error) {
	if errJwTLoads := c.loads(); errJwTLoads != nil {
		return nil, errJwTLoads
	}

	return c.jwt.Produce()
}

func (c *Component) Parse(d string) (*mJwT.Headers, *mJwT.Payload, error) {
	if errJwTLoads := c.loads(); errJwTLoads != nil {
		return nil, nil, errJwTLoads
	}

	return c.jwt.Parse(d)
}

func (c *Component) Verify(d string) (int8, error) {
	if errJwTLoads := c.loads(); errJwTLoads != nil {
		return -1, errJwTLoads
	}

	return c.jwt.Verify(d)
}

func (c *Component) loads() error {
	if errJwTCheck := c.chkJwT(); errJwTCheck != nil {
		return errJwTCheck
	}

	c.jwt.SetSignature(c.Signature)
	c.jwt.SetHeadersTyp(c.Typ)
	c.jwt.SetPayloadIss(c.Iss)
	c.jwt.SetPayloadAud(c.Aud)
	c.jwt.SetPayloadSub(c.Sub)
	c.jwt.SetPayloadIat(c.Iat)
	c.jwt.SetPayloadExp(c.Exp)

	if c.Alg != nil {
		c.jwt.SetHeadersAlg(c.Alg)
	}

	if c.Nbf.IsZero() == false {
		c.jwt.SetPayloadNbf(c.Nbf)
	}

	if c.Jti != nil {
		c.jwt.SetPayloadJti(c.Jti)
	}

	if c.Inf != nil {
		c.jwt.SetPayloadInf(c.Inf)
	}

	return nil
}

func (c *Component) chkJwT() error {
	if c.Signature == "" {
		return err.E(storage.KeyM32005)
	}

	if c.Iss == nil || c.Aud == nil || c.Sub == nil {
		return err.E(storage.KeyM32006)
	}

	return nil
}


