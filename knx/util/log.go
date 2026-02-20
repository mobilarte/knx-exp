// Copyright 2017 Ole Krüger.
// Licensed under the MIT license which can be found in the LICENSE file.

package util

import (
	"fmt"
	"reflect"
	"sync"
)

// A LogTarget is used to log certain messages.
type LogTarget interface {
	Printf(format string, args ...any)
}

// Logger is the log target for asynchronous and non-critical errors.
var Logger LogTarget

var (
	longestLogger = 10
	logMutex      sync.Mutex
)

// Log sends a message to the Logger.
func Log(value any, format string, args ...any) {
	if Logger == nil {
		return
	}

	typ := reflect.TypeOf(value).String()

	logMutex.Lock()
	if len(typ) > longestLogger {
		longestLogger = len(typ)
	}
	lLogger := longestLogger // Local copy for formatting
	logMutex.Unlock()

	Logger.Printf(
		fmt.Sprintf("%%%ds[%%p]: %%s\n", lLogger),
		reflect.TypeOf(value).String(), value, fmt.Sprintf(format, args...),
	)
}
