package q

import (
	"fmt"
	"strings"
)

const (
	DefaultSkip  int64 = 0
	DefaultLimit int64 = 100
	MaxLimit     int64 = 1000
)

type Sort struct {
	Field string
	Desc  bool
}

func (s Sort) IsValid() bool {
	return s.Field != ""
}

func (s Sort) Direction() string {
	if s.Desc {
		return "desc"
	}
	return "asc"
}

func (s Sort) DirectionValue() int {
	if s.Desc {
		return -1
	}
	return 1
}
func (s Sort) String() string {
	return fmt.Sprintf("%s:%s", s.Field, s.Direction())
}

type Sorts []Sort

func DefaultSorts() Sorts {
	return Sorts{}
}

func NewSorts(sorts ...Sort) Sorts {
	validSorts := make(Sorts, 0, len(sorts))
	for _, sort := range sorts {
		if sort.IsValid() {
			validSorts = append(validSorts, sort)
		}
	}
	return validSorts
}

func (s *Sorts) Add(field string, desc bool) {
	if field == "" {
		return
	}
	*s = append(*s, Sort{Field: field, Desc: desc})
}
func (s *Sorts) AddAsc(field string) {
	s.Add(field, false)
}

func (s *Sorts) AddDesc(field string) {
	s.Add(field, true)
}

func (s *Sorts) AddUnique(field string, desc bool) {
	if field == "" {
		return
	}

	for i, existingSort := range *s {
		if existingSort.Field == field {
			(*s)[i].Desc = desc
			return
		}
	}

	s.Add(field, desc)
}

func (s *Sorts) Clear() {
	*s = nil
}

func (s Sorts) IsEmpty() bool {
	return len(s) == 0
}

func (s *Sorts) String() string {
	var parts []string
	for _, sort := range *s {
		parts = append(parts, sort.String())
	}
	return strings.Join(parts, ",")
}

type QueryOptions struct {
	skip  *int64
	limit *int64
	sorts Sorts
}

func NewQuery() *QueryOptions {
	return &QueryOptions{
		sorts: DefaultSorts(),
	}
}

func (q *QueryOptions) Skip() int64 {
	if q.skip == nil || *q.skip < 0 {
		return DefaultSkip
	}
	return *q.skip
}

func (q *QueryOptions) Limit() int64 {
	if q.limit == nil || *q.limit < 0 {
		return DefaultLimit
	}
	if *q.limit > MaxLimit {
		return MaxLimit
	}
	return *q.limit
}

func (q *QueryOptions) Sorts() Sorts {
	return q.sorts
}

func (q *QueryOptions) SetSkip(skip int64) *QueryOptions {
	if skip < 0 {
		skip = DefaultSkip
	}
	q.skip = &skip
	return q
}

func (q *QueryOptions) SetLimit(limit int64) *QueryOptions {
	if limit < 0 {
		limit = DefaultLimit
	} else if limit > MaxLimit {
		limit = MaxLimit
	}
	q.limit = &limit
	return q
}

func (q *QueryOptions) AddSort(field string, desc bool) *QueryOptions {
	if field != "" {
		q.sorts.Add(field, desc)
	}
	return q
}

func (q *QueryOptions) AddSortAsc(field string) *QueryOptions {
	return q.AddSort(field, false)
}

func (q *QueryOptions) AddSortDesc(field string) *QueryOptions {
	return q.AddSort(field, true)
}

func (q *QueryOptions) SetSorts(sorts ...Sort) *QueryOptions {
	q.sorts = NewSorts(sorts...)
	return q
}

func (q *QueryOptions) ClearSorts() *QueryOptions {
	q.sorts.Clear()
	return q
}

func (q *QueryOptions) Page() int64 {
	skip := q.Skip()
	limit := q.Limit()
	if limit <= 0 {
		return 1
	}
	return (skip / limit) + 1
}

func (q *QueryOptions) SetPage(page, size int64) *QueryOptions {
	if page < 1 {
		page = 1
	}
	return q.SetSkip((page - 1) * size).SetLimit(size)
}
func (q *QueryOptions) Normalize() *QueryOptions {
	// Normalize Skip and Limit
	q.SetSkip(q.Skip()).SetLimit(q.Limit())
	q.sorts = NewSorts(q.sorts...)

	return q
}

func MergeQueryOptions(opts ...*QueryOptions) *QueryOptions {
	merged := NewQuery()
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		if opt.skip != nil {
			merged.skip = opt.skip
		}

		if opt.limit != nil {
			merged.limit = opt.limit
		}

		// For slices/pointers to slices, decide on merge strategy.
		// Here, we replace the entire slice.
		for _, sort := range opt.sorts {
			merged.sorts.AddUnique(sort.Field, sort.Desc)
		}
	}
	return merged.Normalize()
}

func WithSkip(skip int64) *QueryOptions {
	return NewQuery().SetSkip(skip)
}

func WithLimit(limit int64) *QueryOptions {
	return NewQuery().SetLimit(limit)
}

func WithPage(page, size int64) *QueryOptions {
	return NewQuery().SetPage(page, size)
}

func WithSort(field string, desc bool) *QueryOptions {
	return NewQuery().AddSort(field, desc)
}

func WithSortAsc(field string) *QueryOptions {
	return NewQuery().AddSortAsc(field)
}

func WithSortDesc(field string) *QueryOptions {
	return NewQuery().AddSortDesc(field)
}

func WithSorts(sorts ...Sort) *QueryOptions {
	return NewQuery().SetSorts(sorts...)
}
