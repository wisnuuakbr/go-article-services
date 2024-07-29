package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/wisnuuakbr/sagala/internal/entities/repository"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/cache"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/datastore"
)

type ArticleUsecase struct {
	ArticleRepo *datastore.ArticleRepository
	Cache       *cache.RedisCache
}

func (u *ArticleUsecase) CreateArticle(ctx context.Context, article *repository.Article) error {
	article.CreatedAt = time.Now()
	err := u.ArticleRepo.Create(ctx, article)
	if err != nil {
		return err
	}
	return u.Cache.Set(fmt.Sprintf("article_%d", article.ID), article, time.Minute*10)
}

func (u *ArticleUsecase) GetArticle(ctx context.Context, id int) ([]*repository.Article, error) {
	var articles []*repository.Article
	if err := u.Cache.Get("articles", &articles); err == nil {
		return articles, nil
	}

	articles, err := u.ArticleRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	_ = u.Cache.Set("articles", articles, time.Minute*10)
	return articles, nil
}

func (u *ArticleUsecase) GetArticleByID(ctx context.Context, id int) (*repository.Article, error) {
	var article repository.Article
	if err := u.Cache.Get(fmt.Sprintf("article_%d", id), &article); err == nil {
		return &article, nil
	}

	articleRepository, err := u.ArticleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	_ = u.Cache.Set(fmt.Sprintf("article_%d", id), articleRepository, time.Minute*10)
	return articleRepository, nil
}
