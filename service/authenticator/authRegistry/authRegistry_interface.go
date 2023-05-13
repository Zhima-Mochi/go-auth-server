package authregistry

import (
	oauth2providers "github.com/Zhima-Mochi/go-oauth2-providers"
)

type AuthRegistry interface {
	Register(providerType ProviderType, auth Auth)
	GetAuth(providerType ProviderType) (Auth, error)
}

type Auth oauth2providers.Auth

type ProviderType oauth2providers.ProviderType
