// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

const (
	Fw string = "golang mvc framework"
	FwVersion string = "v1.0.0"
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
)
