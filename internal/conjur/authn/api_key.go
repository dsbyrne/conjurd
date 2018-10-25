package authn

import (
	"fmt"

	"github.com/cyberark/conjur-api-go/conjurapi"
	conjurAuthn "github.com/cyberark/conjur-api-go/conjurapi/authn"
	"github.com/rs/zerolog/log"
)

// APIKey represents API key authentication provided by:
// - environment
// - netrc
// - hostfactory
type APIKey struct {
	loginPair *conjurAuthn.LoginPair
	client    *conjurapi.Client
}

// Authenticate does the authentication to Conjur and returns an access token
func (authn *APIKey) Authenticate() ([]byte, error) {
	if authn.client == nil {
		return nil, fmt.Errorf("no client available")
	}

	return authn.client.Authenticate(*authn.loginPair)
}

// IsAvailable returns whether or not this authentication method should be tried
func (authn *APIKey) IsAvailable() bool {
	if authn.client == nil {
		err := authn.initialize()
		if err != nil {
			log.Debug().
				Err(err).
				Str("method", "api_key").
				Msg("requirements not met, method unavailable")
			return false
		}
	}
	return true
}

func (authn *APIKey) hasValidLogin() bool {
	if authn.loginPair == nil {
		return false
	}

	return len(authn.loginPair.APIKey) != 0 && len(authn.loginPair.Login) != 0
}

func (authn *APIKey) initialize() error {
	config, err := conjurapi.LoadConfig()
	if err != nil {
		return err
	}

	authn.client, err = conjurapi.NewClientFromEnvironment(config)
	if err != nil {
		return err
	}

	// HACK
	// Our client's login pair is not exported, so we must duplicate it
	authn.loginPair, _ = conjurapi.LoginPairFromEnv()
	if authn.hasValidLogin() {
		return nil
	}

	authn.loginPair, err = conjurapi.LoginPairFromNetRC(config)
	return err
}
