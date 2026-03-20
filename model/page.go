package model

func WithPageSkip[T any](skip int64) func(page *Page[T]) {
	return func(page *Page[T]) {
		page.Skip = skip
	}
}

func WithPageLimit[T any](limit int64) func(page *Page[T]) {
	return func(page *Page[T]) {
		page.Limit = limit
	}
}

func WithPageTotal[T any](total int64) func(page *Page[T]) {
	return func(page *Page[T]) {
		page.Total = total
	}
}

func WithPageContents[T any](contents []T) func(page *Page[T]) {
	return func(page *Page[T]) {
		page.Contents = contents
	}
}

type Page[T any] struct {
	Skip     int64
	Limit    int64
	Total    int64
	Contents []T
}

func NewPage[T any](optionFuncs ...func(page *Page[T])) *Page[T] {
	var p Page[T]
	for _, optionFunc := range optionFuncs {
		optionFunc(&p)
	}
	return &p
}
