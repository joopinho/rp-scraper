package remote

import (
	"log"

	"github.com/gocolly/colly"
	"github.com/joopinho/rp-scarper/configs"
	"github.com/joopinho/rp-scarper/internal/profile"
	"github.com/joopinho/rp-scarper/internal/tools"
)

type PageProfileEnricher struct {
	cfg *configs.RemotePageProfile
	fc  *colly.Collector
	p   *profile.Profile
}

func NewPageProfileEnricher(cfg *configs.RemotePageProfile) *PageProfileEnricher {
	en := &PageProfileEnricher{cfg: cfg}
	en.Init()

	return en
}

func (en *PageProfileEnricher) Init() {

	en.fc = colly.NewCollector()
	en.fc.OnHTML("#clip_kpp", func(e *colly.HTMLElement) {
		en.p.SetKPP(e.Text)
	})

}

func (en *PageProfileEnricher) Enrich(p *profile.Profile) error {
	en.p = p

	url := en.cfg.Url + p.INN()
	log.Printf("visiting remote page: %s", url)
	err := en.fc.Visit(url)

	if err != nil {
		return &tools.ServiceError{Code: 13, Err: err}
	}
	return nil
}
