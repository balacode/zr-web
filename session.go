// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2019-05-17 16:30:46 A9DC96                            zr-web/[session.go]
// -----------------------------------------------------------------------------

package web

//   (ob *Session) ID() string
//   (ob *Session) GetSetting(name string) string
//   (ob *Session) SetSetting(name string, value interface{})

import (
	"sync"

	"github.com/balacode/zr"
)

// Session __
type Session struct {
	id    string
	m     map[string]string
	mutex sync.Mutex
} //                                                                     Session

// ID __
func (ob *Session) ID() string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return "" // error val
	}
	var ret string
	ob.mutex.Lock()
	ret = ob.id
	ob.mutex.Unlock()
	return ret
} //                                                                          ID

// GetSetting __
func (ob *Session) GetSetting(name string) string {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return "" // error val
	}
	var ret string
	ob.mutex.Lock()
	ret = ob.m[name]
	ob.mutex.Unlock()
	return ret
} //                                                                  GetSetting

// SetSetting __
func (ob *Session) SetSetting(name string, value interface{}) {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return
	}
	ob.mutex.Lock()
	ob.m[name] = zr.String(value)
	ob.mutex.Unlock()
} //                                                                  SetSetting

//end
