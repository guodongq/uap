// Package log provides unified logging abstraction.
package log

import (
	"fmt"
	"runtime/debug"
)

// HandlePanic will log the errors when programer panic
var HandlePanic = func(recovered any, funcName string) {
	Error(fmt.Sprintf("%s panic: %v", funcName, recovered))
	Error(string(debug.Stack()))
}
