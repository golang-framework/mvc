// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/modules/exception"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

var Instance = new(Container)

type (
	arr map[string]map[*key]*Ahc
	key struct {
		srv string
		ctl string
		act string
		rel string
		mod string
	}

	Ahc gin.HandlersChain

	M map[string]*H
	H struct {
		Middleware *Ahc
		Adapter []*Ahc
	}

	Container struct {
		arr *arr
		M *M
		E *Ahc
	}

	Exp struct {
		NoR gin.HandlerFunc
		NoM gin.HandlerFunc
	}
)

func (container *Container) Generate() {
	container.arr = &arr{}
	rArr := property.Instance.Get("route.Arr", map[string]interface{}{}).(map[string]interface{})
	if len(rArr) == 0 {
		panic("Unknown Error! - need add the explanation for error")
	}

	for namespace, val := range rArr {
		if val == nil || len(val.([]interface{})) == 0 {
			continue
		}

		for i, arr := range val.([]interface{}) {
			if arr == nil {
				continue
			}

			if _, ok := arr.(map[interface{}]interface{})["ctl"].(string); ok == false {
				continue
			}

			if _, ok := arr.(map[interface{}]interface{})["act"].(string); ok == false {
				continue
			}

			if _, ok := arr.(map[interface{}]interface{})["rel"].(string); ok == false {
				continue
			}

			if _, ok := arr.(map[interface{}]interface{})["mod"].(string); ok == false {
				continue
			}

			if _, ok := (*container.arr)[namespace]; ok == false {
				(*container.arr)[namespace] = make(map[*key]*Ahc)
			}

			(*container.arr)[namespace][&key {
				srv: namespace,
				ctl: arr.(map[interface{}]interface{})["ctl"].(string),
				act: arr.(map[interface{}]interface{})["act"].(string),
				rel: arr.(map[interface{}]interface{})["rel"].(string),
				mod: arr.(map[interface{}]interface{})["mod"].(string),
			}] = (*container.M)[namespace].Adapter[i]
		}
	}
}

func (container *Container) Engine(r *gin.Engine) {
	exp := exception.New()
	noRouter := exp.NoRoute
	noMethod := exp.NoMethod

	if container.E != nil {
		tools := tool.New()
		if tools.ContainSliceIndex(*container.E, 0) == 1 {
			noRouter = (*container.E)[0]
		}

		if tools.ContainSliceIndex(*container.E, 1) == 1 {
			noMethod = (*container.E)[1]
		}
	}

	// todo: no route then redirect
	r.NoRoute(noRouter)

	// todo: no method then redirect
	r.NoMethod(noMethod)

	rTag := property.Instance.Get("route.Tag", nil)
	if rTag == nil {
		panic("Unknown Error! - need add the explanation for error")
	}

	rRel := property.Instance.Get("route.Rel", nil)
	if rRel == nil {
		panic("Unknown Error! - need add the explanation for error")
	}

	for _, namespace := range rTag.(map[string]interface{}) {
		rRelativePath, okRelativePath := rRel.(map[string]interface{})[namespace.(string)]
		if okRelativePath == false {
			continue
		}

		if _, ok := (*container.arr)[namespace.(string)]; ok == false {
			continue
		}

		if len((*container.arr)[namespace.(string)]) == 0 {
			continue
		}

		if _, ok := (*container.M)[namespace.(string)]; ok == false {
			continue
		}

		hc := &Ahc{}
		if (*container.M)[namespace.(string)].Middleware != nil {
			hc = (*container.M)[namespace.(string)].Middleware
		}

		container.groups(rRelativePath.(string), r, (*container.arr)[namespace.(string)], hc)
	}
}

func (container *Container) groups(relativePath string, r *gin.Engine, to map[*key]*Ahc, hc *Ahc) {
	container.to(r.Group(relativePath, *hc ...), to)
}

func (container *Container) to(ctx *gin.RouterGroup, to map[*key]*Ahc) {
	for x, ctrl := range to {
		add := crypto.New()

		add.Mode = storage.Common
		add.D = []interface{}{storage.Md5, x.ctl + x.srv + x.act}

		k, err := add.Engine()
		if err != nil {
			panic(err)
		}

		(*routeMap)[cast.ToString(k)] = x.rel

		switch x.mod {
		case Any:
			ctx.Any(x.rel, *ctrl ...)
			continue

		case Get:
			ctx.GET(x.rel, *ctrl ...)
			continue

		case Put:
			ctx.PUT(x.rel, *ctrl ...)
			continue

		case Post:
			ctx.POST(x.rel, *ctrl ...)
			continue

		case Head:
			ctx.HEAD(x.rel, *ctrl ...)
			continue

		case Patch:
			ctx.PATCH(x.rel, *ctrl ...)
			continue

		case Delete:
			ctx.DELETE(x.rel, *ctrl ...)
			continue

		case Options:
			ctx.OPTIONS(x.rel, *ctrl ...)
			continue

		default:
			continue
		}
	}
}
