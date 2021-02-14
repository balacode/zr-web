// -----------------------------------------------------------------------------
// ZR Library - Web Package                                     zr-web/[func.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package web

// # HTTP Functions
//   BaseReferer(req *http.Request) string
//
// # Other Functions
//   CompactCSS(css []byte) []byte
//   FormatURL(url string) string

import (
	"bytes"
	"net/http"
	"strings"
)

// -----------------------------------------------------------------------------
// # HTTP Functions

// BaseReferer _ _
func BaseReferer(req *http.Request) string {
	return strings.Trim(req.Referer(), " \a\b\f\n\r\t\v"+`/\#-0123456789`)
} //                                                                 BaseReferer

// -----------------------------------------------------------------------------
// # Other Functions

// CompactCSS removes all comments, empty
// lines and indentation from CSS bytes.
func CompactCSS(css []byte) []byte {
	ret := make([]byte, 0, len(css))
	for {
		// find position of nearest single-line or multi-line comment
		var (
			c1 = bytes.Index(css, []byte{'/', '/'})
			c2 = bytes.Index(css, []byte{'/', '*'})
			i  = c1
		)
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
			c := bytes.IndexByte(css, '\n')
			if c == -1 {
				break
			}
			css = css[c+1:]
		}
		if c2 > -1 {
			c := bytes.Index(css, []byte{'*', '/'})
			if c == -1 {
				break
			}
			css = css[c+2:]
		}
	}
	// change tabs to spaces and remove repeated spaces and line feeds:
	var (
		lf    = []byte{'\n'}
		lf2   = []byte{'\n', '\n'}
		lfSpc = []byte{'\n', ' '}
		spc   = []byte{' '}
		spc2  = []byte{' ', ' '}
		tab   = []byte{'\t'}
	)
	for bytes.Contains(ret, tab) {
		ret = bytes.ReplaceAll(ret, tab, spc)
	}
	for bytes.Contains(ret, spc2) {
		ret = bytes.ReplaceAll(ret, spc2, spc)
	}
	for bytes.Contains(ret, lfSpc) {
		ret = bytes.ReplaceAll(ret, lfSpc, lf)
	}
	for bytes.Contains(ret, lf2) {
		ret = bytes.ReplaceAll(ret, lf2, lf)
	}
	return ret
} //                                                                  CompactCSS

// FormatURL trims spaces, trailing '#', '/', and '?' from url.
// It also replaces backslashes '\' with forward slashes '/'.
// Note: this function does not perform escaping.
func FormatURL(url string) string {
	url = strings.TrimSpace(url)
	if strings.Contains(url, "\\") {
		url = strings.ReplaceAll(url, "\\", "/")
	}
	url = strings.TrimRight(url, "#/?")
	return url
} //                                                                   FormatURL

// end
