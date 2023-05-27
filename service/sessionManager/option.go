package sessionmanager

import (
	"github.com/Zhima-Mochi/go-authentication-service/external"
)

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

func WithCache(cache external.Cache) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.cache = cache
	}
}

func WithEncryptor(encryptor external.Encryptor) SessionManagerOption {
	return func(sm *sessionManager) {
		sm.encryptor = encryptor
	}
}
