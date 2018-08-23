package queue

import (
	"github.com/jrecuero/go-cli/engine"
	"github.com/jrecuero/go-cli/grafo"
)

// Server represents ...
type Server struct {
	*grafo.Vertex
}

// ServerEvent is ...
func (server *Server) ServerEvent(jobs *[]*Job, followUp func(*Job)) *engine.Event {
	ev := engine.NewEvent("server/1", 0)
	ev.SetCallback(func(params ...interface{}) error {
		//tools.ToDisplay("que1Event callback\n")
		if j, ok := GetServerContent(server).Serve(jobs); !ok {
			followUp(j.(*Job))
		}
		return nil
	}, nil)
	return ev

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
