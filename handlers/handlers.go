package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/linjiansi/blog-api/models"
	"github.com/linjiansi/blog-api/services"
	"net/http"
	"strconv"
)

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	article, err := services.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int

	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := services.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleIDStr := chi.URLParam(req, "articleID")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	article, err := services.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func PostFavoriteArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := services.PostFavoriteService(reqArticle)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := services.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
