package authn

import (
	"fmt"
	"net"
	"os"

	"github.com/rs/zerolog/log"
)

// Kubernetes authenticates via authn-k8s
type Kubernetes struct {
}

// Authenticate does the authentication to Conjur and returns an access token
func (authn *Kubernetes) Authenticate() ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// IsAvailable returns whether or not this authentication method should be tried
func (authn *Kubernetes) IsAvailable() bool {
	if _, err := os.Stat("/.dockerenv"); os.IsNotExist(err) {
		log.Debug().
			Err(err).
			Str("method", "kubernetes").
			Msg("not running in docker context")
		return false
	}

	_, err := net.LookupIP("kubernetes.default.svc.cluster.local")
	if err != nil {
		log.Debug().
			Err(err).
			Str("method", "kubernetes").
			Msg("not running in Kubernetes")
		return false
	}

	return true
}
