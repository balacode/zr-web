// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-04-28 18:07:08 C6CC3F                            zr-web/[context.go]
// -----------------------------------------------------------------------------

package web

/*
	Context wraps an HTTP request and reply. Its methods such as HREF(),
	Method() and PostData() provide request details. The Reply() method
	is used to send a reply. The most important use of Context is that
	it connects request to sessions. It checks the cookie sent with the
	request and starts a new session or continues an existing session.

	import (
		"net/http"
		"log"
		web "github.com/balacode/zr-web"
	)

	func main() {
		http.HandleFunc("/", mainServe)
		log.Fatal(http.ListenAndServe("localhost:888", nil))
	}

	func mainServe(w http.ResponseWriter, req *http.Request) {
		ctx := web.NewContext(w, req, &ob.Sessions)
		ctx.Reply("<html><h1>Hello World</hq></html>", "html")
	}
*/

//  Context struct
//
// # Constructor
//    NewContext(w http.ResponseWriter, req *http.Request, sess *Sessions,
//        ) Context
//
// # Request Properties (ob *Context)
//   BaseReferer() string
//   Method() string
//   HREF() string
//   PostData() []byte
//   Referer() string
//
// # Reply Method (ob *Context)
//   Reply(data []byte, mediaType string)
//
// # Debug Helper Method
//   (ob *Context) DebugString() string {

// # Support (File Scope)
//   readPostData(req *http.Request) []byte

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"net/http"
	"strings"
	"sync"

	"github.com/balacode/zr"
)

// ContextDebugFunc specifies a function for testing or debugging
// all HTTP traffic received and sent by Context. It is called
// when a request is received, and when a reply is sent.
//
// During normal program operation it should not be assigned
// to avoid serious security and performance issues.
//
// To print to the console, assign fmt.Print to it, or use your
// own custom function if you need to log the output elsewhere
//
// The return values are not used, just specified to match fmt.Print
//
// During normal opration, remove the followig assignment ---> = fmt.Print
var ContextDebugFunc func(a ...interface{}) (n int, err error)

// contextDebugPrint calls ContextDebugFunc (if assigned) and
// sends a loud reminder of this activity to the console.
var contextDebugPrint = func(a ...interface{}) {
	if ContextDebugFunc == nil {
		return
	}
	if !contextDebugWarn {
		contextDebugWarn = true
		for i := 0; i < 10; i++ {
			fmt.Println("CONTEXT HTTP DEBUGGING!")
		}
	}
	ContextDebugFunc(a...)
}

// contextDebugWarn specifies if context debugging warning was issued
var contextDebugWarn = false

// contextDebugMutex ensures every request/reply debug entry is fully made
var contextDebugMutex sync.RWMutex

// nextContextID __
var nextContextID int64

// Context structure wraps a HTTP request and attaches a Session
// object to allow state to be maintained between requests.
type Context struct {
	Session  *Session
	id       int64               // serial number
	w        http.ResponseWriter // an interface, so shouldn't be a pointer
	req      *http.Request
	postData []byte
} //                                                                     Context

// -----------------------------------------------------------------------------
// # Constructor

// NewContext creates a new Context object by wrapping a response writer (w)
// and request (req) from ListenAndServe() and attaching a new or existing
// session using 'sess'. Sessions uses cookies to determine if the
// request is a new session, or a previously started session.
func NewContext(w http.ResponseWriter, req *http.Request, sess *Sessions,
) Context {
	nextContextID++
	ret := Context{
		id:  nextContextID,
		req: req,
		w:   w,
	}
	if sess != nil {
		ret.Session = sess.GetByCookie(w, req)
	}
	// if not debugging, return immediately
	if ContextDebugFunc == nil {
		return ret
	}
	// write request's details for debugging
	contextDebugMutex.Lock() // unlocked by Reply()
	const LE = " \n"         // line end
	var postdata string
	{
		data := ret.PostData()
		if len(data) > 0 {
			if zr.DebugMode() {
				postdata = string(data)
			} else {
				postdata = zr.DebugString(data)
			}
			postdata = strings.Trim(postdata, zr.SPACES)
			postdata = "postdata:" + LE + strings.Repeat("-", 79) + ">" + LE +
				postdata + LE + "<" + strings.Repeat("-", 79) + LE
		}
	}
	var referer string
	if req.Referer() != "" {
		referer = "ref:" + req.Referer() + LE
	}
	contextDebugPrint(
		"REQUEST:", ret.id, " sid:", ret.Session.ID()[:8],
		" path:", req.Method, " ", req.URL.Path, LE,
		referer,
		postdata,
		LE,
	)
	return ret
} //                                                                  NewContext

// -----------------------------------------------------------------------------
// # Request Properties (ob *Context)

// BaseReferer property returns the base referer path of
// the current request. I.e. a path with '/', '\', '#'
// and numbers stripped from the end.
func (ob *Context) BaseReferer() string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := strings.TrimRight(ob.Referer(), SPACES+`/\#-0123456789`)
	return ret
} //                                                                 BaseReferer

// Method returns the request's HTTP method ('GET', 'POST', 'PUT', etc)
func (ob *Context) Method() string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	return strings.ToUpper(ob.req.Method)
} //                                                                      Method

// HREF property returns the URL path of the current request.
func (ob *Context) HREF() string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := strings.Trim(ob.req.URL.Path, `#/\ `)
	if strings.Contains(ret, "\\") {
		ret = strings.Replace(ret, "\\", "/", -1)
	}
	return ret
} //                                                                        HREF

// PostData property returns the POSTDATA of the current request.
func (ob *Context) PostData() []byte {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return []byte{}
	}
	if len(ob.postData) > 0 {
		return ob.postData
	}
	ob.postData = readPostData(ob.req)
	return ob.postData
} //                                                                    PostData

// Referer property returns the referer path of the current request.
func (ob *Context) Referer() string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := strings.Trim(ob.req.Referer(), `#/\ `)
	if strings.Contains(ret, "\\") {
		ret = strings.Replace(ret, "\\", "/", -1)
	}
	return ret
} //                                                                     Referer

// -----------------------------------------------------------------------------
// # Reply Method (ob *Context)

// Reply method sends the reply to a request.
// Specify 'mediaType' to set 'Content-Type' in the HTTP header.
// The media type can be a file extension, such as 'pdf' or 'png'
// in which case it gets converted to a proper MIME type,
// e.g. 'application/pdf' or 'image/png'
// Use the file extension value, e.g. "pdf"
func (ob *Context) Reply(data []byte, mediaType string) {
	//IDEA: it would be good if Reply() accepted strings, io.Reader, etc.
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	mediaType = MediaType(mediaType)
	if mediaType != "" {
		ob.w.Header().Set("Content-Type", mediaType)
	}
	if mediaType == "" {
		zr.Error(zr.EInvalidArg, "^mediaType", ":^", mediaType)
	}
	if ContextDebugFunc != nil {
		const LE = " " + zr.LF           // line end
		defer contextDebugMutex.Unlock() // locked by NewContext()
		crc := fmt.Sprintf("%08X", crc32.ChecksumIEEE(data))
		var sdata string
		if len(data) > 0 &&
			mediaType != "application/javascript" &&
			mediaType != "application/x-font" &&
			mediaType != "image/png" &&
			mediaType != "image/svg+xml" &&
			mediaType != "image/x-icon" &&
			mediaType != "text/css" {
			if zr.DebugMode() {
				sdata = string(data)
			} else {
				sdata = zr.DebugString(data)
				if len(sdata) > 40 {
					sdata = sdata[:40]
				}
			}
			sdata = strings.Trim(sdata, zr.SPACES)
			sdata = strings.Repeat("-", 80) + ">" + LE +
				sdata + LE + "<" + strings.Repeat("-", 80) + LE
		}
		contextDebugPrint(
			"REPLY:", ob.id,
			" sid:", ob.Session.ID()[:8],
			" type:", mediaType,
			" crc:", crc,
			" len:", len(data),
			LE,
			sdata,
			LE,
		)
	}
	ob.w.Write(data)
} //                                                                       Reply

// -----------------------------------------------------------------------------
// # Debug Helper Method

// DebugString __
func (ob *Context) DebugString() string {
	postdata := string(ob.PostData())
	return fmt.Sprint(
		"BaseReferer(): ", ob.BaseReferer(), "\n",
		"Method(): ", ob.Method(), "\n",
		"HREF(): ", ob.HREF(), "\n",
		"PostData(): ", postdata, "\n",
		"Referer(): ", ob.Referer(), "\n",
	)
} //                                                                 DebugString

// -----------------------------------------------------------------------------
// # Support (File Scope)

// readPostData reads the content of a POST HTTP request into a byte array.
// This function should only be called once on a Request.
// Subsequent calls on the same Request will return an empty array.
func readPostData(req *http.Request) []byte {
	var ret []byte
	{
		buf := bytes.NewBuffer(make([]byte, 0, 1024))
		req.Write(buf)
		ret = buf.Bytes()
	}
	pos := bytes.Index(ret, []byte("\r\n\r\n")) // skip HTTP headers
	if pos != -1 {
		ret = ret[pos+4:]
	}
	return ret
} //                                                                readPostData

//end
