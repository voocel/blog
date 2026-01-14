package util

import (
	"blog/pkg/log"
	"runtime/debug"
)

// SafeGo executes a function in a goroutine with panic recovery.
// If the goroutine panics, it logs the error and stack trace instead of crashing the application.
func SafeGo(fn func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Errorw("Goroutine panic recovered",
					log.Pair("panic", r),
					log.Pair("stack", string(debug.Stack())),
				)
			}
		}()
		fn()
	}()
}
