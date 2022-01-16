// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
	"math/rand"
	"net/url"
	"reflect"
	"strings"
	"time"
)

const (
	c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	// Upper Case
	uc = []interface{} {
		"A","B","C","D","E","F","G","H","I","J","K","L","M",
		"N","O","P","Q","R","S","T","U","V","W","X","Y","Z",
	}

	// Lower Case
	lc = []interface{} {
		"a","b","c","d","e","f","g","h","i","j","k","l","m",
		"n","o","p","q","r","s","t","u","v","w","x","y","z",
	}

	// Arabic Numbers
	an = []interface{} {
		"0","1","2","3","4","5","6","7","8","9",
	}

	// Symbol Character
	sc = []interface{} {
		"~","!","@","#","$","%","^","&","*","(",")","_","-",
		"+","=",":",";","'",",",".","<",">","?","/","|","\\",
	}
)

type M struct {

}

func New() *M {
	return &M {

	}
}

func (m *M) SourceFilter(d ... *string) {
	if len(d) < 1 {
		return
	}

	for _, v := range d {
		*v = strings.Trim(*v, " ")
		*v = url.QueryEscape(*v)
	}
}

func (m *M) MatchPattern(v, pattern string, s ... int) int8 {
	if len(s) >= 1 && s[0] > 0 && len(v) < s[0] {
		return -1
	}

	if len(s) >= 2 && s[1] > 0 && len(v) > s[1] {
		return -1
	}

	switch pattern {
	case storage.PatternType01:

		if m.Contains(v,lc ...) == 1 && m.Contains(v,an ...) == 1 {
			return 1
		}
		return -1

	case storage.PatternType02:

		if m.Contains(v,uc ...) == 1 && m.Contains(v,lc ...) == 1 &&
		   m.Contains(v,an ...) == 1 && m.Contains(v,sc ...) == 1 {
			return 1
		}
		return -1

	case storage.PatternType03:

		if m.Contains(v,lc ...) == 1 && m.Contains(v,an ...) == 1 &&
		   m.Contains(v,sc ...) == 1 {
			return 1
		}
		return -1

	default:

		return -1
	}
}

func (m *M) Base64ToEncode(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}

func (m *M) Base64ToDecode(d string) (string, error) {
	decode, err := base64.StdEncoding.DecodeString(d)
	return string(decode), err
}

func (m *M) ToStruct(d, res interface{}) error {
	return json.Unmarshal([]byte(cast.ToString(d)), res)
}

func (m *M) ToJson(d interface{}) (string, error) {
	res, e := json.Marshal(d)
	return string(res), e
}

func (m *M) RandomStr(size int) string {
	rand.NewSource(time.Now().UnixNano())

	var res bytes.Buffer
	for i := 0; i < size; i ++ {
		res.WriteByte(c[rand.Int63() % int64(len(c))])
	}

	return res.String()
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