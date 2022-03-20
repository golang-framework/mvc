// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package languages

import (
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func T(k, ln string, replace ...interface{}) string {
	return message.NewPrinter(language.MustParse(ln)).Sprintf(k, replace...)
}

type M struct {
	TMsg *storage.E
}

func (m *M) Initialized() {
	if m.TMsg == nil || len(*m.TMsg) <= 0 {
		return
	}

	for ln, translation := range *m.TMsg {
		for k, v := range translation {
			message.SetString(language.MustParse(ln), k, cast.ToString(v))
		}
	}
}
