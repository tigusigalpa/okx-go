package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaginatorAllFetchesEveryPage(t *testing.T) {
	calls := 0
	paginator := NewPaginator(func(after string) ([]int, string, error) {
		calls++
		switch after {
		case "":
			return []int{1, 2}, "cursor-1", nil
		case "cursor-1":
			return []int{3}, "", nil
		default:
			return nil, "", nil
		}
	})

	items, err := paginator.All()
	require.NoError(t, err)
	require.Equal(t, []int{1, 2, 3}, items)
	require.Equal(t, 2, calls)

	_, hasMore, err := paginator.Next()
	require.NoError(t, err)
	require.False(t, hasMore)
}

func TestPaginatorReturnsFetchError(t *testing.T) {
	want := errors.New("request failed")
	paginator := NewPaginator(func(string) ([]string, string, error) {
		return nil, "", want
	})

	_, hasMore, err := paginator.Next()
	require.ErrorIs(t, err, want)
	require.False(t, hasMore)
}

func TestPaginatorStopsOnEmptyPage(t *testing.T) {
	paginator := NewPaginator(func(string) ([]string, string, error) {
		return nil, "next", nil
	})

	_, hasMore, err := paginator.Next()
	require.NoError(t, err)
	require.False(t, hasMore)
}
