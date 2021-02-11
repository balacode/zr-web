// -----------------------------------------------------------------------------
// ZR Library - Web Package                                   zr-web/[module.go]
// (c) balarabe@protonmail.com                                      License: MIT
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
