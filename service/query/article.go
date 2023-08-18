package query

import (
	"article-api/domain/article"
	"article-api/libs/errors"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (s *ArticleQueryService) Article(ctx context.Context, query, author string) (res []article.Article, err error) {
	// fetch the data from redis in case it exists on the cache
	// if not fetch it form the database
	// Generate a key for the Redis cache based on the query and author
	cacheKey := fmt.Sprintf("article:%s:%s", author, query)
	// Try fetching the article from Redis cache
	cachedArticle, err := s.redisClient.Get(cacheKey).Result()
	if err == nil {
		// Article found in cache
		err = json.Unmarshal([]byte(cachedArticle), &res)
		if err != nil {
			return res, err
		}
		return res, nil
	}

	// Article not found in cache, fetch from the database
	// Implement your database query here to fetch the article based on the query and author
	// For example:
	res, err = s.psqlRepo.Article.GetArticle(ctx, query, author)
	if err != nil && len(res) == 0 {
		return res, fmt.Errorf(errors.ErrRecordNotFound)
	} else if err != nil {
		return res, err
	}

	go s.saveArticleCache(cacheKey, res)

	return res, nil
}

func (s *ArticleQueryService) saveArticleCache(cacheKey string, res []article.Article) {
	// Store the fetched article in the cache for future use
	cacheDuration := time.Hour * 10 // Adjust cache duration as needed
	dbArticleJSON, _ := json.Marshal(res)
	err := s.redisClient.Set(cacheKey, dbArticleJSON, cacheDuration).Err()
	if err != nil {
		// Log or handle the error as needed
		log.Println(err)
	}
}
