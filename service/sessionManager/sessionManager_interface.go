package sessionmanager

import (
	"context"

	"github.com/Zhima-Mochi/go-authentication-service/service/sessionManager/session"
)

type SessionManager interface {
	CreateSession(ctx context.Context, data map[string]interface{}) (session.Session, error)
	GetSession(ctx context.Context, id string) (session.Session, error)
	UpdateSession(ctx context.Context, id string, session session.Session) error
	DeleteSession(ctx context.Context, id string) error
}
