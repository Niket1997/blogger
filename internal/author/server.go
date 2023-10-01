package author

import (
	"context"

	"connectrpc.com/connect"
	authorv1 "github.com/Niket1997/blogger/gen/author/v1"
	"github.com/Niket1997/blogger/gen/author/v1/authorv1connect"
)

type Server struct {
	authorv1connect.UnimplementedAuthorServiceHandler
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) create(
	ctx context.Context,
	req *connect.Request[authorv1.CreateAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	panic("implement me")
}

func (s *Server) get(
	ctx context.Context,
	req *connect.Request[authorv1.GetAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	panic("implement me")
}
