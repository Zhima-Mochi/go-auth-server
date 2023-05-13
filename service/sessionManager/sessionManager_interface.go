package sessionmanager

import (
	"context"

	"github.com/Zhima-Mochi/go-user-service/service/sessionManager/session"
)

type SessionManager interface {
	NewSession(ctx context.Context) (session.Session, error)
	GetSession(ctx context.Context, id string) (session.Session, error)
	UpdateSession(ctx context.Context, id string, session session.Session) error
	DeleteSession(ctx context.Context, id string) error
}
