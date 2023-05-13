package sessionmanager

import (
	"github.com/Zhima-Mochi/go-user-service/internal/sessionManager/session"
)

var (
	defaultSessionManager = &sessionManager{
		Name:     "session_id",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	}
)

type SessionManager interface {
	NewSession() (session.Session, error)
	GetSession(id string) (session.Session, error)
	UpdateSession(session session.Session) error
	DeleteSession(session session.Session) error
}

type SessionManagerOption func(*sessionManager)

func WithName(name string) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.Name = name
	}
}

func WithMaxAge(maxAge int) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.MaxAge = maxAge
	}
}

func WithSecure(secure bool) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.Secure = secure
	}
}

func WithHttpOnly(httpOnly bool) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.HttpOnly = httpOnly
	}
}

func WithSecretKey(secretKey []byte) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.SecretKey = secretKey
	}
}

func NewSessionManager(options ...SessionManagerOption) *sessionManager {
	sm := defaultSessionManager

	for _, option := range options {
		option(sm)
	}

	return sm
}
