// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package mvc

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang-framework/mvc/modules/caches/redis"
	"github.com/golang-framework/mvc/modules/db"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/languages"
	"github.com/golang-framework/mvc/modules/property"
	"github.com/golang-framework/mvc/routes"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"

	_ "github.com/golang-framework/mvc/autoload"
)

type Framework struct {
	Route       *routes.Container
	Err         *err.M
	Translation *languages.M

	FwgLoggerWithFormat gin.HandlerFunc
}

func New() *Framework {
	return &Framework{
		Route: &routes.Container{
			M: &routes.M{},
			E: &routes.AHC{},
		},
		Err: &err.M{
			EMsg: nil,
		},

		FwgLoggerWithFormat: nil,
	}
}

func (fw *Framework) Fw() {
	fw.Err.Initialized()

	(&db.M{}).Engine()
	(&redis.M{}).Engine()
}

func (fw *Framework) FwRouter() {
	if fw.Route == nil {
		panic(err.E(storage.KeyM31006))
	}

	routes.Instance = fw.Route
	routes.Instance.Load().Generate()
}

func (fw *Framework) FwTranslation(d *storage.E) {
	(&languages.M{TMsg: d}).Initialized()
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

	if fw.FwgLoggerWithFormat != nil {
		r.Use(fw.FwgLoggerWithFormat)
	} else {
		r.Use(fw.mvcLoggerWithFormat())
	}

	routes.Instance.Engine(r)

	/**
	 * Https Power ON&OFF -> PoT.Hssl
	 *   -> PoT.Hssl.Power
	 *   -> PoT.Hssl.CertFile
	 *   -> PoT.Hssl.KeysFile
	**/
	var errStartRun error = nil

	port := ":" + cast.ToString(property.Instance.Get("Common.Port", storage.PropertyPort))
	hSsl := property.Instance.Get("Common.Hssl.Power", storage.PropertyHsslPower)
	if hSsl == 1 {
		hSslcf := cast.ToString(property.Instance.Get("PoT.Hssl.CertFile", ""))
		if hSslcf == "" {
			panic(err.E("KeyM31004"))
		}

		hSslkf := cast.ToString(property.Instance.Get("PoT.Hssl.KeysFile", ""))
		if hSslkf == "" {
			panic(err.E("KeyM31005"))
		}

		fmt.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTPs on %v\n", storage.Fw, port)
		errStartRun = r.RunTLS(port, hSslcf, hSslkf)

	} else {
		fmt.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTP on %v\n", storage.Fw, port)
		errStartRun = r.Run(port)

	}

	if errStartRun != nil {
		panic(errStartRun)
	}
}

func (fw *Framework) mvcLoggerWithFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%v | %v | %v | %v | %v | %v | %v | %v\n",
			storage.Fw,
			param.TimeStamp.Format(storage.FwFormatDateTime),
			param.StatusCode,
			param.ClientIP,
			param.Method,
			param.Latency,
			param.Path,
			param.Request.Proto,
		)
	})
}
