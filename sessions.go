// -----------------------------------------------------------------------------
// ZR Library - Web Package                                  zr-web/[session.go]
// (c) balarabe@protonmail.com                                      License: MIT
// -----------------------------------------------------------------------------

package web

import (
	"net/http"
	"strings"
	"sync"

	"github.com/balacode/zr"
)

// Sessions _ _
type Sessions struct {
	m     map[string]*Session
	mutex sync.Mutex
} //                                                                    Sessions

// GetByCookie _ _
func (ob *Sessions) GetByCookie(
	w http.ResponseWriter,
	req *http.Request,
) *Session {
	if ob == nil {
		zr.Error(zr.ENilReceiver)
		return nil
	}
	ob.mutex.Lock()
	defer ob.mutex.Unlock()
	//
	const CookieName = "app_session_id"
	// if session cookie already exists, use its session ID..
	var id string
	cookie, err := req.Cookie(CookieName)
	if err == nil {
		id = cookie.Value
	} else {
		// ..if not, create new session ID and save it in a cookie
		id = strings.Replace(zr.UUID(), "-", "", -1)
		http.SetCookie(w, &http.Cookie{Name: CookieName, Value: id})
	}
	// if session is already stored, return pointer to stored session
	ptr, exists := ob.m[id]
	if exists {
		return ptr
	}
	// if not, add a new Session to the map
	ses := Session{id: id, m: map[string]string{}}
	ptr = &ses
	if ob.m == nil {
		ob.m = make(map[string]*Session, 0)
	}
	ob.m[id] = ptr
	return ptr
} //                                                                 GetByCookie

//end
