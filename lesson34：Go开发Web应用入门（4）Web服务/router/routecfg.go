package router

import (
	"lesson34/router/index"
	"lesson34/router/news"
	"lesson34/router/post"
)

var Routes = []Route{
	{
		Name:        "index",
		Path:        "/",
		Method:      "GET",
		HandlerFunc: index.Index,
	},
	{
		Name:        "get_post",
		Path:        "/post",
		Method:      "GET",
		HandlerFunc: post.GetPost,
	},
	{
		Name:        "post_post",
		Path:        "/post",
		Method:      "POST",
		HandlerFunc: post.PostPost,
	},
	{
		Name:        "news",
		Path:        "/news/{title}",
		Method:      "GET",
		HandlerFunc: news.GetNews,
	},
}
