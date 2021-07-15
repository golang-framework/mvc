// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/golang-framework/mvc/storage"
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

var routeMap *storage.Y