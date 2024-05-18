package models 

import "time"

var (
	Comment1 = Comment{
		CommentID: 1,
		ArticleID: 1,
		Message: "This is a good article",
		CreatedAt: time.Now(),
	}

	Comment2 = Comment{
		CommentID: 2,
		ArticleID: 1,
		Message: "I think this article is not good",
		CreatedAt: time.Now(),
	}
)

var (
	Article1 = Article{
		ID: 1,
		Title : "Article1",
		Contents: "This is article1",
		UserName: "tacchan",
		NiceNum: 10,
		CommentList: []Comment{Comment1, Comment2},
		CreatedAt: time.Now(),
	}

	Article2 = Article{	
		ID: 2,
		Title : "Article2",
		Contents: "This is article2",
		UserName: "tacchan",
		NiceNum: 20,
		CommentList: []Comment{},
		CreatedAt: time.Now(),
	}
)