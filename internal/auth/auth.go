package auth

import (
	"fmt"
	"log"
	"os"

	oauth2providers "github.com/Zhima-Mochi/go-oauth2-providers"
	"github.com/Zhima-Mochi/go-user-service/internal/auth/config"
)

var (
	// registry is the registry of auth providers.
	registry = &authRegistry{}

	registerOAuth2Provider = func(envVar string, providerType oauth2providers.ProviderType) {
		enable := os.Getenv(envVar)
		if enable == "" {
			return
		}
		cfg, err := config.GetOAuth2ProviderConfig(providerType)
		if err != nil {
			log.Fatalf("failed to get %s config: %v", providerType, err)
		}
		providerOptions := []oauth2providers.ProviderOption{}
		providerOptions = append(providerOptions, oauth2providers.WithProviderClientID(cfg.ClientID))
		providerOptions = append(providerOptions, oauth2providers.WithProviderClientSecret(cfg.ClientSecret))
		providerOptions = append(providerOptions, oauth2providers.WithProviderRedirectURL(cfg.RedirectURL))
		providerOptions = append(providerOptions, oauth2providers.WithProviderScopes(cfg.Scopes))
		providerConfig, err := oauth2providers.NewProviderConfig(providerOptions...)
		if err != nil {
			log.Fatalf("failed to create %s provider config: %v", providerType, err)
		}
		oauth2Provider, err := oauth2providers.NewOAuth2Provider(providerType, providerConfig)
		if err != nil {
			log.Fatalf("failed to create %s oauth2 provider: %v", providerType, err)
		}
		oauth2Auth := oauth2providers.NewOAuth2Auth(oauth2Provider, nil)

		registry.Register(providerType, oauth2Auth)
	}
)

func init() {
	registerOAuth2Provider("LINE_AUTH_ENABLE", oauth2providers.LineOAuth2ProviderType)
	registerOAuth2Provider("GOOGLE_AUTH_ENABLE", oauth2providers.GoogleOAuth2ProviderType)
}

type authRegistry struct {
	auth map[oauth2providers.ProviderType]oauth2providers.Auth
}

func (r *authRegistry) Register(providerType oauth2providers.ProviderType, auth oauth2providers.Auth) {
	r.auth[providerType] = auth
}

func (r *authRegistry) GetAuth(providerType oauth2providers.ProviderType) (oauth2providers.Auth, error) {
	auth, ok := r.auth[providerType]
	if !ok {
		return nil, fmt.Errorf("provider %s is not registered", providerType)
	}
	return auth, nil
}

func GetAuthRegistry() *authRegistry {
	return registry
}
