// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-24 13:36:58 F570C1                          zr-web/[html_test.go]
// -----------------------------------------------------------------------------

package web

import (
	"testing"

	"github.com/balacode/zr"
)

/*
to test all items in html.go use:
    go test --run Test_html_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

// go test --run Test_html_SetClass_
func Test_html_SetClass_(t *testing.T) {
	zr.TBegin(t)
	// SetClass(add bool, input string, classes ...string) string
	//
	// set classes on a blank string
	zr.TEqual(t, SetClass(true, "", "A", "B", "C"), ("A B C"))
	zr.TEqual(t, SetClass(true, "", "A", "BB", "CCC"), ("A BB CCC"))
	zr.TEqual(t, SetClass(true, "", "AA", "BB", "CC"), ("AA BB CC"))
	//
	// cases where the added classes are already existing
	zr.TEqual(t, SetClass(true, "A B C", "A"), ("A B C"))
	zr.TEqual(t, SetClass(true, "A B C", "A", "B"), ("A B C"))
	zr.TEqual(t, SetClass(true, "A B C", "A", "B", "C"), ("A B C"))
	zr.TEqual(t, SetClass(true, "A BB CCC", "A"), ("A BB CCC"))
	zr.TEqual(t, SetClass(true, "A BB CCC", "A", "BB"), ("A BB CCC"))
	zr.TEqual(t, SetClass(true, "A BB CCC", "A", "BB", "CCC"), ("A BB CCC"))
	zr.TEqual(t, SetClass(true, "AA BB CC", "AA"), ("AA BB CC"))
	zr.TEqual(t, SetClass(true, "AA BB CC", "AA", "BB"), ("AA BB CC"))
	zr.TEqual(t, SetClass(true, "AA BB CC", "AA", "BB", "CC"), ("AA BB CC"))
	//
	// remove some classes:
	zr.TEqual(t, SetClass(false, "A B C", "A"), ("B C"))
	zr.TEqual(t, SetClass(false, "A B C", "B"), ("A C"))
	zr.TEqual(t, SetClass(false, "A B C", "C"), ("A B"))
	zr.TEqual(t, SetClass(false, "A BB CCC", "A"), ("BB CCC"))
	zr.TEqual(t, SetClass(false, "A BB CCC", "BB"), ("A CCC"))
	zr.TEqual(t, SetClass(false, "A BB CCC", "CCC"), ("A BB"))
	zr.TEqual(t, SetClass(false, "AA BB CC", "AA"), ("BB CC"))
	zr.TEqual(t, SetClass(false, "AA BB CC", "BB"), ("AA CC"))
	zr.TEqual(t, SetClass(false, "AA BB CC", "CC"), ("AA BB"))
	//
	// remove everything:
	zr.TEqual(t, SetClass(false, "A B C", "A", "B", "C"), (""))
	zr.TEqual(t, SetClass(false, "A B C", "A", "B", "C", "X"), (""))
	zr.TEqual(t, SetClass(false, "A BB CCC", "A", "BB", "CCC"), (""))
	zr.TEqual(t, SetClass(false, "A BB CCC", "A", "BB", "CCC", "X"), (""))
	zr.TEqual(t, SetClass(false, "AA BB CC", "AA", "BB", "CC"), (""))
	zr.TEqual(t, SetClass(false, "AA BB CC", "AA", "BB", "CC", "X"), (""))
} //                                                         Test_html_SetClass_

//end
