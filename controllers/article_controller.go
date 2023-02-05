package controllers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/linjiansi/blog-api/controllers/services"
	"github.com/linjiansi/blog-api/models"
	"net/http"
	"strconv"
)

type ArticleController struct {
	service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
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

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(articleList)
}

func (c *ArticleController) GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleIDStr := chi.URLParam(req, "articleID")
	articleID, err := strconv.Atoi(articleIDStr)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostFavoriteArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostFavoriteService(reqArticle)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(article)
}
