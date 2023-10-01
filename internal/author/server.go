package author

import (
	"context"
	"log"

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

func (s *Server) Create(
	ctx context.Context,
	req *connect.Request[authorv1.CreateAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	log.Println("request received", req)
	return &connect.Response[authorv1.AuthorResponse]{
		Msg: &authorv1.AuthorResponse{
			AuthorId:  "id12345",
			Name:      "Niket",
			EmailId:   "niket@gmail.com",
			Bio:       "I'm awesome",
			CreatedAt: 0,
			UpdatedAt: 0,
		},
	}, nil
}

func (s *Server) Get(
	ctx context.Context,
	req *connect.Request[authorv1.GetAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	panic("implement me")
}
