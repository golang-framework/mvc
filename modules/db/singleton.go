// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"reflect"
	"sync"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

type (
	singleton struct {
		mx *sync.Mutex
		requirement []*requirement
		conns []*conns
		tools *tool.M
	}
)

func newSingleton() *singleton {
	return &singleton {
		mx: &sync.Mutex{},
		tools: tool.New(),
	}
}

func (m *singleton) initialized() {
	if reflect.ValueOf(m.conns).IsNil() {
		return
	}

	if reflect.TypeOf(m.conns).Kind() != reflect.Slice {
		panic(err.E(storage.KeyM33006))
	}

	instance = make([]*xorm.Engine, len(m.conns))
	for i, cfg := range m.conns {
		if m.tools.ContainSliceIndex(m.requirement, cfg.Require) == -1 {
			continue
		}

		if m.tools.Contains(m.requirement[cfg.Require].Driver, mysql, postgreSQL) == -1 {
			continue
		}

		if reflect.TypeOf(cfg).Kind() == reflect.Ptr {
			instance[i] = m.engine(m.requirement[cfg.Require], cfg)
		}
	}
}

func (m *singleton) engine(requirement *requirement, conns *conns) *xorm.Engine {
	m.mx.Lock()
	defer func() { m.mx.Unlock() }()

	var dataSourceName = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
		conns.Username,
		conns.Password,
		conns.Host,
		conns.Port,
		conns.Repo,
	)

	adapter, errXormEngine := xorm.NewEngine(requirement.Driver, dataSourceName)
	if errXormEngine != nil {
		panic(errXormEngine)
	}

	adapter.SetMaxOpenConns(requirement.MaxOpen)
	adapter.SetMaxIdleConns(requirement.MaxIdle)

	adapter.ShowSQL(requirement.ShowedSQL)

	area, _ := time.LoadLocation(requirement.TimeLoaction)
	adapter.SetTZDatabase(area)

	if requirement.MaxLifeTime != 0 {
		adapter.SetConnMaxLifetime(cast.ToDuration(requirement.MaxLifeTime) * time.Second)
	} else {
		adapter.SetConnMaxLifetime(cast.ToDuration(1) * time.Minute)
	}

	if requirement.CachedSQL && requirement.Expire != 0 && requirement.MaxElementSize != 0 {
		c := caches.NewLRUCacher2(caches.NewMemoryStore(), cast.ToDuration(requirement.Expire), requirement.MaxElementSize)
		adapter.SetDefaultCacher(c)
	}

	return adapter
}
