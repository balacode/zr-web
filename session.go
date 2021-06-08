// -----------------------------------------------------------------------------
// ZR Library - Web Package                                  zr-web/[session.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package web

//   (ob *Session) ID() string
//   (ob *Session) GetSetting(name string) string
//   (ob *Session) SetSetting(name string, value interface{})

import (
	"sync"

	"github.com/balacode/zr"
)

// Session _ _
type Session struct {
	id    string
	m     map[string]string
	mutex sync.Mutex
} //                                                                     Session

// ID _ _
func (ob *Session) ID() string {
	var ret string
	ob.mutex.Lock()
	ret = ob.id
	ob.mutex.Unlock()
	return ret
} //                                                                          ID

// GetSetting _ _
func (ob *Session) GetSetting(name string) string {
	var ret string
	ob.mutex.Lock()
	ret = ob.m[name]
	ob.mutex.Unlock()
	return ret
} //                                                                  GetSetting

// SetSetting _ _
func (ob *Session) SetSetting(name string, value interface{}) {
	ob.mutex.Lock()
	ob.m[name] = zr.String(value)
	ob.mutex.Unlock()
} //                                                                  SetSetting

// end
