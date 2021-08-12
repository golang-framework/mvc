// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package db

import (
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"xorm.io/xorm"
)

const (
	mysql      	string = "mysql"
	postgreSQL 	string = "postgres"
)

var (
	instance []*xorm.Engine
	instanceGroup []*xorm.EngineGroup
)

type (
	M struct {

	}

	requirement struct {
		Driver 			string
		MaxOpen			int
		MaxIdle			int
		MaxLifeTime		int
		ShowedSQL 		bool
		CachedSQL 		bool
		Expire			int
		MaxElementSize	int
		TimeLoaction	string
	}

	conns struct {
		Require 		int
		Repo 			string
		Host 			string
		Port 			int
		Username 		string
		Password 		string
	}
)

func (m *M) Engine() {
	singleton := newSingleton()
	_ = property.Instance.Usk("db.requirement", &singleton.requirement)
	_ = property.Instance.Usk("db.conns", &singleton.conns)

	singleton.initialized()
}

func Engine(d int) (*xorm.Engine, error) {
	if tool.New().ContainSliceIndex(instance, d) == -1 {
		return nil, err.E(storage.KeyM33007)
	}

	return instance[d], nil
}


