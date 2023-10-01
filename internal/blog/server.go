package blog

import (
	"context"

	"connectrpc.com/connect"
	blogv1 "github.com/Niket1997/blogger/gen/blog/v1"
	"github.com/Niket1997/blogger/gen/blog/v1/blogv1connect"
)

type Server struct {
	blogv1connect.UnimplementedBlogServiceHandler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Create(
	ctx context.Context,
	req *connect.Request[blogv1.CreateBlogRequest],
) (*connect.Response[blogv1.BlogResponse], error) {
	panic("implement me")
}

func (s *Server) Get(
	ctx context.Context,
	req *connect.Request[blogv1.GetBlogRequest],
) (*connect.Response[blogv1.BlogResponse], error) {
	panic("implement me")
}
