package queue

import "github.com/jrecuero/go-cli/grafo"

// Server represents ...
type Server struct {
	*grafo.Vertex
}

// NewServer is ...
func NewServer(label string, sc *ServerContent) *Server {
	server := &Server{
		grafo.NewVertex(label),
	}
	server.Content = sc
	return server
}

// ServerToVertex is ...
func ServerToVertex(server *Server) *grafo.Vertex {
	return server.Vertex
}

// ToServer is ...
func ToServer(vertex *grafo.Vertex) *Server {
	return &Server{
		vertex,
	}
}

// GetServerContent is ...
func GetServerContent(server *Server) *ServerContent {
	return server.Content.(*ServerContent)
}
