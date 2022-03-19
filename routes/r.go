// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/modules/crypto"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/exception"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

var Instance = new(Container)

type (
	arr map[string]map[*key]*AHC
	key struct {
		srv string
		ctl string
		act string
		rel string
		mod string
	}

	AHC gin.HandlersChain

	M map[string]*H
	I []interface{}

	H struct {
		Middleware *AHC
		Adapter    map[*I]*AHC
	}

	Container struct {
		arr    *arr
		toolTP *tool.M

		M *M
		E *AHC
	}

	Exp struct {
		NoR gin.HandlerFunc
		NoM gin.HandlerFunc
	}
)

func (container *Container) Load() *Container {
	container.arr = &arr{}
	container.toolTP = tool.New()

	return container
}

func (container *Container) Generate() {
	for namespace, s := range *container.M {
		for i, ahc := range s.Adapter {
			if len(*i) > 4 {
				continue
			}

			if len(*ahc) != 0 {
				ctl, act, rel, mod := "", "", "", ""

				for _, v := range *ahc {
					spf := strings.Split(container.toolTP.SourceFuncTP(v), ".")
					if len(spf) < 3 {
						break
					}

					if strings.Contains(container.toolTP.SourceFuncTP(v), "Controller") {
						ctl = container.toolTP.HumpToU(strings.Replace(
							container.toolTP.SourceGrab(spf[1], 2, len(spf[1])-3),
							"Controller", "", -1,
						))
						act = container.toolTP.HumpToU(
							strings.Replace(spf[2], "-fm", "", -1),
						)

						if len(*i) == 2 {
							rel = cast.ToString((*i)[0])
							mod = cast.ToString((*i)[1])

							break
						}

						rel = fmt.Sprintf("/%v/%v", ctl, act)

						if len(*i) >= 3 {
							if (*i)[0] == 0 {
								mod = http.MethodGet
							}

							if (*i)[0] == 1 {
								mod = cast.ToString((*i)[1])
							}

							break
						}

						break
					}
				}

				if _, ok := (*container.arr)[namespace]; ok == false {
					(*container.arr)[namespace] = make(map[*key]*AHC)
				}

				(*container.arr)[namespace][&key{
					srv: namespace,
					ctl: ctl,
					act: act,
					rel: rel,
					mod: mod,
				}] = ahc
			}
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
		panic(err.E(storage.KeyM31007))
	}

	rRel := property.Instance.Get("route.Rel", nil)
	if rRel == nil {
		panic(err.E(storage.KeyM31008))
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

		hc := &AHC{}
		if (*container.M)[namespace.(string)].Middleware != nil {
			hc = (*container.M)[namespace.(string)].Middleware
		}

		container.groups(rRelativePath.(string), r, (*container.arr)[namespace.(string)], hc)
	}
}

func (container *Container) groups(relativePath string, r *gin.Engine, to map[*key]*AHC, hc *AHC) {
	container.to(r.Group(relativePath, *hc...), to)
}

func (container *Container) to(ctx *gin.RouterGroup, to map[*key]*AHC) {
	for x, ctrl := range to {
		add := crypto.New()

		add.Mode = storage.Common
		add.D = []interface{}{storage.Md5, x.srv + x.ctl + x.act}

		k, errRoutesCryptoEngine := add.Engine()
		if errRoutesCryptoEngine != nil {
			panic(errRoutesCryptoEngine)
		}

		(*routeMap)[cast.ToString(k)] = x.rel

		switch x.mod {
		case any:
			ctx.Any(x.rel, *ctrl...)
			continue

		case http.MethodGet:
			ctx.GET(x.rel, *ctrl...)
			continue

		case http.MethodPut:
			ctx.PUT(x.rel, *ctrl...)
			continue

		case http.MethodPost:
			ctx.POST(x.rel, *ctrl...)
			continue

		case http.MethodHead:
			ctx.HEAD(x.rel, *ctrl...)
			continue

		case http.MethodPatch:
			ctx.PATCH(x.rel, *ctrl...)
			continue

		case http.MethodDelete:
			ctx.DELETE(x.rel, *ctrl...)
			continue

		case http.MethodOptions:
			ctx.OPTIONS(x.rel, *ctrl...)
			continue

		default:
			continue
		}
	}
}
