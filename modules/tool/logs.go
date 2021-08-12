// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package tool

import (
	"io"
	"log"
	"os"
)

var (
	Trace	*log.Logger
	Info	*log.Logger
	Warning	*log.Logger
	Error	*log.Logger
)

type Logs struct {
	files *File
}

func NewLogs() *Logs {
	return &Logs {
		files: NewFile(),
	}
}

func (l *Logs) Initialized(name string) error {
	f, e := l.files.OpenFile(name, 0666)
	if e != nil {
		return e
	}

	Trace = l.initializedLog(f, "Trace: ")
	Info = l.initializedLog(f, "Info: ")
	Warning = l.initializedLog(f, "Warning: ")
	Error = l.initializedLog(f, "Error: ")

	return nil
}

func (l *Logs) initializedLog(f *os.File, prefix string) *log.Logger {
	return log.New(io.MultiWriter(f, os.Stderr), prefix, log.Ldate|log.Ltime|log.Lshortfile)
}


