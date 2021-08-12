// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"fmt"
	"github.com/spf13/cast"
)

/**
 * modules:crypto -> M33001 ~ M33999
**/
const (
	ErrPrefix string = "go_mvc_error_message_w"

	SuccessOK string = "SuccessOK"
	Incorrect string = "Incorrect"

	// Modules -> M31001 ~ M31999 : property
	KeyM31001 = "modules_property_property_load_01"
	KeyM31002 = "modules_property_property_load_02"
	KeyM31003 = "modules_property_property_load_03"
	KeyM31004 = "fw_start_run_01"
	KeyM31005 = "fw_start_run_02"
	KeyM31006 = "fw_start_mvc_01"
	KeyM31007 = "fw_route_r_engine_01"

	valM31001 = "Environment Error"
	valM31002 = "Environment Error Exclude"
	valM31003 = "Failure to Create the Environment Yaml file"
	valM31004 = "Https Hssl.CertFile Empty"
	valM31005 = "Https Hssl.KeysFile Empty"
	valM31006 = "Empty Routes"
	valM31007 = "Empty Routes"

	// Error Message for Component -> M32001 ~ M32999
	KeyM32001 = "components_caches_redis_singleton_check_01"
	KeyM32002 = "components_db_singleton_select_01"
	KeyM32003 = "components_db_singleton_select_02"
	KeyM32004 = "components_db_singleton_select_03"

	valM32001 = "Missing Redis Client"
	valM32002 = "Missing DB Table"
	valM32003 = "Error Order Type"
	valM32004 = "Error Select Type"

	// Error Message for Modules -> M33001 ~ M33999
	KeyM33001 = "modules_crypto_common_engine_01"
	KeyM33002 = "modules_crypto_common_engine_02"
	KeyM33003 = "modules_crypto_crypto_engine_01"
	KeyM33004 = "modules_caches_redis_singleton_initialized_01"
	KeyM33005 = "modules_caches_redis_r_01"
	KeyM33006 = "modules_db_singleton_initialized_01"
	KeyM33007 = "modules_db_Engine_01"

	valM33001 = "Parameters Error"
	valM33002 = "Error Crypto Common Engine Type"
	valM33003 = "Error Crypto Engine Type "
	valM33004 = "Empty Redis Parameters"
	valM33005 = "Empty Redis Client"
	valM33006 = "Empty db Engine"
	valM33007 = "Empty db Engine"
)

var msg *E = &E {
	ErrPrefix: {
		SuccessOK: "Success",
		Incorrect: "Unknown",

		KeyM31001: valM31001,
		KeyM31002: valM31002,
		KeyM31003: valM31003,
		KeyM31004: valM31004,
		KeyM31005: valM31005,
		KeyM31006: valM31006,
		KeyM31007: valM31007,

		KeyM32001: valM32001,
		KeyM32002: valM32002,
		KeyM32003: valM32003,
		KeyM32004: valM32004,

		KeyM33001: valM33001,
		KeyM33002: valM33002,
		KeyM33003: valM33003,
		KeyM33004: valM33004,
		KeyM33005: valM33005,
		KeyM33006: valM33006,
		KeyM33007: valM33007,
	},
}

func SetError(add *E) {
	if add != nil {
		for k, v := range *add {
			if k != ErrPrefix {
				(*msg)[k] = v
			}
		}
	}
}

func GetError(pfx, k string, content ... interface{}) string {
	str := cast.ToString((*msg)[ErrPrefix][Incorrect])

	s, ok := (*msg)[pfx][k]
	if ok {
		str = cast.ToString(s)
		if len(content) != 0 {
			str = str + "," + fmt.Sprint(content ...)
		}
	}

	return str
}
