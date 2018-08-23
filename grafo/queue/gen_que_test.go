package queue_test

import (
	"testing"
	"time"

	"github.com/jrecuero/go-cli/engine"
	"github.com/jrecuero/go-cli/grafo/queue"
)

// TestGenQue_Engine is ...
func TestGenQue_Engine(t *testing.T) {
	pipe := queue.NewSystem("pipe/1")
	gen0 := queue.NewGenerator("gen/0", queue.NewGeneratorContent("gen/0", 3))
	gen1 := queue.NewGenerator("gen/1", queue.NewGeneratorContent("gen/1", 10))
	srv1 := queue.NewServer("que/1", queue.NewServerContent("que/1", 5))
	pipe.AddQueue(nil, gen0, 0)
	pipe.AddQueue(nil, gen1, 0)
	pipe.AddQueue(gen0, srv1, 10)
	pipe.AddQueue(gen1, srv1, 10)
	//tools.ToDisplay("%s", pipe.ToMermaid())

	eng := engine.NewEngine()

	srv1Event := func(server *queue.Server, jobs *[]*queue.Job) *engine.Event {
		return server.ServerEvent(jobs, func(job *queue.Job) {
			t.Errorf("Queue:Serve: job should be completed: %#v\n", job)
		})
	}

	que0ToSrv1Event := func(que *queue.Queue, job *queue.Job) *engine.Event {
		return que.QueueEvent(job, func(jobs *[]*queue.Job) {
			eng.AddEvent(srv1Event(que.GetChild().(*queue.Server), jobs))
		})
	}

	gen0ToQue0Event := gen0.GenEvent(1, func(job *queue.Job) {
		eng.AddEvent(que0ToSrv1Event(gen0.Edges[0].(*queue.Queue), job))
	})

	go func() {
		eng.AddEvent(gen0ToQue0Event)
		time.Sleep(10 * time.Millisecond)
		eng.EndLoop()
	}()
	eng.Loop()
}
