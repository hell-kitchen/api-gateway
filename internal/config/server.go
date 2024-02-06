package config

type Server struct {
	BindAddr string
	BindPort int
}

// NewServer return new Server config.
func NewServer() (*Server, error) {
	s := new(Server)
	return s, nil
}
