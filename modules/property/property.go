// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package property

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	err "github.com/golang-framework/mvc/modules/error"
	"github.com/golang-framework/mvc/modules/tool"
	"github.com/golang-framework/mvc/storage"
	"github.com/spf13/viper"
)

var Property *M

type M struct {
	Property *viper.Viper

	tools *tool.M
	files *tool.File
}

func New() *M {
	return &M {
		Property: viper.New(),

		tools: tool.New(),
	}
}

func (m *M) Load() *M {
	m.Property.AddConfigPath(".")
	m.Property.SetConfigType(storage.PropertySuffix)

	env := m.Property.GetString("env")
	if env == "" {
		panic(err.E(storage.KeyM31001))
	}

	if ok := m.tools.Contains(env, storage.PropertyEnv ...); ok == false {
		panic(err.E(storage.KeyM31002))
	}

	dir := "./." + env + "." + storage.PropertySuffix

	if ok, _ := m.files.IsExists(dir); ok == false {
		f, e := m.files.Create(dir)
		defer func() {
			f.Close()
		}()

		if e != nil {
			panic(err.E(storage.KeyM31003))
		}

		f.WriteString(m.tpl())
	}

	m.Property.SetConfigName("." + env)
	if err := m.Property.ReadInConfig(); err != nil {
		panic(err)
	}

	m.Property.WatchConfig()
	m.Property.OnConfigChange(func (e fsnotify.Event){
		// Todo: do something ...
	})

	return m
}

func (m *M) Get(key string, val interface{}) interface{} {
	if m.Property.IsSet(key) {
		return m.Property.Get(key)
	}

	return val
}

func (m *M) Usk(key string, val interface{}, opts ...viper.DecoderConfigOption) error {
	return m.Property.UnmarshalKey(key, val, opts ...)
}

func (m *M) tpl() string {
	return fmt.Sprintf(`## %v - %v ##
Common:
  Name: "%v"
  Port: "%v"
  TimeLocation: "%v"
  Addr: ""
  Hssl:
    Power: %d
    CertFile: ""
    KeysFile: ""
`,
storage.Fw,
storage.FwVersion,
storage.Fw,
storage.PropertyPort,
storage.PropertyTimeLocation,
storage.PropertyHsslPower,)
}


