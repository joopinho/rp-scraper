package remote

import (
	"github.com/joopinho/rp-scarper/internal/profile"
)

type RemoteProfileEnricher interface {
	Enrich(p *profile.Profile) error
}
