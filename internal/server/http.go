package server

type Http struct {
	Port string
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(port string) *Http {
	return &Http{
		Port: port,
	}
}