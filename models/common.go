package models

// OKXResponse is the generic response envelope for all OKX REST API responses.
type OKXResponse[T any] struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data []T    `json:"data"`
}

// Paginator provides cursor-based pagination support.
type Paginator[T any] struct {
	client     interface{}
	fetchFunc  func(after string) ([]T, string, error)
	after      string
	hasMore    bool
	currentBuf []T
	index      int
}

// NewPaginator creates a new paginator.
func NewPaginator[T any](fetchFunc func(after string) ([]T, string, error)) *Paginator[T] {
	return &Paginator[T]{
		fetchFunc: fetchFunc,
		hasMore:   true,
	}
}

// Next returns the next item in the paginated result set.
func (p *Paginator[T]) Next() (T, bool, error) {
	var zero T

	if p.index < len(p.currentBuf) {
		item := p.currentBuf[p.index]
		p.index++
		return item, true, nil
	}

	if !p.hasMore {
		return zero, false, nil
	}

	items, after, err := p.fetchFunc(p.after)
	if err != nil {
		return zero, false, err
	}

	if len(items) == 0 {
		p.hasMore = false
		return zero, false, nil
	}

	p.currentBuf = items
	p.index = 1
	p.after = after

	if after == "" {
		p.hasMore = false
	}

	return items[0], true, nil
}

// All returns all items from all pages.
func (p *Paginator[T]) All() ([]T, error) {
	var result []T
	for {
		item, hasMore, err := p.Next()
		if err != nil {
			return nil, err
		}
		if !hasMore {
			break
		}
		result = append(result, item)
	}
	return result, nil
}
