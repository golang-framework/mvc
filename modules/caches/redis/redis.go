// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"github.com/go-redis/redis/v8"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
)

var (
	instance []*redis.Client
)

type M struct {

}

func (m *M) Engine() {
	singleton := newSingleton()
	_ = property.Instance.Usk("redis", &singleton.conns)

	singleton.initialized()
}

func R(d int) (*redis.Client, error) {
	if i := tool.New().ContainSliceIndex(instance, d); i == -1 {
		return nil, err.E(storage.KeyM33005)
	}

	return instance[d], nil
}






