// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

import (
	"net/http"
)

const (
	StatusOK = http.StatusOK

	StatusUnknown = -1
	StatusNotFound = http.StatusNotFound
	StatusNoContent = http.StatusNoContent
)
