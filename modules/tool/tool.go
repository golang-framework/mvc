// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"github.com/spf13/cast"
	"strings"
)

type M struct {

}

func New() *M {
	return &M {}
}

func (m *M) Contains(k interface{}, d ... interface{}) bool {
	if len(d) < 1 {
		return false
	}

	for _, val := range d {
		if strings.Contains(cast.ToString(k), cast.ToString(val)) {
			return true
		}
	}

	return false
}
