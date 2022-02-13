// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/cast"
)

const (
	c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	// Upper Case
	uc = []interface{}{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	}

	// Lower Case
	lc = []interface{}{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	}

	// Arabic Numbers
	an = []interface{}{
		"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	}

	// Symbol Character
	sc = []interface{}{
		"~", "!", "@", "#", "$", "%", "^", "&", "*", "(", ")", "_", "-",
		"+", "=", ":", ";", "'", ",", ".", "<", ">", "?", "/", "|", "\\",
	}
)

type M struct {
}

func New() *M {
	return &M{}
}

func (m *M) MatchPattern(v, pattern string, s ...int) int8 {
	if len(s) >= 1 && s[0] > 0 && len(v) < s[0] {
		return -1
	}

	if len(s) >= 2 && s[1] > 0 && len(v) > s[1] {
		return -1
	}

	switch pattern {
	case storage.PatternType01:

		if m.Contains(v, lc...) == 1 && m.Contains(v, an...) == 1 {
			return 1
		}
		return -1

	case storage.PatternType02:

		if m.Contains(v, uc...) == 1 && m.Contains(v, lc...) == 1 &&
			m.Contains(v, an...) == 1 && m.Contains(v, sc...) == 1 {
			return 1
		}
		return -1

	case storage.PatternType03:

		if m.Contains(v, lc...) == 1 && m.Contains(v, an...) == 1 &&
			m.Contains(v, sc...) == 1 {
			return 1
		}
		return -1

	default:

		return -1
	}
}

func (_ *M) SourceGrab(d string, start, length int) string {
	s := []rune(d)
	l := len(s)
	z := 0

	if start < 0 {
		start = l - 1 + start
	}

	z = start + length

	if start > z {
		start, z = z, start
	}
	if start < 0 {
		start = 0
	}
	if start > length {
		start = l
	}

	if z < 0 {
		z = 0
	}
	if z > l {
		z = l
	}

	return string(s[start:z])
}

func (_ *M) SourceFuncTP(d interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(d).Pointer()).Name()
}

func (_ *M) SourceFilter(d ...*string) {
	if len(d) < 1 {
		return
	}

	for _, v := range d {
		*v = strings.Trim(*v, " ")
	}
}

func (_ *M) Base64ToEncode(d []byte) string {
	return base64.StdEncoding.EncodeToString(d)
}

func (_ *M) Base64ToDecode(d string) (string, error) {
	decode, errToolBase64ToDecode := base64.StdEncoding.DecodeString(d)
	return string(decode), errToolBase64ToDecode
}

func (_ *M) UToHump(d string) string {
	d = strings.Replace(d, "_", " ", -1)
	d = strings.Title(d)

	return strings.Replace(d, " ", "", -1)
}

func (m *M) HumpToU(d string) string {
	var (
		ul = make([]string, 0)
		hp = []rune(d)
	)

	if len(hp) == 0 {
		return ""
	}

	for i := 0; i < len(hp); i++ {
		if m.Contains(string(hp[i]), uc...) == 1 {
			ul = append(ul, "_")
			ul = append(ul, strings.ToLower(string(hp[i])))
		} else {
			ul = append(ul, string(hp[i]))
		}
	}

	if len(hp) > 0 && ul[0] == "_" {
		ul = append(ul[:0], ul[1:]...)
	}

	return strings.Join(ul, "")
}

func (_ *M) ToStruct(d, res interface{}) error {
	return json.Unmarshal([]byte(cast.ToString(d)), res)
}

func (_ *M) ToJson(d interface{}) (string, error) {
	res, e := json.Marshal(d)
	return string(res), e
}

func (_ *M) RandomStr(size int) string {
	rand.NewSource(time.Now().UnixNano())

	var res bytes.Buffer
	for i := 0; i < size; i++ {
		res.WriteByte(c[rand.Int63()%int64(len(c))])
	}

	return res.String()
}

func (_ *M) Contains(k interface{}, d ...interface{}) int8 {
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

func (_ *M) ContainSliceIndex(src interface{}, k int) int8 {
	switch reflect.TypeOf(src).Kind() {
	case reflect.Slice:
		var res = reflect.ValueOf(src)
		for i := 0; i < res.Len(); i++ {
			if i == k {
				return 1
			}
		}
		return -1

	default:
		return -1
	}
}
