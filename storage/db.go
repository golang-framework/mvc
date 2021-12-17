// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package storage

const (
	SelectOne 	= "ONE"
	SelectAll 	= "ALL"
	ByAsc 		= "ASC"
	ByEsc 		= "DESC"

	JoinLeft 	= "LEFT"
	JoinInner 	= "INNER"
	JoinRight 	= "RIGHT"
)

type (
	Conditions struct {
		Types 		string // Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
		Table 		string
		Field 		[]string
		Joins 		[]*Join
		Limit 		int
		Start 		int
		Query 		interface{}
		QueryArgs 	[]interface{}
		Columns 	[]string
		OrderType 	string
		OrderArgs 	[]string
	}

	Join struct {
		JoinOperator 	string
		TableName 		interface{}
		Condition 		string
	}
)


