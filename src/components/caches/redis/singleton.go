// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	r "github.com/golang-framework/mvc/modules/caches/redis"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"strings"
	"time"
)

type Component struct {
	sep string
	pfx string
	ctx context.Context
	client *redis.Client
}

func New(d int) *Component {
	client, _ := r.R(d)

	return &Component {
		sep:	"::",
		pfx:	"",
		ctx:    context.Background(),
		client: client,
	}
}

func (c *Component) check() error {
	if c.client == nil {
		return err.E(storage.KeyM32001)
	}

	return nil
}

func (c *Component) addPrefix(key string) string {
	if c.pfx == "" {
		return key
	}

	return strings.Join([]string{c.pfx, key}, c.sep)
}

func (c *Component) SetPrefix(pfx string) {
	c.pfx = pfx
}

func (c *Component) SetPrefixSeparate(sep string) {
	c.sep = sep
}

func (c *Component) Key(key string) string {
	return c.addPrefix(key)
}

func (c *Component) Keys(pattern interface{}) (interface{}, error) {
	if e := c.check(); e != nil {
		return nil, e
	}

	return c.client.Keys(c.ctx, cast.ToString(pattern)).Result()
}

func (c *Component) Set(key, val interface{}, d ... time.Duration) (interface{}, error) {
	if e := c.check(); e != nil {
		return nil, e
	}

	var expire time.Duration = 0
	if len(d) == 1 {
		expire = d[0]
	}

	return c.client.Set(c.ctx, c.addPrefix(cast.ToString(key)), val, expire).Result()
}

func (c *Component) Get(key interface{}) (interface{}, error) {
	if e := c.check(); e != nil {
		return nil, e
	}

	return c.client.Get(c.ctx, c.addPrefix(cast.ToString(key))).Result()
}

func (c *Component) Del(key interface{}) (interface{}, error) {
	if e := c.check(); e != nil {
		return nil, e
	}

	return c.client.Del(c.ctx, c.addPrefix(cast.ToString(key))).Result()
}

func (c *Component) IsExist(key interface{}) (interface{}, error) {
	if e := c.check(); e != nil {
		return nil, e
	}

	return c.client.Exists(c.ctx, c.addPrefix(cast.ToString(key))).Result()
}













