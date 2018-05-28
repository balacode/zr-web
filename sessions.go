// -----------------------------------------------------------------------------
// (c) balarabe@protonmail.com                                      License: MIT
// :v: 2018-05-24 13:36:58 5EFA16                            [zr-web/session.go]
// -----------------------------------------------------------------------------

package web

import (
	"net/http"
	str "strings"
	"sync"

	"github.com/balacode/zr"
)

// Sessions __
type Sessions struct {
	m     map[string]*Session
	mutex sync.Mutex
} //                                                                    Sessions

// GetByCookie __
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
	var cookie, err = req.Cookie(CookieName)
	if err == nil {
		id = cookie.Value
	} else {
		// ..if not, create new session ID and save it in a cookie
		id = str.Replace(zr.UUID(), "-", "", -1)
		http.SetCookie(w, &http.Cookie{Name: CookieName, Value: id})
	}
	// if session is already stored, return pointer to stored session
	var ptr, exists = ob.m[id]
	if exists {
		return ptr
	}
	// if not, add a new Session to the map
	var ses = Session{id: id, m: map[string]string{}}
	ptr = &ses
	if ob.m == nil {
		ob.m = make(map[string]*Session, 0)
	}
	ob.m[id] = ptr
	return ptr
} //                                                                 GetByCookie

//end