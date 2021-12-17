// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid

import (
	"github.com/golang-framework/mvc/modules/tool"
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


	_ = m.tools.RandomStr(d[0].(int))
	_ = time.Now().UnixNano()

	return nil, nil
}