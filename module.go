// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-18 17:41:25 9354FF                             zr-web/[module.go]
// -----------------------------------------------------------------------------

// Package web provides HTML generation functions
// and HTTP session management.
package web

import (
	"fmt"

	"github.com/balacode/zr"
)

var (
	// PL is fmt.Println() but is used only for debugging.
	PL = fmt.Println

	// VL is zr.VerboseLog() but is used only for debugging.
	VL = zr.VerboseLog
)

//end
