package sessionmanager

import (
	"context"
	"errors"
	"time"

	"github.com/Zhima-Mochi/go-authentication-service/external"
	"github.com/Zhima-Mochi/go-authentication-service/service/sessionManager/session"
	"github.com/Zhima-Mochi/go-authentication-service/service/utility"
	"github.com/google/uuid"
)

var _ SessionManager = (*sessionManager)(nil)

var (
	timeNow = time.Now

	defaultSessionManager = &sessionManager{
		Name:      "session_id",
		MaxAge:    3600,
		encryptor: utility.NewEncryptor(),
		cache:     utility.NewCache(),
	}
)

type sessionManager struct {
	Name      string
	MaxAge    int
	encryptor external.Encryptor
	cache     external.Cache
}

// generateSessionID generates a new session id.
func generateSessionID() string {
	return uuid.New().String()
}

// NewSessionManager creates a new session manager.
func NewSessionManager(options ...SessionManagerOption) *sessionManager {
	sm := defaultSessionManager

	for _, option := range options {
		option(sm)
	}

	return sm
}

// CreateSession creates a new session.
func (sm *sessionManager) CreateSession(ctx context.Context, data map[string]interface{}) (session.Session, error) {
	// generate session id
	id := generateSessionID()

	// calculate session expiration time
	expiresAt := timeNow().Add(time.Duration(sm.MaxAge) * time.Second)

	// create new session
	s := session.NewSession(id, expiresAt, data)

	err := sm.cache.Set(ctx, id, s, time.Duration(sm.MaxAge)*time.Second)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// GetSession gets session by id.
func (sm *sessionManager) GetSession(ctx context.Context, id string) (session.Session, error) {
	// get session data from cache
	s, err := sm.cache.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.(session.Session), nil
}

// UpdateSession updates session by id.
func (sm *sessionManager) UpdateSession(ctx context.Context, id string, s session.Session) error {
	if id != s.GetID() {
		return errors.New("session id not match")
	}

	err := sm.cache.Set(ctx, id, s, time.Duration(sm.MaxAge)*time.Second)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSession deletes session by id.
func (sm *sessionManager) DeleteSession(ctx context.Context, id string) error {
	err := sm.cache.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
