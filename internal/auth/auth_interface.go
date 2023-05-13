package auth

import oauth2providers "github.com/Zhima-Mochi/go-oauth2-providers"

type AuthRegistry interface {
	Register(providerType oauth2providers.ProviderType, auth oauth2providers.Auth)
	GetAuth(providerType oauth2providers.ProviderType) (oauth2providers.Auth, error)
}
