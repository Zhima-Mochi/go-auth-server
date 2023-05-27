package session

import "time"

type Session interface {
	GetID() string
	GetExpires() time.Time
	GetData() map[string]interface{}
}
