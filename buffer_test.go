// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-24 13:36:58 F9EB16                        [zr-web/buffer_test.go]
// -----------------------------------------------------------------------------

package web

// # Builders
//   Test_buff_Embed_
//   Test_buff_NewBuffer_
//
// # Methods
//   Test_buff_Buffer_Bytes_
//   Test_buff_Buffer_Len_
//   Test_buff_Buffer_Reset_
//   Test_buff_Buffer_String_
//   Test_buff_Buffer_Write_
//   Test_buff_Buffer_WriteBytes_
//   Test_buff_Buffer_WriteString_

/*
to test all items in buffer.go use:
    go test --run Test_buff_

to generate a test coverage report for the whole module use:
    go test -coverprofile cover.out
    go tool cover -html=cover.out
*/

import (
	// "bytes"
	"testing"

	"github.com/balacode/zr"
)

// -----------------------------------------------------------------------------
// # Builders

// go test --run Test_buff_Embed_
func Test_buff_Embed_(t *testing.T) {
	zr.TBegin(t)
	// Embed_(html string) Buffer
	//
	{
		var buf = Embed("")
		zr.TEqual(t, buf.String(), "")
	}
	{
		var buf = Embed("ABC")
		zr.TEqual(t, buf.String(), "ABC")
	}
} //                                                     Test_buff_Buffer_Embed_

// go test --run Test_buff_Buffer_NewBuffer_
func Test_buff_NewBuffer_(t *testing.T) {
	zr.TBegin(t)
	// NewBuffer_(size int) Buffer
	//
} //                                                 Test_buff_Buffer_NewBuffer_

// -----------------------------------------------------------------------------
// # Methods

// go test --run Test_buff_Buffer_Bytes_
func Test_buff_Buffer_Bytes_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) Bytes() []byte
	//
	//TODO: add check for nil receiver
} //                                                     Test_buff_Buffer_Bytes_

// go test --run Test_buff_Buffer_Len_
func Test_buff_Buffer_Len_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) Len() int
	//
	//TODO: add check for nil receiver
} //                                                       Test_buff_Buffer_Len_

// go test --run Test_buff_Buffer_Reset_
func Test_buff_Buffer_Reset_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) Reset()
	//
	//TODO: add check for nil receiver
} //                                                     Test_buff_Buffer_Reset_

// go test --run Test_buff_Buffer_String_
func Test_buff_Buffer_String_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) String() string
	//
	//TODO: add check for nil receiver
} //                                                    Test_buff_Buffer_String_

// go test --run Test_buff_Buffer_Write_
func Test_buff_Buffer_Write_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) Write(html ...*Buffer)
	//
	//TODO: add check for nil receiver
} //                                                     Test_buff_Buffer_Write_

// go test --run Test_buff_Buffer_WriteBytes_
func Test_buff_Buffer_WriteBytes_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) WriteBytes(arrays ...[]byte)
	//
	//TODO: add check for nil receiver
} //                                                Test_buff_Buffer_WriteBytes_

// go test --run Test_buff_Buffer_WriteString_
func Test_buff_Buffer_WriteString_(t *testing.T) {
	zr.TBegin(t)
	// (ob *Buffer) WriteString(strings ...string)
	//
	//TODO: add check for nil receiver
} //                                               Test_buff_Buffer_WriteString_

//end
