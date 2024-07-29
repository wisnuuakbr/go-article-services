package di

import (
	"database/sql"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/wisnuuakbr/sagala/internal/adapters/httphandler"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/cache"
	"github.com/wisnuuakbr/sagala/internal/infrastructure/datastore"
	"github.com/wisnuuakbr/sagala/internal/usecases"
)

func NewRouter(db *sql.DB, redisClient *redis.Client) *mux.Router {
	articleRepo := &datastore.ArticleRepository{DB: db}
	redisCache := cache.NewRedisCache(redisClient)
	articleUsecase := &usecases.ArticleUsecase{ArticleRepo: articleRepo, Cache: redisCache}
	articleHandler := &httphandler.ArticleHandler{Usecase: articleUsecase}

	r := mux.NewRouter()
	r.HandleFunc("/articles", articleHandler.GetArticle).Methods("GET")
	r.HandleFunc("/articles", articleHandler.CreateArticle).Methods("POST")

	return r
}
