package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wisnuuakbr/sagala/internal/entities/repository"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/cache"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/datastore"
)

// ArticleUsecase is a struct that holds the dependencies for article-related use cases.
type ArticleUsecase struct {
	ArticleRepo *datastore.ArticleRepository
	Cache       *cache.RedisCache
}

// CreateArticle is a use case that creates a new article and updates the cache.
func (u *ArticleUsecase) CreateArticle(ctx context.Context, article *repository.Article) error {
	article.CreatedAt = time.Now()
	err := u.ArticleRepo.CreateArticle(ctx, article)
	if err != nil {
		return err
	}

	// invalidate cache for all articles
    err = u.Cache.DeleteCache("articles")
    if err != nil {
        fmt.Printf("failed to delete cache: %v\n", err)
    }

	// set cache for the newly created article
	return u.Cache.SetCache(fmt.Sprintf("article_%d", article.ID), article, time.Minute*10)
}

// GetArticle retrieves all articles, checking cache first, then the database.
func (u *ArticleUsecase) GetArticle(ctx context.Context, id int) ([]*repository.Article, error) {
	var articles []*repository.Article
	if err := u.Cache.GetCache("articles", &articles); err == nil {
		return articles, nil
	}

	articles, err := u.ArticleRepo.GetAllArticle(ctx)
	if err != nil {
		return nil, err
	}

	_ = u.Cache.SetCache("articles", articles, time.Minute*10)
	return articles, nil
}

// GetArticleByID retrieves a single article by ID, checking cache first, then the database.
func (u *ArticleUsecase) GetArticleByID(ctx context.Context, id int) (*repository.Article, error) {
	var article repository.Article
	if err := u.Cache.GetCache(fmt.Sprintf("article_%d", id), &article); err == nil {
		return &article, nil
	}

	articleRepository, err := u.ArticleRepo.GetArticleByID(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = u.Cache.SetCache(fmt.Sprintf("article_%d", id), articleRepository, time.Minute*10)
	return articleRepository, nil
}

// SearchArticles retrieves articles based on keyword and author, checking cache first, then the database.
func (u *ArticleUsecase) SearchArticles(ctx context.Context, keyword, author string) ([]*repository.Article, error) {
	var articles []*repository.Article
	cacheKey := fmt.Sprintf("search_articles_%s_%s", keyword, author)
	if err := u.Cache.GetCache(cacheKey, &articles); err == nil {
		return articles, nil
	}

	articles, err := u.ArticleRepo.SearchArticle(ctx, keyword, author)
	if err != nil {
		return nil, err
	}

	if len(articles) == 0 {
		return nil, errors.New("no articles found matching the criteria")
	}

	_ = u.Cache.SetCache(cacheKey, articles, time.Minute*10)
	return articles, nil
}