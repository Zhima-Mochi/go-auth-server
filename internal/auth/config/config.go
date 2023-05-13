package config

import (
	"fmt"
	"io/ioutil"
	"os"

	oauth2providers "github.com/Zhima-Mochi/go-oauth2-providers"
	"gopkg.in/yaml.v2"
)

type ProviderConfig struct {
	// ClientID is the OAuth client ID.
	ClientID string `yaml:"client_id"`
	// ClientSecret is the OAuth client secret.
	ClientSecret string `yaml:"client_secret"`
	// Scopes is a list of OAuth scopes.
	Scopes []string `yaml:"scopes"`
	// RedirectURL is the OAuth redirect URL.
	RedirectURL string `yaml:"redirect_url"`
}

func GetOAuth2ProviderConfig(providerType oauth2providers.ProviderType) (*ProviderConfig, error) {
	switch providerType {
	case oauth2providers.FacebookOAuth2ProviderType:
		return getFacebookConfig()
	case oauth2providers.GoogleOAuth2ProviderType:
		return getGoogleConfig()
	case oauth2providers.LineOAuth2ProviderType:
		return getLineConfig()
	default:
		return nil, fmt.Errorf("provider %s is not supported", providerType)
	}
}

func getLineConfig() (*ProviderConfig, error) {
	// read yaml file
	filePath := os.Getenv("LINE_CONFIG_FILE")
	if filePath == "" {
		return nil, fmt.Errorf("LINE_CONFIG_FILE environment variable is not set")
	}
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// unmarshal yaml content into struct
	providerConfig := &ProviderConfig{}
	err = yaml.Unmarshal(fileContent, providerConfig)
	if err != nil {
		return nil, err
	}

	return providerConfig, nil
}

func getFacebookConfig() (*ProviderConfig, error) {
	// read yaml file
	filePath := os.Getenv("FACEBOOK_CONFIG_FILE")
	if filePath == "" {
		return nil, fmt.Errorf("FACEBOOK_CONFIG_FILE environment variable is not set")
	}
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// unmarshal yaml content into struct
	providerConfig := &ProviderConfig{}
	err = yaml.Unmarshal(fileContent, providerConfig)
	if err != nil {
		return nil, err
	}

	return providerConfig, nil
}

func getGoogleConfig() (*ProviderConfig, error) {
	// read yaml file
	filePath := os.Getenv("GOOGLE_CONFIG_FILE")
	if filePath == "" {
		return nil, fmt.Errorf("GOOGLE_CONFIG_FILE environment variable is not set")
	}
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// unmarshal yaml content into struct
	providerConfig := &ProviderConfig{}
	err = yaml.Unmarshal(fileContent, providerConfig)
	if err != nil {
		return nil, err
	}

	return providerConfig, nil
}
