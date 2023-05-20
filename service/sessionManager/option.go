package sessionmanager

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
