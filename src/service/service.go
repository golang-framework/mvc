// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package service

import (
	"math"

	"github.com/golang-framework/mvc/storage"
)

const (
	defaultPage int = 1
	defaultSize int = 10
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (_ *Service) Paginator(page, nums, size int) *storage.Y {
	var previousPage int
	var nextPage int

	if size <= 0 {
		size = defaultSize
	}

	var totalPage int = int(math.Ceil(float64(nums) / float64(size)))

	if page > totalPage {
		page = totalPage
	}

	if page <= 0 {
		page = 1
	}

	var pages []int

	switch {
	case page >= totalPage-5 && totalPage > 5:
		start := totalPage - 5 + 1
		previousPage = page - 1
		nextPage = int(math.Min(float64(totalPage), float64(page+1)))
		pages = make([]int, 5)
		for i, _ := range pages {
			pages[i] = start + i
		}

	case page >= 3 && totalPage > 5:
		start := page - 3 + 1
		pages = make([]int, 5)
		previousPage = page - 3
		for i, _ := range pages {
			pages[i] = start + i
		}

		previousPage = page - 1
		nextPage = page + 1

	default:
		pages = make([]int, int(math.Min(5, float64(totalPage))))
		for i, _ := range pages {
			pages[i] = i + 1
		}

		previousPage = int(math.Max(float64(1), float64(page-1)))
		nextPage = page + 1
	}

	return &storage.Y{
		"pages":         pages,
		"total":         totalPage,
		"previous_page": previousPage,
		"next_page":     nextPage,
		"front_page":    page,
	}
}
