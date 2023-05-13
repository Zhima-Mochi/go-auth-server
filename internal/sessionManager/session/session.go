package session

import (
	"time"
)

type session struct {
	ID      string
	Expires time.Time
	Data    map[string]interface{}
}

func NewSession(sessionID string, expires time.Time, data map[string]interface{}) *session {
	return &session{
		ID:      sessionID,
		Expires: expires,
		Data:    data,
	}
}

func (s *session) GetID() string {
	return s.ID
}

func (s *session) GetExpires() time.Time {
	return s.Expires
}

func (s *session) GetData() map[string]interface{} {
	return s.Data
}
