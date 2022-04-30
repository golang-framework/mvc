// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"database/sql"
	"reflect"
	"strings"

	"github.com/golang-framework/mvc/modules/db"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
	"xorm.io/xorm"
)

type Model struct {
	conn *xorm.Engine
}

func New(d int) *Model {
	engine := &Model{}
	engine.conn, _ = db.Engine(d)

	return engine
}

func (mod *Model) Conn() *xorm.Engine {
	return mod.conn
}

func (mod *Model) Exec(sql ...interface{}) (sql.Result, error) {
	return mod.conn.Exec(sql...)
}

func (mod *Model) Query(sql ...interface{}) ([]map[string][]byte, error) {
	return mod.conn.Query(sql...)
}

func (mod *Model) Insert(d ...interface{}) (int64, error) {
	db := mod.conn.NewSession()
	defer func() { db.Close() }()

	return db.Insert(d...)
}

func (mod *Model) Update(conditions *storage.Conditions, d interface{}, bean ...interface{}) (int64, error) {
	db := mod.conn.NewSession()
	defer func() { db.Close() }()

	return db.Where(conditions.Query, conditions.QueryArgs...).Update(d, bean...)
}

func (mod *Model) Delete(conditions *storage.Conditions, d ...interface{}) (int64, error) {
	db := mod.conn.NewSession()
	defer func() { db.Close() }()

	return db.Where(conditions.Query, conditions.QueryArgs...).Delete(d...)
}

func (mod *Model) Count(conditions *storage.Conditions, d ...interface{}) (int64, error) {
	db := mod.conn.NewSession()
	defer func() { db.Close() }()

	// add where conditions
	if conditions.Query != nil && conditions.QueryArgs != nil {
		db = db.Where(conditions.Query, conditions.QueryArgs...)
	}

	return db.Count(d...)
}

func (mod *Model) Select(conditions *storage.Conditions, d interface{}) (int8, error) {
	db := mod.conn.NewSession()
	defer func() { db.Close() }()

	if conditions.Table != "" && reflect.TypeOf(conditions.Field).Kind() == reflect.Slice {
		if ok, _ := db.IsTableExist(conditions.Table); ok == false {
			return -1, err.E(storage.KeyM32002)
		}

		db = db.Table(conditions.Table).Select(strings.Join(conditions.Field, ","))
	} else {
		db = db.Cols(conditions.Columns...)
	}

	// join table
	if conditions.Joins != nil {
		for _, join := range conditions.Joins {
			db = db.Join(join.JoinOperator, join.TableName, join.Condition)
		}
	}

	// add where conditions
	if conditions.Query != nil && conditions.QueryArgs != nil {
		db = db.Where(conditions.Query, conditions.QueryArgs...)
	}

	// order by
	if len(conditions.OrderArgs) > 0 {
		switch conditions.OrderType {
		case storage.ByAsc:
			db = db.Asc(conditions.OrderArgs...)

		case storage.ByEsc:
			db = db.Desc(conditions.OrderArgs...)

		default:
			return -1, err.E(storage.KeyM32003)
		}
	}

	switch conditions.Types {
	case storage.SelectOne:
		_, errXormEngine := db.Get(d)
		if errXormEngine != nil {
			return -1, errXormEngine
		}

		return 1, nil

	case storage.SelectAll:

		// limit & start
		if conditions.Limit > 0 && conditions.Start >= 0 {
			db.Limit(conditions.Limit, conditions.Start)
		}

		return 1, db.Find(d)

	default:
		return -1, err.E(storage.KeyM32004)
	}
}
