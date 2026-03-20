package auth

import "maps"

type Properties struct {
	values map[any]any
}

func (m *Properties) Get(key any) any {
	m.lazyInit()
	return m.values[key]
}

func (m *Properties) Set(key, value any) {
	m.lazyInit()
	m.values[key] = value
}

func (m *Properties) Has(key any) bool {
	m.lazyInit()
	_, ok := m.values[key]
	return ok
}

func (m *Properties) SetAll(other *Properties) {
	if other.values == nil {
		return
	}

	m.lazyInit()
	maps.Copy(m.values, other.values)
}

func (m *Properties) Values() map[any]any {
	return maps.Clone(m.values)
}

func (m *Properties) lazyInit() {
	if m.values == nil {
		m.values = map[any]any{}
	}
}

type (
	authOptionsKey struct{}
)

type Option struct {
	SchemeID           string
	IdentityProperties Properties
	SignerProperties   Properties
}

func GetAuthOptions(p *Properties) ([]*Option, bool) {
	v, ok := p.Get(authOptionsKey{}).([]*Option)
	return v, ok
}

func SetAuthOptions(p *Properties, options []*Option) {
	p.Set(authOptionsKey{}, options)
}
