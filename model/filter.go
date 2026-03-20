package model

type FilterMap map[string]any

type Filter interface {
	ToMap() FilterMap
	Or(filters ...Filter) Filter
	And(filters ...Filter) Filter
}

// DefaultFilter implements Filter  with optimized composition
type DefaultFilter struct {
	FilterMap
}

func NewDefaultFilter() *DefaultFilter {
	return &DefaultFilter{FilterMap: make(FilterMap)}
}

func (m *DefaultFilter) ToMap() FilterMap {
	return m.FilterMap
}

func (m *DefaultFilter) Or(filters ...Filter) Filter {
	var alts []FilterMap
	if len(m.FilterMap) > 0 {
		alts = append(alts, m.FilterMap)
	}

	for _, f := range filters {
		if f == nil {
			continue
		}
		fm := f.ToMap()
		if len(fm) > 0 {
			alts = append(alts, fm)
		}
	}

	switch len(alts) {
	case 0:
		return NewDefaultFilter()
	case 1:
		return &DefaultFilter{FilterMap: alts[0]}
	default:
		return &DefaultFilter{FilterMap: FilterMap{"$or": alts}}
	}
}

func (m *DefaultFilter) And(filters ...Filter) Filter {
	var alts []FilterMap
	if len(m.FilterMap) > 0 {
		alts = append(alts, m.FilterMap)
	}

	for _, f := range filters {
		if f == nil {
			continue
		}
		fm := f.ToMap()
		if len(fm) > 0 {
			alts = append(alts, fm)
		}
	}

	switch len(alts) {
	case 0:
		return NewDefaultFilter()
	case 1:
		return &DefaultFilter{FilterMap: alts[0]}
	default:
		return &DefaultFilter{FilterMap: FilterMap{"$and": alts}}
	}
}
