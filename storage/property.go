// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

const (
	// Property File Suffix
	EnvDEV = "dev"
	EnvSTG = "stg"
	EnvPRD = "prd"
	PropertySuffix = "yaml"

	PropertyPort = "8577"
	PropertyTimeLocation = "Asia/Shanghai"
	PropertyHsslPower int = 0
)

var (
	/**
	 * Modules: Property
	 *   - dev
	 *   - stg
	 *   - prd
	 */
	PropertyEnv []interface{} = []interface{}{
		EnvDEV,
		EnvSTG,
		EnvPRD,
	}
)
