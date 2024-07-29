package datastore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/wisnuuakbr/sagala/internal/entities/repository"
)

type ArticleRepository struct {
	DB *sql.DB
}

// function for insert article data to database
func (a *ArticleRepository) Create(ctx context.Context, article *repository.Article) error {
	query := `INSERT INTO articles (author, title, body, created_at) VALUES ($1, $2, $3, $4)`
	_, err := a.DB.ExecContext(ctx, query, article.Author, article.Title, article.Body, article.CreatedAt)
	return err
}

// function for get all article data from database
func (a *ArticleRepository) GetAll(ctx context.Context) ([]*repository.Article, error) {
	query := `SELECT id, author, title, body, created_at 
			FROM articles
			ORDER BY created_at 
			DESC`
	rows, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*repository.Article
	for rows.Next() {
		var article repository.Article
		err := rows.Scan(&article.ID, &article.Author, &article.Title, &article.Body, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, &article)
	}
	return articles, nil
}

// function for get article by id from database
func (a *ArticleRepository) GetByID(ctx context.Context, id int) (*repository.Article, error) {
	query := `SELECT id, author, title, body, created_at 
            FROM articles
            WHERE id = $1`
	row := a.DB.QueryRowContext(ctx, query, id)

	var article repository.Article
	err := row.Scan(&article.ID, &article.Author, &article.Title, &article.Body, &article.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("article not found")
		}
		return nil, err
	}
	return &article, nil
}
