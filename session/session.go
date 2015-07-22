package session

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
)

var providers = make(map[string]Provider)

// Register makes a session provide available by the provided name.
// If Register is called twice with the same name or if driver is nil,
// it panics.
func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	providers[name] = provider
}

func init() {
	fmt.Println("--- session package init ---")

}

type Manager struct {
	cookieName  string // private cookiename
	lock        sync.Mutex
	provider    Provider //
	maxlifetime int64
}

func (manager *Manager) sessionId() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	//	if _, err := rand.Read(b); err != nil {
	//		return ""
	//	}
	return base64.URLEncoding.EncodeToString(b)
}

//我们需要为每个来访用户分配或获取与他相关连的Session

func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	fmt.Println("Session Start!!!!!")
	//	manager.lock.Lock()
	//	defer manager.lock.Unlock()

	cookie, err := r.Cookie(manager.cookieName)
	fmt.Println("cookie", cookie)
	return
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		//客户端已经有cookie
		//从cookie里读出sid,到provider里读出Session对象
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return

}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

func NewManager(provideName, cookieName string, maxlifetime int64) (*Manager, error) {
	provider, ok := providers[provideName] //
	if !ok {
		return nil, fmt.Errorf("session: unknow provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}
