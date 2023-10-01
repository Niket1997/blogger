package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	authorv1 "github.com/Niket1997/blogger/gen/author/v1"
	"github.com/Niket1997/blogger/gen/author/v1/authorv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

const address = "localhost:8080"

func main() {
	mux := http.NewServeMux()
	authorPath, authorHandler := authorv1connect.NewAuthorServiceHandler(&authorServiceServer{})
	//blogPath, blogHandler := blogv1connect.NewBlogServiceHandler(&blogServiceServer{})
	mux.Handle(authorPath, authorHandler)
	//mux.Handle(blogPath, blogHandler)

	fmt.Println("... Listening on", address)
	err := http.ListenAndServe(
		address,
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
}

type authorServiceServer struct {
	authorv1connect.UnimplementedAuthorServiceHandler
}

func (s *authorServiceServer) create(
	ctx context.Context,
	req *connect.Request[authorv1.CreateAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	panic("implement me")
}

func (s *authorServiceServer) get(
	ctx context.Context,
	req *connect.Request[authorv1.GetAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	panic("implement me")
}

//type blogServiceServer struct {
//	blogv1connect.UnimplementedBlogServiceHandler
//}
