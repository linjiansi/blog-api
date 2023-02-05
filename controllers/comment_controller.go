package controllers

import (
	"encoding/json"
	"github.com/linjiansi/blog-api/controllers/services"
	"github.com/linjiansi/blog-api/models"
	"net/http"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	comment, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to get request body\n", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(comment)
}
