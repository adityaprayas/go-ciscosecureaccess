//go:build functional

package tests

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContentCategoriesFunctionalList(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetContentCategoriesClient(env.ctx)

	// Fetch first page of category settings.
	categories, httpResp, err := apiClient.ContentCategoriesAPI.
		GetCategorySettings(env.ctx).
		Page(1).
		Limit(100).
		Execute()
	if err != nil && !isHTTP2xx(httpResp) {
		require.NoError(t, err)
	}

	require.NotEmpty(t, categories, "expected at least one content category setting")

	for _, cat := range categories {
		require.NotZero(t, cat.GetId(), "category Id should be set")
		require.NotEmpty(t, cat.GetName(), "category Name should be set")
	}
}

func TestContentCategoriesFunctionalPagination(t *testing.T) {
	env := newFunctionalEnv(t)
	apiClient := env.factory.GetContentCategoriesClient(env.ctx)

	// Fetch page 1 with a small limit.
	page1, httpResp, err := apiClient.ContentCategoriesAPI.
		GetCategorySettings(env.ctx).
		Page(1).
		Limit(5).
		Execute()
	if err != nil && !isHTTP2xx(httpResp) {
		require.NoError(t, err)
	}

	if len(page1) < 5 {
		// Fewer total results than the page size; nothing more to paginate.
		t.Skip("not enough content categories to test pagination")
	}

	// Fetch page 2 and verify it differs from page 1.
	page2, httpResp2, err := apiClient.ContentCategoriesAPI.
		GetCategorySettings(env.ctx).
		Page(2).
		Limit(5).
		Execute()
	if err != nil && !isHTTP2xx(httpResp2) {
		require.NoError(t, err)
	}

	if len(page2) == 0 {
		t.Skip("second page returned no results; total count <= page size")
	}

	require.NotEqual(t, page1[0].GetId(), page2[0].GetId(),
		"first item on page 2 should differ from first item on page 1")
}
