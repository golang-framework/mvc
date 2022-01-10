// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

const (
	Fw 					= "golang-mvc framework"
	FwVersion		 	= "v1.0.0"
	FwTimeLocation	 	= "Asia/Shanghai"
	FwSeparate 			= "__::__"
)

type (
	// default array construct
	Y map[string]interface{}

	// default error array construct
	E map[string]Y

	// default response template
	Tpl struct {
		Status int
		Msg interface{}
		Res *Y
	}

	TplCookie struct {
		Name string
		Value string
		MaxAge int
		Path string
		Domain string
		Secure bool
		HttpOnly bool
	}
)

func FwTpl(e error) *Tpl {
	return &Tpl {
		Status: StatusOK,
		Msg: e.Error(),
		Res: &Y{},
	}
}
