// -----------------------------------------------------------------------------
// ZR Library - Web Package                                   zr-web/[buffer.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package web

// # Constructor
//   Buffer struct
//   Embed(html string) Buffer
//   NewBuffer(size int) Buffer
//
// # Methods (ob *Buffer)
//   ) Bytes() []byte
//   ) Len() int
//   ) Reset()
//   ) String() string
//   ) Write(html ...*Buffer)
//   ) WriteBytes(arrays ...[]byte)
//   ) WriteString(strings ...string)

import (
	"bytes"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Constructor

// Buffer represents a buffer holding HTML output.
type Buffer struct {
	html bytes.Buffer
} //                                                                      Buffer

// Embed creates a Buffer from a HTML string.
func Embed(html string) Buffer {
	var ret Buffer
	ret.html = *bytes.NewBuffer([]byte(html))
	return ret
} //                                                                       Embed

// NewBuffer creates and initializes a new Buffer of the specified size.
// It is intended to prepare a buffer for writing HTML markup.
// In most cases, new(Buffer) (or just declaring a Buffer variable)
// is sufficient to initialize a Buffer.
func NewBuffer(size int) Buffer {
	var ret Buffer
	ret.html = *bytes.NewBuffer(make([]byte, 0, size))
	return ret
} //                                                                   NewBuffer

// -----------------------------------------------------------------------------
// # Methods (ob *Buffer)

// Bytes returns the contents of the buffer as an array of bytes.
func (ob *Buffer) Bytes() []byte {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return []byte{}
	}
	return ob.html.Bytes()
} //                                                                       Bytes

// Len returns the length of the buffer.
func (ob *Buffer) Len() int {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return 0
	}
	return ob.html.Len()
} //                                                                         Len

// Reset makes the buffer empty, but retains
// allocated storage for future reuse.
func (ob *Buffer) Reset() {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	ob.html.Reset()
} //                                                                       Reset

// String returns the contents of the buffer as a string
// and implements the fmt.Stringer interface.
func (ob *Buffer) String() string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	return ob.html.String()
} //                                                                      String

// Write writes 'html' to the buffer.
func (ob *Buffer) Write(html ...*Buffer) {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	for _, iter := range html {
		if iter != nil && iter.Len() > 0 {
			ob.html.Write(iter.Bytes())
		}
	}
} //                                                                       Write

// WriteBytes writes a string to the buffer.
func (ob *Buffer) WriteBytes(arrays ...[]byte) {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	for _, b := range arrays {
		ob.html.Write(b)
	}
} //                                                                  WriteBytes

// WriteString writes a string to the buffer.
func (ob *Buffer) WriteString(strings ...string) {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	for _, s := range strings {
		ob.html.WriteString(s)
	}
} //                                                                 WriteString

//end
