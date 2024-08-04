package env

type Server struct {
	host string
	port string
}

func loadServer() (*Server, error) {
	host, err := load("SERVER_HOST")
	if err != nil {
		return nil, err
	}

	port, err := load("SERVER_PORT")
	if err != nil {
		return nil, err
	}

	return &Server{
		host: host,
		port: port,
	}, nil
}

func (s *Server) Host() string {
	return s.host
}

func (s *Server) Port() string {
	return s.port
}
