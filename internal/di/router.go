package di

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/wisnuuakbr/sagala/internal/adapters/httphandler"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/datastore"
	"github.com/wisnuuakbr/sagala/internal/usecases"
	"github.com/wisnuuakbr/sagala/pkg/cache"
)

func NewRouter(db *sql.DB, redisClient *redis.Client) *mux.Router {
	articleRepo := &datastore.ArticleRepository{DB: db}
	redisCache := cache.NewRedisCache(redisClient)
	articleUsecase := &usecases.ArticleUsecase{ArticleRepo: articleRepo, Cache: redisCache}
	articleHandler := &httphandler.ArticleHandler{Usecase: articleUsecase}

	r := mux.NewRouter()
	r.HandleFunc("/articles", articleHandler.GetArticle).Methods("GET")
	r.HandleFunc("/articles", articleHandler.CreateArticle).Methods("POST")
	r.HandleFunc("/articles/search", articleHandler.SearchArticles).Methods("GET")
	r.HandleFunc("/articles/{id}", articleHandler.GetArticleByID).Methods("GET")

	return r
}
