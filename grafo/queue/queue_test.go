package queue_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo/queue"
)

// TestQueue_Queue is ...
func TestQueue_Queue(t *testing.T) {
	system := queue.NewSystem("system/1")
	if system == nil {
		t.Errorf("Queue:NewSystem: system can not be <nil>")
	}
	root := queue.NewServer("root/1", nil)
	if root == nil {
		t.Errorf("Network:NewServer: server can not be <nil>")

	}
	queue1 := queue.NewServer("queue/1", nil)
	if queue1 == nil {
		t.Errorf("Network:NewServer: server can not be <nil>")

	}
	if err := system.AddQueue(nil, queue.ServerToVertex(root), 0); err != nil {
		t.Errorf("Queue:AddQueue: return error code: %#v\n", err)
	}
	if err := system.AddQueue(queue.ServerToVertex(root), queue.ServerToVertex(queue1), 10); err != nil {
		t.Errorf("Queue:AddQueue: return error code: %#v\n", err)
	}
	//tools.ToDisplay("%s\n", system.ToMermaid())
}

// TestQueue_Check is ...
func TestQueue_Check(t *testing.T) {
	system := queue.NewSystem("system/1")
	root := queue.NewServer("root/1", nil)
	queue1 := queue.NewServer("queue/1", nil)
	system.AddQueue(nil, queue.ServerToVertex(root), 0)
	system.AddQueue(queue.ServerToVertex(root), queue.ServerToVertex(queue1), 10)
	b := root.Edges[0]
	b.Check(queue.NewJob("job/1", 1))
	b.Check(queue.NewJob("job/2", 2))
	if len(root.Edges[0].(*queue.Queue).Jobs) != 2 {
		t.Errorf("Queue:Check: edges len mismatch: exp: %d got: %d\n", 2, len(root.Edges[0].(*queue.Queue).Jobs))
	}
	//tools.ToDisplay("%#v\n", root.Edges[0].(*queue.Queue).Jobs)
	//for _, j := range root.Edges[0].(*queue.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
}

// TestQueue_Serve is ...
func TestQueue_Serve(t *testing.T) {
	system := queue.NewSystem("system/1")
	root := queue.NewServer("root/1", nil)
	queue1 := queue.NewServer("queue/1", queue.NewServerContent("queue/1", 5))
	system.AddQueue(nil, queue.ServerToVertex(root), 0)
	system.AddQueue(queue.ServerToVertex(root), queue.ServerToVertex(queue1), 10)
	b := root.Edges[0]
	b.Check(queue.NewJob("job/1", 10))
	b.Check(queue.NewJob("job/2", 2))
	//for _, j := range root.Edges[0].(*queue.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
	jobs := &root.Edges[0].(*queue.Queue).Jobs
	if job, ok := queue1.Content.(*queue.ServerContent).Serve(jobs); ok {
		t.Errorf("Queue:Serve: job can not completed: %#v\n", job)
	}
	//for _, j := range root.Edges[0].(*queue.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
	if job, ok := queue1.Content.(*queue.ServerContent).Serve(jobs); !ok {
		t.Errorf("Queue:Serve: job should completed: %#v\n", job)
	}
	//for _, j := range root.Edges[0].(*queue.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
}
