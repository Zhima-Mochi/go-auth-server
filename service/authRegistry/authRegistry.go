package authregistry

import (
	"fmt"
)

type authRegistry struct {
	auth map[ProviderType]Auth
}

// Register registers an auth provider
func (r *authRegistry) Register(providerType ProviderType, auth Auth) {
	r.auth[providerType] = auth
}

// GetAuth returns an auth provider
func (r *authRegistry) GetAuth(providerType ProviderType) (Auth, error) {
	auth, ok := r.auth[providerType]
	if !ok {
		return nil, fmt.Errorf("provider %s is not registered", providerType)
	}
	return auth, nil
}

// NewAuthRegistry returns a new auth registry
func NewAuthRegistry() *authRegistry {
	return &authRegistry{
		auth: map[ProviderType]Auth{},
	}
}
