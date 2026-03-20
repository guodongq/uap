package set

// Set is a generic set data structure backed by a map.
type Set[T comparable] map[T]struct{}

// New creates a new Set with the given items.
func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, item := range items {
		s[item] = struct{}{}
	}
	return s
}

func (s Set[T]) Add(item T) {
	s[item] = struct{}{}
}

func (s Set[T]) Remove(item T) {
	delete(s, item)
}

func (s Set[T]) Exists(item T) bool {
	_, ok := s[item]
	return ok
}

func (s Set[T]) Len() int {
	return len(s)
}

func (s Set[T]) Items() []T {
	items := make([]T, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}

func (s Set[T]) Intersection(other Set[T]) Set[T] {
	newSet := make(Set[T])

	// iterate over the smaller set for efficiency
	small, big := s, other
	if len(s) > len(other) {
		small, big = other, s
	}

	for key := range small {
		if _, in := big[key]; in {
			newSet[key] = struct{}{}
		}
	}
	return newSet
}

func (s Set[T]) Union(other Set[T]) Set[T] {
	newSet := make(Set[T], len(s)+len(other))
	for key := range s {
		newSet[key] = struct{}{}
	}
	for key := range other {
		newSet[key] = struct{}{}
	}
	return newSet
}

func (s Set[T]) Difference(other Set[T]) Set[T] {
	newSet := make(Set[T])
	for key := range s {
		if _, in := other[key]; !in {
			newSet[key] = struct{}{}
		}
	}
	return newSet
}

func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for key := range s {
		if _, in := other[key]; !in {
			return false
		}
	}
	return true
}
