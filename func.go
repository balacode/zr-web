// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-24 13:36:58 A9F043                               [zr-web/func.go]
// -----------------------------------------------------------------------------

package web

// # HTTP Functions
//   BaseReferer(req *http.Request) string
//
// # Other Functions
//   CompactCSS(css []byte) []byte

import (
	"bytes"
	"net/http"
	str "strings"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # HTTP Functions

// BaseReferer __
func BaseReferer(req *http.Request) string {
	return str.Trim(req.Referer(), zr.SPACES+`/\#-0123456789`)
} //                                                                 BaseReferer

// -----------------------------------------------------------------------------
// # Other Functions

// CompactCSS removes all comments, empty
// lines and indentation from CSS bytes.
func CompactCSS(css []byte) []byte {
	var ret = make([]byte, 0, len(css))
	for {
		// find position of nearest single-line or multi-line comment
		var c1 = bytes.Index(css, []byte{'/', '/'})
		var c2 = bytes.Index(css, []byte{'/', '*'})
		var i = c1
		if (c2 > -1 && c2 < c1) || c1 == -1 {
			i = c2
			c1 = -1
		} else {
			c2 = -1
		}
		// append the part before the comment
		if i == -1 {
			ret = append(ret, css...)
			break
		}
		ret = append(ret, css[:i]...)
		css = css[i+2:]
		// skip comment body
		if c1 > -1 {
			var c = bytes.IndexByte(css, '\n')
			if c == -1 {
				break
			}
			css = css[c+1:]
		}
		if c2 > -1 {
			var c = bytes.Index(css, []byte{'*', '/'})
			if c == -1 {
				break
			}
			css = css[c+2:]
		}
	}
	// change tabs to spaces and remove repeated spaces and line feeds:
	var lf = []byte{'\n'}
	var lf2 = []byte{'\n', '\n'}
	var lfSpc = []byte{'\n', ' '}
	var spc = []byte{' '}
	var spc2 = []byte{' ', ' '}
	var tab = []byte{'\t'}
	for bytes.Contains(ret, tab) {
		ret = bytes.Replace(ret, tab, spc, -1)
	}
	for bytes.Contains(ret, spc2) {
		ret = bytes.Replace(ret, spc2, spc, -1)
	}
	for bytes.Contains(ret, lfSpc) {
		ret = bytes.Replace(ret, lfSpc, lf, -1)
	}
	for bytes.Contains(ret, lf2) {
		ret = bytes.Replace(ret, lf2, lf, -1)
	}
	return ret
} //                                                                  CompactCSS

//end
