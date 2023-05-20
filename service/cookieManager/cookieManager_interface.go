package cookiemanager

import (
	"net/http"
)

type CookieManager interface {
	SetCookie(w http.ResponseWriter, value string)
	GetCookie(r *http.Request) (string, error)
	DeleteCookie(w http.ResponseWriter)
}
