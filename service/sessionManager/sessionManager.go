package sessionmanager

import (
	"context"
	"errors"
	"time"

	"github.com/Zhima-Mochi/go-user-service/external"
	"github.com/Zhima-Mochi/go-user-service/service/sessionManager/session"
	"github.com/google/uuid"
)

var (
	timeNow = time.Now

	defaultSessionManager = &sessionManager{
		Name:   "session_id",
		MaxAge: 3600,
	}
)

type sessionManager struct {
	Name   string
	MaxAge int
	cache  external.Cache
}

// generateSessionID generates a new session id.
func generateSessionID() string {
	return uuid.New().String()
}

// NewSessionManager creates a new session manager.
func NewSessionManager(cache external.Cache, options ...SessionManagerOption) *sessionManager {
	sm := defaultSessionManager

	sm.cache = cache

	for _, option := range options {
		option(sm)
	}

	return sm
}

// NewSession creates a new session.
func (sm *sessionManager) NewSession(ctx context.Context, data map[string]interface{}) (*session.Session, error) {
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
func (sm *sessionManager) GetSession(ctx context.Context, id string) (*session.Session, error) {
	// get session data from cache
	s, err := sm.cache.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.(*session.Session), nil
}

// UpdateSession updates session by id.
func (sm *sessionManager) UpdateSession(ctx context.Context, id string, s *session.Session) error {
	if id != s.ID {
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
