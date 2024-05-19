package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"errors"

	"github.com/gorilla/mux"
	"github.com/tacchanmaru/test_go_api/models"
)


func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request){

	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	reqBodybuffer := make([]byte, length)

	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	defer req.Body.Close()

	var reqArticle models.Article
	if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	article := reqArticle
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)

}

func ArticleListHandler(w http.ResponseWriter, req *http.Request){
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json(page %d)", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request){
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid article ID", http.StatusBadRequest)
		return
	}

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (articleID %d)", articleID)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request){
	article := models.Article1.Contents
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request){
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	
	w.Write(jsonData)
}