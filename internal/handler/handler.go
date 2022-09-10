package handler

import (
	"database/sql"
	api "github.com/dena-gohost/gohost-server/gen/api"
)

var _ api.ServerInterface = (*Server)(nil)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}
