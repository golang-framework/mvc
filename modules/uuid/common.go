// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid

import (
	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"strings"
	"time"
)

type common struct {
	tools *tool.M
}

func newCommon() *common {
	return &common {
		tools: tool.New(),
	}
}

func (m *common) Generate(d ... interface{}) (interface{}, error) {
	var strServiceLabel string
	if len(d) > 0 {
		strServiceLabel = cast.ToString(d[0])
	} else {
		strServiceLabel = m.tools.RandomStr(10)
	}

	var strRandomNumber string
	if len(d) == 2 {
		strRandomNumber = m.tools.RandomStr(d[1].(int))
	} else {
		strRandomNumber = m.tools.RandomStr(10)
	}

	strTimeUnixNano := cast.ToString(time.Now().UnixNano())

	var strCommonUUID string = strings.Join(
		[]string { strServiceLabel,strRandomNumber,strTimeUnixNano }, storage.FwSeparate,
	)

	cry := crypto.New()
	cry.Mode = storage.Common
	cry.D = []interface{}{storage.Md5, strCommonUUID}

	return cry.Engine()
}