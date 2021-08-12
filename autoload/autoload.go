// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/golang-framework/mvc/modules/property"
	"github.com/spf13/pflag"
)

type autoload struct {
	p *property.M
}

func init() {
	ad := newAutoload()

	/**
	 * Initialized Property
	 *   -> assign the property.Property
	 *   -> Adapters
	**/
	ad.property()
}

func newAutoload() *autoload {
	return &autoload {
		p: property.New(),
	}
}

func (ad *autoload) property() {
	pflag.String("env", "", "environment configure")
	pflag.Parse()

	if e := ad.p.Property.BindPFlags(pflag.CommandLine); e != nil {
		panic(e)
	}

	property.Instance = ad.p.Load()
}

