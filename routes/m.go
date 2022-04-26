// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

const (
	any                 = "ANY"
	DefaultRelativePath = "{_}"
)

var routeMap = &storage.Y{}

func AiANY(relativePath, method interface{}) *I {
	return ai(relativePath, any)
}

func AiGET(relativePath string) *I {
	return ai(relativePath, http.MethodGet)
}

func AiPUT(relativePath string) *I {
	return ai(relativePath, http.MethodPut)
}

func AiPOST(relativePath string) *I {
	return ai(relativePath, http.MethodPost)
}

func AiHEAD(relativePath string) *I {
	return ai(relativePath, http.MethodHead)
}

func AiPATCH(relativePath string) *I {
	return ai(relativePath, http.MethodPatch)
}

func AiDELETE(relativePath string) *I {
	return ai(relativePath, http.MethodDelete)
}

func AiOPTIONS(relativePath string) *I {
	return ai(relativePath, http.MethodOptions)
}

func Path(srv, ctl, act string) (interface{}, error) {
	add := crypto.New()

	add.Mode = storage.Common
	add.D = []interface{}{storage.Md5, srv + ctl + act}

	k, errRoutesCryptoEngine := add.Engine()
	if errRoutesCryptoEngine != nil {
		return nil, errRoutesCryptoEngine
	}

	_, ok := (*routeMap)[cast.ToString(k)]
	if ok == false {
		return nil, nil
	}

	return (*routeMap)[cast.ToString(k)], nil
}

func ai(d ...interface{}) *I {
	if len(d) == 0 {
		return &I{0, time.Now().UnixNano(), rand.Intn(100000)}
	}

	if len(d) == 1 {
		if d[0] == DefaultRelativePath {
			return &I{0, time.Now().UnixNano(), rand.Intn(100000)}
		} else {
			return &I{d[0], http.MethodGet}
		}
	}

	if len(d) == 2 {
		if d[0] == DefaultRelativePath {
			return &I{1, d[1], time.Now().UnixNano(), rand.Intn(100000)}
		} else {
			return &I{d[0], d[1]}
		}
	} else {
		return &I{0, time.Now().UnixNano(), rand.Intn(100000)}
	}
}
