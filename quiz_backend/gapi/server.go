package gapi

import (
	db "github.com/achintya-7/quiz/db/sqlc"
	"github.com/achintya-7/quiz/pb"
	"github.com/achintya-7/quiz/util"
)

type Server struct {
	pb.UnimplementedQuizServiceServer
	store  *db.Store
	config util.Config
}

func NewServer(store *db.Store, config util.Config) (*Server, error) {
	server := &Server{store: store, config: config}

	return server, nil
}
