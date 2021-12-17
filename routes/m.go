// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

const any = "ANY"
var routeMap = &storage.Y{}

func Path(srv, ctl, act string) (interface{}, error) {
	add := crypto.New()

	add.Mode = storage.Common
	add.D = []interface{}{storage.Md5, srv + ctl + act}

	k, e := add.Engine()
	if e != nil {
		return nil, e
	}

	_, ok := (*routeMap)[cast.ToString(k)]
	if ok == false {
		return nil, nil
	}

	return (*routeMap)[cast.ToString(k)], nil
}