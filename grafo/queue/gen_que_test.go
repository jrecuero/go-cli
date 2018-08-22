package queue_test

import (
	"testing"

	"github.com/jrecuero/go-cli/engine"
	"github.com/jrecuero/go-cli/grafo/queue"
	"github.com/jrecuero/go-cli/tools"
)

// TestGenQue_Engine is ...
func TestGenQue_Engine(t *testing.T) {
	pipe := queue.NewSystem("pipe/1")
	gen0 := queue.NewGenerator("gen/0", queue.NewGeneratorContent("gen/0", 3))
	gen1 := queue.NewGenerator("gen/1", queue.NewGeneratorContent("gen/1", 10))
	que1 := queue.NewServer("que/1", queue.NewServerContent("que/1", 5))
	pipe.AddQueue(nil, queue.GeneratorToVertex(gen0), 0)
	pipe.AddQueue(nil, queue.GeneratorToVertex(gen1), 0)
	pipe.AddQueue(queue.GeneratorToVertex(gen0), queue.ServerToVertex(que1), 10)
	pipe.AddQueue(queue.GeneratorToVertex(gen1), queue.ServerToVertex(que1), 10)
	//tools.ToDisplay("%s", pipe.ToMermaid())

	eng := engine.NewEngine()

	que1Event := func(server *queue.Server, jobs *[]*queue.Job) *engine.Event {
		ev := engine.NewEvent("que/1", 0)
		ev.SetCallback(func(params ...interface{}) error {
			//tools.ToDisplay("que1Event callback\n")
			if j, ok := queue.GetServerContent(server).Serve(jobs); !ok {
				t.Errorf("Queue:Serve: job should be completed: %#v\n", j)
			}
			return nil
		}, nil)
		return ev
	}

	queTo0Event := func(gen *queue.Generator, server *queue.Server, job *queue.Job) *engine.Event {
		ev := engine.NewEvent("event/gen-to-que/0", 2)
		ev.SetCallback(func(params ...interface{}) error {
			//tools.ToDisplay("queTo0Event callback\n")
			params[0].(*queue.Queue).Check(job)
			eng.AddEvent(que1Event(server, params[1].(*[]*queue.Job)))
			return nil
		}, gen.Edges[0], &gen.Edges[0].(*queue.Queue).Jobs)
		return ev
	}

	gen0Event := func(gen *queue.Generator) *engine.Event {
		ev := engine.NewEvent("event/gen/0", 1)
		ev.SetCallback(func(params ...interface{}) error {
			//tools.ToDisplay("gen callback\n")
			if job, ok := queue.GetGeneratorContent(gen).Generate(); ok {
				eng.AddEvent(queTo0Event(gen, que1, job.(*queue.Job)))
				return nil
			}
			return tools.ERROR(nil, false, "Generator error")
		}, nil)
		return ev
	}

	eng.AddEvent(gen0Event(gen0))
	if err := eng.Loop(); err != nil {
		t.Errorf("Engine:Loop: return error code: %#v\n", err)
	}

	//if job, ok := queue.GetGeneratorContent(gen0).Generate(); ok {
	//    edge := gen0.Edges[0]
	//    edge.Check(job)
	//    jobs := &gen0.Edges[0].(*queue.Queue).Jobs
	//    if j, ok := queue.GetServerContent(que1).Serve(jobs); !ok {
	//        t.Errorf("Queue:Serve: job should be completed: %#v\n", j)
	//    }
	//}
}
