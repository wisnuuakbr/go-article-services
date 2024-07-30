package httphandler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wisnuuakbr/sagala/internal/entities/repository"
	"github.com/wisnuuakbr/sagala/internal/usecases"
)

type ArticleHandler struct {
	Usecase *usecases.ArticleUsecase
}

// CreateArticle handles create requests for articles
func (h *ArticleHandler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	var article repository.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	if err := h.Usecase.CreateArticle(ctx, &article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// custom response message
	responseMessage := "Data Created Successfully"

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": responseMessage}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetArticle handles get requests for all articles
func (h *ArticleHandler) GetArticle(w http.ResponseWriter, r *http.Request) {
	articles, err := h.Usecase.GetArticle(r.Context(), 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// GetArticleByID handles get requests for articles by ID
func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	article, err := h.Usecase.GetArticleByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

// SearchArticles handles search requests for articles
func (h *ArticleHandler) SearchArticles(w http.ResponseWriter, r *http.Request) {
    keyword := r.URL.Query().Get("keyword")
    author := r.URL.Query().Get("author")

    articles, err := h.Usecase.SearchArticles(context.Background(), keyword, author)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(articles); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}