// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/modules/exception"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

var (
	Instance *Container = new(Container)
)

type (
	R interface {
		Tag() string
		Put(r *gin.Engine, to map[*Key][]gin.HandlerFunc)
	}

	Src []R
	Arr map[string]map[*Key][]gin.HandlerFunc

	Container struct {
		Src *Src
		Arr *Arr
	}

	Key struct {
		Srv string
		Ctl string
		Act string
		Mod string
		Rel string
	}
)

func (container *Container) Engine(r *gin.Engine) {
	var exp *exception.M = exception.New()

	// todo: no route then redirect
	r.NoRoute(exp.NoRoute)

	// todo: no method then redirect
	r.NoMethod(exp.NoMethod)

	if len(*container.Src) != 0 {
		for _, to := range *container.Src {
			if _, ok := (*container.Arr)[to.Tag()]; ok == false {
				continue
			}

			if len((*container.Arr)[to.Tag()]) == 0 {
				continue
			}

			to.Put(r, (*container.Arr)[to.Tag()])
		}
	}
}

func To(ctx *gin.RouterGroup, to map[*Key][]gin.HandlerFunc) {
	for x, ctrl := range to {
		add := crypto.New()

		add.Mode = storage.Common
		add.D = []interface{}{storage.Md5, x.Ctl + x.Srv + x.Act}

		key, err := add.Engine()
		if err != nil {
			panic(err)
		}

		(*routeMap)[cast.ToString(key)] = x.Rel

		switch x.Mod {
		case Any:
			ctx.Any(x.Rel, ctrl ...)
			continue

		case Get:
			ctx.GET (x.Rel, ctrl ...)
			continue

		case Put:
			ctx.PUT(x.Rel, ctrl ...)
			continue

		case Post:
			ctx.POST(x.Rel, ctrl ...)
			continue

		case Delete:
			ctx.DELETE(x.Rel, ctrl ...)
			continue

		case Head:
			ctx.HEAD(x.Rel, ctrl ...)
			continue

		case Patch:
			ctx.PATCH(x.Rel, ctrl ...)
			continue

		case Options:
			ctx.OPTIONS(x.Rel, ctrl ...)
			continue

		default:
			continue
		}
	}
}
