// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/spf13/cast"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

const (
	c = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type M struct {

}

func New() *M {
	return &M {

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