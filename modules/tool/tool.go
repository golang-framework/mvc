// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"github.com/spf13/cast"
	"reflect"
	"strings"
)

type M struct {

}

func New() *M {
	return &M {}
}

func (m *M) Contains(k interface{}, d ... interface{}) int8 {
	if len(d) < 1 {
		return -1
	}

	for _, val := range d {
		if strings.Contains(cast.ToString(k), cast.ToString(val)) {
			return 1
		}
	}

	return -1
}

func (m *M) ContainSliceIndex(src interface{}, k int) int8 {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		var res = reflect.ValueOf(src)
		for i := 0; i < res.Len(); i ++ {
			if i == k {
				return 1
			}
		}
		return -1

	default:
		return -1
	}
}