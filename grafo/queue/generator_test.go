package queue_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo/queue"
)

// TestGen_Generate is ...
func TestGen_Generate(t *testing.T) {
	system := queue.NewSystem("system/1")
	gen0 := queue.NewGenerator("gen/0", queue.NewGeneratorContent("gen/0", 3))
	queue1 := queue.NewServer("queue/1", queue.NewServerContent("queue/1", 5))
	system.AddQueue(nil, queue.GeneratorToVertex(gen0), 0)
	system.AddQueue(queue.GeneratorToVertex(gen0), queue.ServerToVertex(queue1), 10)

	if job, ok := queue.GetGeneratorContent(gen0).Generate(); ok {
		edge := gen0.Edges[0]
		edge.Check(job)
		jobs := &gen0.Edges[0].(*queue.Queue).Jobs
		if j, ok := queue.GetServerContent(queue1).Serve(jobs); !ok {
			t.Errorf("Queue:Serve: job should be completed: %#v\n", j)
		}
	}
}
