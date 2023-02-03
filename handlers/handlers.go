package handlers

import (
	"encoding/json"
	"github.com/linjiansi/blog-api/models"
	"net/http"
)

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {

	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(reqArticle)
}

func GetArticleListHandler(w http.ResponseWriter, req *http.Request) {
	//queryMap := req.URL.Query()
	//
	//var page int
	//
	//if p, ok := queryMap["page"]; ok && len(p) > 0 {
	//	var err error
	//	page, err = strconv.Atoi(p[0])
	//	if err != nil {
	//		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
	//		return
	//	}
	//} else {
	//	page = 1
	//}

	article1 := models.Article1
	article2 := models.Article2
	json.NewEncoder(w).Encode([]models.Article{article1, article2})
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	//articleID := chi.URLParam(req, "articleID")
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

func PostFavoriteArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqArticle)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(reqComment)
}
