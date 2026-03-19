package set

type void = struct{}

type Set map[any]void

func (s Set) Add(item any) {
	s[item] = void{}
}

func (s Set) Exists(item any) bool {
	_, ok := s[item]

	return ok
}

func (s Set) Items() []any {
	var items []any
	for item := range s {
		items = append(items, item)
	}

	return items
}

func (s Set) Intersection(other Set) Set {
	if len(s) == 0 {
		return other
	}

	if len(other) == 0 {
		return s
	}

	newSet := Set{}
	for key := range s {
		if _, in := other[key]; in {
			newSet.Add(key)
		}
	}
	return newSet
}

func (s Set) Difference(other Set) Set {
	newSet := Set{}
	for key := range s {
		if _, in := other[key]; !in {
			newSet[key] = struct{}{}
		}
	}
	return newSet
}

func (s Set) Equal(other Set) bool {
	if len(s) != len(other) {
		return false
	}
	return len(s.Intersection(other)) == len(s)
}
