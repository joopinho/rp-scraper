package services

import (
	"github.com/joopinho/rp-scarper/internal/profile"
	"github.com/joopinho/rp-scarper/internal/profile/remote"
	"github.com/joopinho/rp-scarper/internal/tools"
)

type OrganizationService struct {
	Enrichers *[]remote.RemoteProfileEnricher
}

func (o *OrganizationService) GetOrganization(inn string) (profile.Profile, error) {

	p := profile.NewProfile(inn)

	err := tools.ValidateInn(inn)
	if err != nil {
		return (*p), err
	}
	for _, enr := range *o.Enrichers {
		err = enr.Enrich(p)
		if err != nil {
			return (*p), err
		}
	}

	return (*p), nil
}
