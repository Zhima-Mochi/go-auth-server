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
