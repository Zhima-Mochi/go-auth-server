package cookiemanager

import (
	"net/http"
	"time"
)

type cookieManager struct {
	cookieName   string
	cookieMaxAge int
}

// NewCookieManager creates new cookie manager.
func NewCookieManager(cookieName string, cookieMaxAge int) *cookieManager {
	return &cookieManager{
		cookieName:   cookieName,
		cookieMaxAge: cookieMaxAge,
	}
}

// SetCookie sets cookie.
func (cm *cookieManager) SetCookie(w http.ResponseWriter, value string) {
	cookie := &http.Cookie{
		Name:     cm.cookieName,
		Value:    value,
		Expires:  time.Now().Add(time.Duration(cm.cookieMaxAge) * time.Second),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}

// GetCookie gets cookie.
func (cm *cookieManager) GetCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie(cm.cookieName)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// DeleteCookie deletes cookie.
func (cm *cookieManager) DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     cm.cookieName,
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
}
