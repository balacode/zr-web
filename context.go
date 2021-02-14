// -----------------------------------------------------------------------------
// ZR Library - Web Package                                  zr-web/[context.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package web

//	Context wraps an HTTP request and reply. Its methods such as HREF(),
//	Method() and PostData() provide request details. The Reply() method
//	is used to send a reply. The most important use of Context is that
//	it connects request to sessions. It checks the cookie sent with the
//	request and starts a new session or continues an existing session.
//
//	import (
//		"net/http"
//		"log"
//
//		web "github.com/balacode/zr-web"
//	)
//
//	func main() {
//		http.HandleFunc("/", mainServe)
//		log.Fatal(http.ListenAndServe("localhost:888", nil))
//	}
//
//	func mainServe(w http.ResponseWriter, req *http.Request) {
//		context := web.NewContext(w, req, &sessions)
//		context.Reply("<html><h1>Hello World</hq></html>", "html")
//	}

//  Context struct
//
// # Constructor
//    NewContext(w http.ResponseWriter, req *http.Request, sess *Sessions,
//        ) Context
//
// # Request Properties (ctx *Context)
//   BaseReferer() string
//   Method() string
//   HREF() string
//   PostData() []byte
//   Referer() string
//   URI() string
//
// # Methods (ctx *Context)
//   Redirect(url string)
//   Reply(data []byte, mediaType string)
//   ResetPostData()
//
// # Debug Helper Method
//   (ctx *Context) DebugString() string {

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

// nextContextID _ _
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
			postdata = strings.TrimSpace(postdata)
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
// # Request Properties (ctx *Context)

// BaseReferer property returns the base referer path of
// the current request. I.e. a path with '/', '\', '#'
// and numbers stripped from the end.
func (ctx *Context) BaseReferer() string {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := strings.TrimRight(ctx.Referer(), " \a\b\f\n\r\t\v"+`/\#-0123456789`)
	return ret
} //                                                                 BaseReferer

// Method returns the request's HTTP method ('GET', 'POST', 'PUT', etc)
func (ctx *Context) Method() string {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	return strings.ToUpper(ctx.req.Method)
} //                                                                      Method

// HREF property returns the URL path of the current request.
// This property does not return the query parameters.
func (ctx *Context) HREF() string {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := ctx.req.URL.Path
	if strings.HasPrefix(ret, "/") {
		// preserve existing trimming of '/' from the start
		// of the path, but this may need to be changed!
		ret = strings.TrimLeft(ret, `/`)
	}
	ret = FormatURL(ret)
	return ret
} //                                                                        HREF

// PostData property returns the POSTDATA of the current request.
func (ctx *Context) PostData() []byte {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return []byte{}
	}
	if len(ctx.postData) > 0 {
		return ctx.postData
	}
	ctx.postData = readPostData(ctx.req)
	return ctx.postData
} //                                                                    PostData

// Referer property returns the referer path of the current request.
func (ctx *Context) Referer() string {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := strings.Trim(ctx.req.Referer(), `#/\ `)
	if strings.Contains(ret, "\\") {
		ret = strings.ReplaceAll(ret, "\\", "/")
	}
	return ret
} //                                                                     Referer

// URI property returns the full URI path of the current request.
// This includes the HREF() part and any query parameters.
func (ctx *Context) URI() string {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return ""
	}
	ret := strings.Trim(ctx.req.RequestURI, `#/\ `)
	if strings.Contains(ret, "\\") {
		ret = strings.ReplaceAll(ret, "\\", "/")
	}
	return ret
} //                                                                        URI

// -----------------------------------------------------------------------------
// # Methods (ctx *Context)

// Redirect redirects the client to another url using
// HTTP redirect code 302 (temporary redirect).
func (ctx *Context) Redirect(url string) {
	http.Redirect(ctx.w, ctx.req, url, http.StatusFound)
} // 																	Redirect

// Reply method sends the reply to a request.
// Specify 'mediaType' to set 'Content-Type' in the HTTP header.
// The media type can be a file extension, such as 'pdf' or 'png'
// in which case it gets converted to a proper MIME type,
// e.g. 'application/pdf' or 'image/png'
// Use the file extension value, e.g. "pdf"
func (ctx *Context) Reply(data []byte, mediaType string) {
	//IDEA: it would be good if Reply() accepted strings, io.Reader, etc.
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	mediaType = MediaType(mediaType)
	if mediaType != "" {
		ctx.w.Header().Set("Content-Type", mediaType)
	}
	if mediaType == "" {
		zr.Error(zr.EInvalidArg, "^mediaType", ":^", mediaType)
	}
	if ContextDebugFunc != nil {
		const LE = " \n"                 // line end
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
			sdata = strings.TrimSpace(sdata)
			sdata = strings.Repeat("-", 80) + ">" + LE +
				sdata + LE + "<" + strings.Repeat("-", 80) + LE
		}
		contextDebugPrint(
			"REPLY:", ctx.id,
			" sid:", ctx.Session.ID()[:8],
			" type:", mediaType,
			" crc:", crc,
			" len:", len(data),
			LE,
			sdata,
			LE,
		)
	}
	ctx.w.Write(data)
} //                                                                       Reply

// ResetPostData _ _
func (ctx *Context) ResetPostData() {
	if ctx == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	ctx.postData = []byte{}
} //                                                               ResetPostData

// -----------------------------------------------------------------------------
// # Debug Helper Method

// DebugString _ _
func (ctx *Context) DebugString() string {
	postdata := string(ctx.PostData())
	return fmt.Sprint(
		"BaseReferer(): ", ctx.BaseReferer(), "\n",
		"Method(): ", ctx.Method(), "\n",
		"HREF(): ", ctx.HREF(), "\n",
		"PostData(): ", postdata, "\n",
		"Referer(): ", ctx.Referer(), "\n",
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

// end
