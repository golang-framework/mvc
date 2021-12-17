// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package uuid

import "sync"

type snowflake struct {
	mx sync.Mutex
}

func newSnowFlake() *snowflake {
	return &snowflake {

	}
}

func (m *snowflake) Generate(d ... interface{}) (interface{}, error) {

	return nil, nil
}

