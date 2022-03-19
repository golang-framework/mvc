// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

type M struct {
}

func New() *M {
	return &M{}
}

func (_ *M) Abort(ctx *gin.Context, e error, d ...interface{}) {
	res := storage.FwTpl(e)

	if len(d) >= 1 {
		res.Status = cast.ToInt(d[0])
	} else {
		res.Status = storage.StatusUnknown
	}

	if len(d) == 2 {
		if reflect.TypeOf(d[1]).Kind() == reflect.Ptr && reflect.TypeOf(d[1]).String() == "*storage.Y" {
			res.Res = d[1].(*storage.Y)
		}
	}

	ctx.AbortWithStatusJSON(storage.StatusOK, res)
	return
}
