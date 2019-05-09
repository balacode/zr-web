// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-09 17:19:39 A87E90                             zr-web/[module.go]
// -----------------------------------------------------------------------------

// Package web provides HTML generation functions
// and HTTP session management.
package web

import (
	"fmt"

	"github.com/balacode/zr"
)

// SPACES is a string of all white-space characters,
// which includes spaces, tabs, and newline characters.
const SPACES = " \a\b\f\n\r\t\v"

// PL is fmt.Println() but is used only for debugging.
var PL = fmt.Println

// VL is zr.VerboseLog() but is used only for debugging.
var VL = zr.VerboseLog

//end
