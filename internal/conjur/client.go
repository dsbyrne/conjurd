package conjur

import (
	"fmt"

	"github.com/dsbyrne/conjurd/internal/conjur/authn"
	"github.com/rs/zerolog/log"
)

var accessToken []byte
var authnMethods = map[string]authn.Method{
	"api_key":    &authn.APIKey{},
	"aws":        &authn.AWS{},
	"kubernetes": &authn.Kubernetes{},
}
var preferredMethod authn.Method

func authenticate() ([]byte, error) {
	if preferredMethod == nil {
		return nil, fmt.Errorf("no preferred method")
	}

	token, err := preferredMethod.Authenticate()
	if err != nil {
		return nil, err
	}

	if len(token) == 0 {
		return nil, fmt.Errorf("no authentications methods able to succeed")
	}

	return token, nil
}

func Initialize() {
	for name, authnMethod := range authnMethods {
		if authnMethod.IsAvailable() {
			log.Info().Msgf("%s selected as authentication method", name)
			preferredMethod = authnMethod
			return
		}
	}

	log.Warn().Msg("no method of authentication identified")
}

func GetToken() []byte {
	token, err := authenticate()
	if err != nil {
		log.Error().
			Err(err).
			Msg("failed to authenticate")
	}

	return token
}
