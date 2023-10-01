package blog

import (
	"context"
	"log"

	"connectrpc.com/connect"
	blogv1 "github.com/Niket1997/blogger/gen/blog/v1"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Create(
	ctx context.Context,
	req *connect.Request[blogv1.CreateBlogRequest],
) (*connect.Response[blogv1.BlogResponse], error) {
	log.Println("request received", req)
	res := connect.NewResponse(&blogv1.BlogResponse{
		BlogId:      "blog12345",
		AuthorId:    "author12345",
		Title:       "How to write gRPC clients in go",
		Body:        "Some Body",
		PublishedAt: 1,
		UpdatedAt:   1,
	})
	res.Header().Set("Blog-Version", "v1")
	return res, nil
}

func (s *Server) Get(
	ctx context.Context,
	req *connect.Request[blogv1.GetBlogRequest],
) (*connect.Response[blogv1.BlogResponse], error) {
	log.Println("request received", req)
	res := connect.NewResponse(&blogv1.BlogResponse{
		BlogId:      "blog12345",
		AuthorId:    "author12345",
		Title:       "How to write gRPC clients in go",
		Body:        "Some Body",
		PublishedAt: 1,
		UpdatedAt:   1,
	})
	res.Header().Set("Blog-Version", "v1")
	return res, nil
}
