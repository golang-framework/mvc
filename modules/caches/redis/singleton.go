// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
	"reflect"
	"sync"
)

type (
	singleton struct {
		mx *sync.Mutex
		conns []*params
	}

	params struct {
		Network		string
		Addr 		string
		Username	string
		Password 	string
		DB 			int
		PoolSize	int
	}
)

func newSingleton() *singleton {
	return &singleton {
		mx: &sync.Mutex{},
	}
}

func (m *singleton) initialized() {
	if ok := reflect.ValueOf(m.conns).IsNil(); ok {
		return
	}

	if kind := reflect.TypeOf(m.conns).Kind(); kind != reflect.Slice {
		panic(err.E(storage.KeyM33004))
	}

	instance = make([]*redis.Client, len(m.conns))
	for i, cfg := range m.conns {
		if kind := reflect.TypeOf(cfg).Kind(); kind == reflect.Ptr {
			instance[i] = m.engine(cfg)
		}
	}
}

func (m *singleton) engine(cfg *params) *redis.Client {
	m.mx.Lock()
	defer func() { m.mx.Unlock() }()

	client := redis.NewClient(&redis.Options {
		Network:            cfg.Network,
		Addr:               cfg.Addr,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           cfg.Username,
		Password:           cfg.Password,
		DB:                 cfg.DB,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolSize:           cfg.PoolSize,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	_, errRedisSingletonEngine := client.Ping(context.Background()).Result()
	if errRedisSingletonEngine != nil {
		panic(errRedisSingletonEngine)
	}

	return client
}


