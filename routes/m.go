// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/golang-framework/mvc/modules/crypto"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

const (
	Any	= "ANY"
	Get = "GET"
	Put = "PUT"
	Head = "HEAD"
	Post = "POST"
	Patch = "PATCH" // RFC 5789
	Trace = "TRACE"
	Delete = "DELETE"
	Connect = "CONNECT"
	Options = "OPTIONS"
)

var routeMap = &storage.Y{}

func Path(ctl, srv, act string) (interface{}, error) {
	add := crypto.New()

	add.Mode = storage.Common
	add.D = []interface{}{storage.Md5, ctl + srv + act}

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