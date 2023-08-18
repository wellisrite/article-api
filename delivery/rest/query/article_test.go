package query

import (
	"article-api/domain/article"
	"article-api/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"context"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock cache client interface
type CacheClient interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string, expiration time.Duration) error
	Del(ctx context.Context, key string) error
}

// Define a mock cache client struct using testify/mock
type MockCacheClient struct {
	mock.Mock
}

func (m *MockCacheClient) Get(ctx context.Context, key string) (string, error) {
	args := m.Called(ctx, key)
	return args.String(0), args.Error(1)
}

func (m *MockCacheClient) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
	args := m.Called(ctx, key, value, expiration)
	return args.Error(0)
}

func (m *MockCacheClient) Del(ctx context.Context, key string) error {
	args := m.Called(ctx, key)
	return args.Error(0)
}

func TestArticleHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/articles", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockSvc := service.Service{}

	handler := NewQueryHandler(mockSvc)
	mockCache := &MockCacheClient{}

	// Set expectations on the mock cache client
	mockCache.On("Get", mock.Anything, "article:John:Sample").Return(nil, nil)

	// Call the handler
	err := handler.ArticleHandler(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseArticles []article.Article
	err = json.Unmarshal(rec.Body.Bytes(), &responseArticles)
	assert.NoError(t, err)
}
