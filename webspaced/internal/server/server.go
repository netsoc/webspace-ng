package server

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Server is the main webspaced server struct
type Server struct {
	http *http.Server
}

// NewServer returns an initialized Server
func NewServer() *Server {
	r := mux.NewRouter()
	httpSrv := &http.Server{
		Handler: r,
	}

	srv := &Server{
		http: httpSrv,
	}
	r.HandleFunc("/", srv.index)

	return srv
}

// Start begins listening
func (s *Server) Start(sockPath string) error {
	listener, err := net.Listen("unix", sockPath)
	if err != nil {
		return err
	}

	// Socket needs to be u=rw,g=rw,o=rw so anyone can access it (we'll do auth later)
	err = os.Chmod(sockPath, 0o666)
	if err != nil {
		return err
	}

	return s.http.Serve(listener)
}

// Stop shuts down the server and listener
func (s *Server) Stop() error {
	return s.http.Close()
}

func (s *Server) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello, world!\n")
}