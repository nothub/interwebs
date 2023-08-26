package interwebs

import "net/http"

type Server struct {
	*http.ServeMux
}

func NewServer() (srv *Server) {
	srv = &Server{http.NewServeMux()}

	// TODO: default handlers

	return srv
}
