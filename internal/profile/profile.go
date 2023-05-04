package profile

type Profile struct {
	inn  string
	kpp  string
	ceo  string
	name string
}

func NewProfile(inn string) *Profile {
	p := &Profile{inn: inn}
	return p
}

func (p *Profile) SetKPP(kpp string) {
	if p.kpp == "" {
		p.kpp = kpp
	}
}

func (p *Profile) SetCEO(ceo string) {
	if p.ceo == "" {
		p.ceo = ceo
	}
}

func (p *Profile) SetName(name string) {
	if p.name == "" {
		p.name = name
	}
}

func (p *Profile) INN() string {
	inn := p.inn
	return inn
}

func (p *Profile) CEO() string {
	ceo := p.ceo
	return ceo
}

func (p *Profile) Name() string {
	name := p.name
	return name
}

func (p *Profile) KPP() string {
	kpp := p.kpp
	return kpp
}
