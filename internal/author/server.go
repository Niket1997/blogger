package author

import (
	"context"
	"log"

	"connectrpc.com/connect"
	authorv1 "github.com/Niket1997/blogger/gen/author/v1"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Create(
	ctx context.Context,
	req *connect.Request[authorv1.CreateAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	log.Println("request received", req)
	res := connect.NewResponse(&authorv1.AuthorResponse{
		AuthorId:  "id12345",
		Name:      "Niket",
		EmailId:   "nik@gmail.com",
		Bio:       "I'm awesome",
		CreatedAt: 1,
		UpdatedAt: 1,
	})
	res.Header().Set("Author-Version", "v1")
	return res, nil
}

func (s *Server) Get(
	ctx context.Context,
	req *connect.Request[authorv1.GetAuthorRequest],
) (*connect.Response[authorv1.AuthorResponse], error) {
	panic("implement me")
}
