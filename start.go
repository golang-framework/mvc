// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mvc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/routes"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"

	_ "github.com/golang-framework/mvc/autoload"
)

type Framework struct {
	Route *routes.Container
	Err *err.M
}

func New() *Framework {
	return &Framework {

	}
}

func (fw *Framework) Run() {
	// console -> golang framework version
	fmt.Fprintf(gin.DefaultWriter, "[%v] %v\n", storage.Fw, storage.FwVersion)

	// console -> Disable Console Color
	gin.DisableConsoleColor()

	// configs -> gin mode release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(fw.loadLoggerWithFormat())

	routes.Instance.Engine(r)

	/**
	 * Https Power ON&OFF -> PoT.Hssl
	 *   -> PoT.Hssl.Power
	 *   -> PoT.Hssl.CertFile
	 *   -> PoT.Hssl.KeysFile
	**/
	var e error = nil

	port := ":" + cast.ToString(property.Property.Get("Common.Port", storage.PropertyPort))
	hSsl := property.Property.Get("Common.Hssl.Power", storage.PropertyHsslPower)
	if hSsl == 1 {
		hSslcf := cast.ToString(property.Property.Get("PoT.Hssl.CertFile", ""))
		if hSslcf == "" {
			panic(err.E("KeyM31004"))
		}

		hSslkf := cast.ToString(property.Property.Get("PoT.Hssl.KeysFile", ""))
		if hSslkf == "" {
			panic(err.E("KeyM31005"))
		}

		fmt.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTPs on %v\n", storage.Fw, port)
		e = r.RunTLS(port, hSslcf, hSslkf)

	} else {
		fmt.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTP on %v\n", storage.Fw, port)
		e = r.Run(port)

	}

	if e != nil {
		panic(e)
	}
}

func (fw *Framework) loadLoggerWithFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[%v] %v |	%v |	%v |	%v |	%v |	%v(%v)\n",
			storage.Fw,
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
		)
	})
}
