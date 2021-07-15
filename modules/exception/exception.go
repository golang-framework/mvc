// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package exception

import (
	"github.com/gin-gonic/gin"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/storage"
)

type M struct {

}

func New() *M {
	return &M {}
}

func (m *M) NoRoute(ctx *gin.Context) {
	ctx.AbortWithError(storage.StatusNotFound, err.E(storage.Incorrect))
	return
}

func (m *M) NoMethod(ctx *gin.Context) {
	ctx.AbortWithError(storage.StatusNotFound, err.E(storage.Incorrect))
	return
}



