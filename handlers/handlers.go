package handlers

import (
	"encoding/json"
	"github.com/linjiansi/blog-api/models"
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world! \n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
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
	jsonData, err := json.Marshal([]models.Article{article1, article2})
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func GetArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	//articleID := chi.URLParam(req, "articleID")
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func PostFavoriteArticleHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}
