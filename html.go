// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-07-11 10:45:59 A46878                               zr-web/[html.go]
// -----------------------------------------------------------------------------

package web

// # Global Settings
//   UseNthChild() bool
//   SetUseNthChild(val bool)
//
// # Functions
// # Top-Level Container Elements
// # Container Elements
// # HTML Attributes
// # Helper Tags (non-standard tags that simplify markup)
// # Non-Container Elements
// # General Wrappers
//
// # Functions
//   SetClass(add bool, input string, classes ...string) string
//
// # Top-Level Container Elements
//   HTML(content ...interface{}) []byte
//   Head(content ...interface{}) *Buffer
//   Body(content ...interface{}) *Buffer
//
// # Container Elements
//   A(href string, content ...interface{}) *Buffer
//   Article(content ...interface{}) *Buffer
//   Div(content ...interface{}) *Buffer
//   Form(content ...interface{}) *Buffer
//   H1(content ...interface{}) *Buffer
//   H2(content ...interface{}) *Buffer
//   H3(content ...interface{}) *Buffer
//   H4(content ...interface{}) *Buffer
//   H5(content ...interface{}) *Buffer
//   H6(content ...interface{}) *Buffer
//   Header(content ...interface{}) *Buffer
//   IFrame(content ...interface{}) *Buffer
//   Img(content ...interface{}) *Buffer
//   Label(content ...interface{}) *Buffer
//   Li(content ...interface{}) *Buffer
//   Nav(content ...interface{}) *Buffer
//   P(content ...interface{}) *Buffer
//   Span(content ...interface{}) *Buffer
//   Title(content ...interface{}) *Buffer
//   Ul(content ...interface{}) *Buffer
//
// # HTML Attributes
//   Attr(name, val string) Attribute
//   Charset(locale string) Attribute
//   Class(classList ...string) Attribute
//   Content(locale string) Attribute
//   HREF(href string) Attribute
//   ID(locale string) Attribute
//   Lang(locale string) Attribute
//   Name(locale string) Attribute
//   OnClick(jsCall string) Attribute
//   OnLoad(jsCall string) Attribute
//   Type(locale string) Attribute
//
// # Helper Tags (non-standard tags that simplify markup)
//   COLUMNS(cols []string, class string, useNthChild bool) *Buffer
//   CSS(styles ...string) *Buffer
//   JOIN(content... *Buffer) *Buffer
//   JS(scripts ...string) *Buffer
//   NAV(href string, content ...interface{}) *Buffer
//   TEXT(texts ...string) *Buffer
//
// # Non-Container Elements
//   Br(attributes ...Attribute) *Buffer
//   Hr(attributes ...Attribute) *Buffer
//   Input(attributes ...Attribute) *Buffer
//   Meta(attributes ...Attribute) *Buffer
//
// # General Wrappers
//   Comment(s string) *Buffer
//   Container(elementName string, content ...interface{}) *Buffer
//   Element(elementName string, attributes ...Attribute) *Buffer

import (
	"bytes"
	"fmt"
	"reflect"
	str "strings"

	"github.com/balacode/zr"
)

// oldBrowsers constant specifies if older browsers should
// be supported by the framework. False by default.
const oldBrowsers = false

// useNthChild constant specifies if the framework should use
// 'nth-child' CSS pseudo-selectors when generating content.
// Using 'nth-child' makes HTML output shorter than setting
// a column-number class for each output column.
// True by default.
var useNthChild = true

// -----------------------------------------------------------------------------
// # Global Settings

// UseNthChild __
func UseNthChild() bool {
	return useNthChild
} //                                                                 UseNthChild

// SetUseNthChild __
func SetUseNthChild(val bool) {
	useNthChild = val
} //                                                              SetUseNthChild

// -----------------------------------------------------------------------------
// # Functions

// SetClass appends or removes the specified class(es) to the given string.
// 'add' specifies if the class should be added (or removed if false).
//  'input' is the existing class string.
//  'classes' is a list of class names to add or remove.
//  Returns the input string with the required classes added or removed.
func SetClass(add bool, input string, classes ...string) string {
	var classList []string
	if input != "" {
		classList = str.Split(input, " ")
	}
	for _, class := range classes {
		if add {
			var found bool
			for _, s := range classList {
				if s == class {
					found = true
					break
				}
			}
			if !found {
				classList = append(classList, class)
			}
			continue
		}
		// remove:
		for i := 0; i < len(classList); i++ {
			if classList[i] == class {
				classList = classList[:i+copy(classList[i:], classList[i+1:])]
			}
		}
	}
	return str.Join(classList, " ")
} //                                                                    SetClass

// -----------------------------------------------------------------------------
// # Top-Level Container Elements

// HTML tag is the root of an HTML document and the top-level
// container for all other HTML tags.
// Attributes: manifest, xmlns
func HTML(content ...interface{}) []byte {
	var retBuf = NewBuffer(4096)
	retBuf.WriteString("<!DOCTYPE html>" + LB)
	retBuf.Write(Container("html", content...))
	return retBuf.Bytes()
} //                                                                        HTML

// Head element is the container for a page's head elements,
// including the title, scripts and metadata.
// Allowed child elements: <base> <link> <meta> <noscript> <script> <style>
// <title> (required)
func Head(content ...interface{}) *Buffer {
	return Container("head", content...)
} //                                                                        Head

// Body tag: contains the visible content of a page,
// including text, images, hyperlinks, etc.
// Attributes: alink, background, bgcolor, link, text, vlink
func Body(content ...interface{}) *Buffer {
	return Container("body", content...)
} //                                                                        Body

// -----------------------------------------------------------------------------
// # Container Elements

// A tag specifies a hyperlink, which links other web
// pages and locations in the current document.
// Attributes: charset coords download href hreflang
// media name rel rev shape target type
func A(href string, content ...interface{}) *Buffer {
	//TODO: prevent multiple href attributes
	content = append(content, HREF(href))
	return Container("a", content...)
} //                                                                           A

// Article defines self-contained content.
// Attributes: Global, Event
func Article(content ...interface{}) *Buffer {
	if oldBrowsers {
		return Container("div", content...)
	}
	return Container("article", content...)
} //                                                                     Article

// Div tag defines an arbitrary division in the document.
// Attributes: align (left, right, center, justify)
func Div(content ...interface{}) *Buffer {
	return Container("div", content...)
} //                                                                         Div

// Form tag.
func Form(content ...interface{}) *Buffer {
	return Container("form", content...)
} //                                                                        Form

// H1 tag specifies headings (level 1)
// Attributes: align (left, center, right, justify)
func H1(content ...interface{}) *Buffer {
	return Container("h1", content...)
} //                                                                          H1

// H2 tag specifies headings (level 2)
// Attributes: align (left, center, right, justify)
func H2(content ...interface{}) *Buffer {
	return Container("h2", content...)
} //                                                                          H2

// H3 tag specifies headings (level 3)
// Attributes: align (left, center, right, justify)
func H3(content ...interface{}) *Buffer {
	return Container("h3", content...)
} //                                                                          H3

// H4 tag specifies headings (level 4)
// Attributes: align (left, center, right, justify)
func H4(content ...interface{}) *Buffer {
	return Container("h4", content...)
} //                                                                          H4

// H5 tag specifies headings (level 5)
// Attributes: align (left, center, right, justify)
func H5(content ...interface{}) *Buffer {
	return Container("h5", content...)
} //                                                                          H5

// H6 tag specifies headings (level 6)
// Attributes: align (left, center, right, justify)
func H6(content ...interface{}) *Buffer {
	return Container("h6", content...)
} //                                                                          H6

// Header tag contains the headings of a section or navigational links.
func Header(content ...interface{}) *Buffer {
	if oldBrowsers {
		Div(
			Class("header"),
			Container("div", content...),
		)
	}
	return Div(
		Class("header"),
		Container("header", content...),
	)
} //                                                                      Header

// IFrame represents an <iframe> tag, i.e. inline frame.
func IFrame(content ...interface{}) *Buffer {
	return Container("iframe", content...)
} //                                                                      IFrame

// Img inserts an image element.
// For example Img("folder/filename.png") will become
// <img src="folder/filename.png"> in the output HTML
func Img(content ...interface{}) *Buffer {
	for i, it := range content {
		if s, ok := it.(string); ok {
			content[i] = Attribute{Name: "src", Value: s}
		}
	}
	return Container("img", content...)
} //                                                                         Img

// Label tag represents a label element.
func Label(content ...interface{}) *Buffer {
	return Container("label", content...)
} //                                                                       Label

// Li tag defines a list item in unordered (ul), ordered (ol)
// or menu lists (menu). Attributes: global attributes and value.
func Li(content ...interface{}) *Buffer {
	return Container("li", content...)
} //                                                                          Li

// Nav tag defines a section with navigation links.
func Nav(content ...interface{}) *Buffer {
	if oldBrowsers {
		return Container("div", content...)
	}
	return Container("nav", content...)
} //                                                                         Nav

// P tag defines a paragraph.
// Attribute: align, global attributes
func P(content ...interface{}) *Buffer {
	return Container("p", content...)
} //                                                                           P

// Span tag groups inline elements.
func Span(content ...interface{}) *Buffer {
	return Container("span", content...)
} //                                                                        Span

// Title tag specifies the title or name of the document.
// It is required to be placed within the <head> element.
func Title(content ...interface{}) *Buffer {
	return Container("title", content...)
} //                                                                       Title

// Ul defines an unordered list. (A list in which the ordering of items
// is not important). The list can contain zero or more <li> elements.
func Ul(content ...interface{}) *Buffer {
	return Container("ul", content...)
} //                                                                          Ul

// -----------------------------------------------------------------------------
// # HTML Attributes

// Attribute holds the name and value of a single HTML attribute.
type Attribute struct {
	Name  string
	Value string
} //                                                                   Attribute

// Attr specifies any element's attribute.
func Attr(name, val string) Attribute {
	return Attribute{Name: name, Value: val}
} //                                                                        Attr

// Charset attribute applies to <meta> tags.
func Charset(locale string) Attribute {
	return Attribute{Name: "charset", Value: locale}
} //                                                                     Charset

// Class represents the 'class' attribute. You can specify multiple class
// strings in which case they will be delimited by a space. E.g.
// Class("currency", "sum") will become class="currency sum"
func Class(classList ...string) Attribute {
	var class string
	for _, s := range classList {
		if class != "" {
			class += " "
		}
		class += s
	}
	return Attribute{Name: "class", Value: class}
} //                                                                       Class

// Content attribute applies to <meta> tags.
func Content(locale string) Attribute {
	return Attribute{Name: "content", Value: locale}
} //                                                                     Content

// HREF attribute applies to <a> tags.
func HREF(href string) Attribute {
	return Attribute{Name: "href", Value: href}
} //                                                                        HREF

// ID attribute applies to various tags.
func ID(locale string) Attribute {
	return Attribute{Name: "id", Value: locale}
} //                                                                          ID

// Lang attribute applies to <html> tags.
func Lang(locale string) Attribute {
	return Attribute{Name: "lang", Value: locale}
} //                                                                        Lang

// Name attribute applies to various tags.
func Name(locale string) Attribute {
	return Attribute{Name: "name", Value: locale}
} //                                                                        Name

// OnClick attribute applies to various tags.
func OnClick(jsCall string) Attribute {
	if jsCall == "" {
		return Attribute{}
	}
	return Attribute{Name: "onclick", Value: jsCall}
} //                                                                     OnClick

// OnLoad attribute applies to various tags.
func OnLoad(jsCall string) Attribute {
	if jsCall == "" {
		return Attribute{}
	}
	return Attribute{Name: "onload", Value: jsCall}
} //                                                                      OnLoad

// Type attribute.
func Type(locale string) Attribute {
	return Attribute{Name: "type", Value: locale}
} //                                                                        Type

// -----------------------------------------------------------------------------
// # Helper Tags (non-standard tags that simplify markup)

// COLUMNS creates a group of <p> tags with class c1, c2, c3, etc.
// for every string passed in 'columns'. Used to create tabular listings
// using CSS and <p> tags, without the need to use HTML tables.
func COLUMNS(cols []string, class string, useNthChild bool) *Buffer {
	var retBuf Buffer
	var ws = retBuf.WriteString
	ws("<div>" + LB)
	for i, col := range cols {
		ws("<p")
		var hasClass = str.Contains(col, "class::") || class != ""
		if !useNthChild || hasClass {
			ws(` class="`)
			if class != "" {
				ws(class)
			}
			if !useNthChild {
				if class != "" {
					ws(" ")
				}
				ws(fmt.Sprintf("c c%d", i+1))
			}
			if hasClass {
				var part = zr.GetPart(col, "class::", ";")
				col = str.Replace(col, "class::"+part+";", "", -1)
			}
			ws(`"`)
		}
		ws(">", col, "</p>"+LB)
	}
	ws("</div>" + LB)
	return &retBuf
} //                                                                     COLUMNS

// CSS links or embeds one or more CSS (Cascading Style Sheet)
// files or scripts.
//
// To link to a CSS file, just specify the filename in 'styles',
// for example: CSS("global_style.css")
//
// To embed a style locally, specify the style in styles, for example,
// CSS("body { font: normal 11pt Helvetica }")
//
// Should be placed within the <head> element under <html>.
func CSS(styles ...string) *Buffer {
	var retBuf Buffer
	var ws = retBuf.WriteString
	for _, style := range styles {
		style = str.Trim(style, zr.SPACES)
		if style == "" {
			continue
		}
		if zr.ContainsI(style, ".css") {
			ws(`<link rel="stylesheet" type="text/css" href="` +
				style + `">` + LB)
			continue
		}
		ws(`<style type="text/css">` + LB + style + LB + "</style>" + LB)
	}
	return &retBuf
} //                                                                         CSS

// JOIN concatenates the content of multiple buffers into a single buffer.
// It can be useful to merge together several HTML elements,
// with one element listed after the other.
func JOIN(content ...*Buffer) *Buffer {
	var retBuf = NewBuffer(64)
	for _, buf := range content {
		retBuf.Write(buf)
	}
	return &retBuf
} //                                                                        JOIN

// JS links JavaScript (.js) script files or embeds JS code snippets.
func JS(scripts ...string) *Buffer {
	var retBuf Buffer
	var ws = retBuf.WriteString
	for _, js := range scripts {
		js = str.Trim(js, zr.SPACES)
		if js == "" {
			continue
		}
		if zr.ContainsI(js, ".js") {
			ws(`<script type="text/javascript" src="` +
				js + `"></script>` + LB)
			continue
		}
		ws(`<script type="text/javascript">` + js + "</script>" + LB)
	}
	return &retBuf
} //                                                                          JS

// NAV helper tag specifies a hyperlink, which links other web pages and
// locations in the current document. It is similar to the 'A' tag,
// but uses zr.go() in JS to save the current page reference.
// Attributes: charset coords download href hreflang
//             media name rel rev shape target type
func NAV(href string, content ...interface{}) *Buffer {
	//TODO: prevent multiple href attributes
	var isFuncCall = str.Contains(href, "(") && str.Contains(href, ")")
	if !isFuncCall {
		href = fmt.Sprintf("zr.go('%s')", href)
	}
	content = append(content, Attr("onclick", href))
	return A("#", content...)
} //                                                                         NAV

// TEXT helper tag is a non-standard tag that helps
// inject literal strings into HTML content.
func TEXT(texts ...string) *Buffer {
	var retBuf Buffer
	var ws = retBuf.WriteString
	for _, s := range texts {
		ws(s)
	}
	return &retBuf
} //                                                                        TEXT

// -----------------------------------------------------------------------------
// # Non-Container Elements

// Br tag inserts a line break.
// This tag has no closing tag and is not a container.
// Attributes: Global, Event
func Br(attributes ...Attribute) *Buffer {
	return Element("br", attributes...)
} //                                                                          Br

// Hr tag inserts a thematic break (horizontal rule pre-HTML5).
// This tag has no closing tag and is not a container.
// Attributes: Global, Event
func Hr(attributes ...Attribute) *Buffer {
	return Element("hr", attributes...)
} //                                                                          Hr

// Input tag represents an element for user input.
// This tag has no closing tag and is not a container.
func Input(attributes ...Attribute) *Buffer {
	return Element("input", attributes...)
} //                                                                       Input

// Meta tag: Provides information about the HTML document.
// Attributes: charset, content, http-equiv, name, scheme
func Meta(attributes ...Attribute) *Buffer {
	return Element("meta", attributes...)
} //                                                                        Meta

// -----------------------------------------------------------------------------
// # General Wrappers

// Comment composes an HTML comment.
func Comment(s string) *Buffer {
	//TODO: change 's string' to 'args ...interface{}' and use fmt.Sprint()
	var retBuf = NewBuffer(64)
	var ws = retBuf.WriteString
	ws("<!--")
	ws(s)
	ws("-->" + LB)
	return &retBuf
} //                                                                     Comment

// Container composes an arbitrary HTML container tag.
func Container(elementName string, content ...interface{}) *Buffer {
	var retBuf = NewBuffer(64)
	var w = retBuf.Write
	var wb = retBuf.WriteBytes
	var ws = retBuf.WriteString
	// write the opening tag and its attributes
	ws("<" + elementName)
	for _, val := range content {
		switch val := val.(type) {
		case Attribute:
			if val.Name != "" && val.Value != "" {
				ws(fmt.Sprintf(` %s="%s"`, val.Name, val.Value))
			}
		}
	}
	ws(">")
	if zr.StrOneOf(elementName,
		"article", "body", "div", "head", "header", "html", "nav", "ul") {
		ws(LB)
	}
	// write the inner content (usually the byte buffers of child tags)
	for i, val := range content {
		switch val := val.(type) {
		case Attribute:
			// do nothing: already handled above
		case []byte:
			wb(val)
		case *[]byte:
			if val != nil {
				wb(*val)
			}
		// web.Buffer
		case Buffer:
			w(&val)
		case *Buffer:
			if val != nil {
				w(val)
			}
		case []Buffer:
			for _, val := range val {
				w(&val)
			}
		case []*Buffer:
			for _, val := range val {
				w(val)
			}
		// bytes.Buffer
		case bytes.Buffer:
			wb(val.Bytes())
		case *bytes.Buffer:
			if val != nil {
				wb(val.Bytes())
			}
		// numbers
		case float32, float64:
			ws(fmt.Sprintf("%f", val))
		case int, int8, int16, int32, int64,
			uint, uint8, uint16, uint32, uint64:
			ws(fmt.Sprintf("%d", val))
		// strings
		case string:
			ws(val)
		case []string:
			for _, s := range val {
				ws(s)
			}
		case fmt.Stringer:
			ws(val.String())
		default:
			zr.Error("Content item", i, "of type",
				reflect.TypeOf(val), "not handled")
		}
	}
	// write closing tag
	ws("</" + elementName + ">")
	if elementName != "a" {
		ws(LB)
	}
	return &retBuf
} //                                                                   Container

// Element composes a HTML tag with optional attributes but no child tags.
func Element(elementName string, attributes ...Attribute) *Buffer {
	var retBuf = NewBuffer(64)
	var ws = retBuf.WriteString
	ws("<" + elementName)
	for _, attr := range attributes {
		ws(fmt.Sprintf(` %s="%s"`, attr.Name, attr.Value))
	}
	ws(">" + LB)
	return &retBuf
} //                                                                     Element

//end
