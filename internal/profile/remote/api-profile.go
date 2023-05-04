package remote

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/joopinho/rp-scarper/configs"
	"github.com/joopinho/rp-scarper/internal/profile"
	"github.com/joopinho/rp-scarper/internal/tools"
)

type ApiProfileEnricher struct {
	client *http.Client
	cfg    *configs.RemoteApiProfile
}

type Organization struct {
	Name    string `json:"name"`
	CeoName string `json:"ceo_name"`
}

type RPResponse struct {
	UlCount int            `json:"ul_count"`
	UlList  []Organization `json:"ul"`
	IPCount int            `json:"ip_count"`
	IPList  []Organization `json:"ip"`
}

func NewApiProfileEnricher(cfg *configs.RemoteApiProfile) *ApiProfileEnricher {
	en := &ApiProfileEnricher{cfg: cfg}
	en.Init()
	return en
}

func (en *ApiProfileEnricher) Init() {
	en.client = &http.Client{Timeout: 10 * time.Second}
}

func (en *ApiProfileEnricher) Enrich(p *profile.Profile) error {

	url := en.cfg.Url + p.INN()

	log.Printf("requesting remote url: %s", url)
	r, err := en.client.Get(url)
	if err != nil {
		return &tools.ServiceError{Code: 13, Err: err}
	}
	defer r.Body.Close()

	var rpResponse RPResponse
	err = json.NewDecoder(r.Body).Decode(&rpResponse)

	if err != nil {
		return &tools.ServiceError{Code: 13, Err: err}
	}

	if rpResponse.IPCount > 1 || rpResponse.UlCount > 1 {
		return &tools.ServiceError{Code: 3, Err: errors.New("more then one organization found")}
	}
	if rpResponse.IPCount == 0 && rpResponse.UlCount == 0 {
		return &tools.ServiceError{Code: 5, Err: errors.New("organization not found")}
	}

	if rpResponse.IPCount == 1 {
		p.SetCEO(rpResponse.IPList[0].Name)
		p.SetName(rpResponse.IPList[0].Name)
	}

	if rpResponse.UlCount == 1 {
		p.SetCEO(rpResponse.UlList[0].CeoName)
		p.SetName(rpResponse.UlList[0].Name)
	}

	return nil
}
