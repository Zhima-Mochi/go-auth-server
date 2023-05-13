package session

import (
	"time"
)

type Session struct {
	ID      string
	Expires time.Time
	Data    map[string]interface{}
}

func NewSession(sessionID string, expires time.Time, data map[string]interface{}) *Session {
	return &Session{
		ID:      sessionID,
		Expires: expires,
		Data:    data,
	}
}

func (s *Session) GetID() string {
	return s.ID
}

func (s *Session) GetExpires() time.Time {
	return s.Expires
}

func (s *Session) GetData() map[string]interface{} {
	return s.Data
}
