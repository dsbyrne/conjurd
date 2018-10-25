package authn

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
)

const (
	uuidPath = "/sys/hypervisor/uuid"
)

// AWS authenticates via IAM role
type AWS struct {
}

// Authenticate does the authentication to Conjur and returns an access token
func (authn *AWS) Authenticate() ([]byte, error) {
	return nil, fmt.Errorf("not yet implemented")
}

// IsAvailable returns whether or not this authentication method should be tried
func (authn *AWS) IsAvailable() bool {
	if _, err := os.Stat(uuidPath); os.IsNotExist(err) {
		return false
	}

	content, err := ioutil.ReadFile(uuidPath)
	if err != nil {
		log.Error().
			Err(err).
			Str("method", "aws").
			Msgf("failed reading %s", uuidPath)
		return false
	}

	if len(content) < 3 {
		log.Debug().
			Str("method", "aws").
			Msgf("%s does not meet expected content length", uuidPath)
		return false
	}

	header := strings.ToLower(string(content[:3]))
	if header != "ec2" {
		log.Debug().
			Str("method", "aws").
			Msgf("%s does not begin with 'ec2'")
		return false
	}

	resp, err := http.Get("http://169.254.169.254/latest/meta-data/instance-id")
	if err != nil {
		log.Debug().
			Err(err).
			Str("method", "aws").
			Msg("failed making request to ec2 meta-data service")
		return false
	}

	if resp.StatusCode != 200 {
		log.Debug().
			Str("method", "aws").
			Msgf("received %s response from ec2 meta-data service", resp.StatusCode)
		return false
	}

	return true
}
