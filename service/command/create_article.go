package command

import (
	article "article-api/domain/article"
	"context"
	"fmt"
)

func (s *ArticleCommandService) CreateArticle(ctx context.Context, articleBody article.ArticleWrite) (err error) {
	// Convert the articleBody to the Article struct that can be inserted into the database
	newArticle := article.Article{
		Author: articleBody.Author,
		Title:  articleBody.Title,
		Body:   articleBody.Body,
	}

	// Insert the new article into the database using GORM
	err = s.psqlRepo.Article.CreateArticle(ctx, newArticle)
	if err != nil {
		// Handle the error appropriately, possibly return an error or log it
		return err
	}

	// Invalidate cache for the specific query and author
	pattern := "article:*"
	keysToDelete, err := s.redisClient.Keys(pattern).Result()
	if err != nil {
		// Handle cache key retrieval error
		fmt.Println("Error retrieving cache keys:", err)
	} else {
		// Delete all matching cache keys
		for _, key := range keysToDelete {
			err := s.redisClient.Del(key).Err()
			if err != nil {
				// Handle cache key deletion error
				fmt.Printf("Error deleting cache key %s: %v\n", key, err)
			}
		}
	}

	return nil
}
