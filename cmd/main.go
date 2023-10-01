package main

import (
	"log"
	"net/http"

	"github.com/Niket1997/blogger/gen/blog/v1/blogv1connect"
	"github.com/Niket1997/blogger/internal/author"
	"github.com/Niket1997/blogger/internal/blog"

	"github.com/Niket1997/blogger/gen/author/v1/authorv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	authorServer := author.NewServer()
	blogServer := blog.NewServer()
	authorPath, authorHandler := authorv1connect.NewAuthorServiceHandler(authorServer)
	blogPath, blogHandler := blogv1connect.NewBlogServiceHandler(blogServer)
	mux.Handle(authorPath, authorHandler)
	mux.Handle(blogPath, blogHandler)

	log.Println("... Listening on", address)

	err := http.ListenAndServe(
		address,
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}
