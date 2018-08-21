package grafo_test

import (
	"testing"

	"github.com/jrecuero/go-cli/grafo"
)

// TestQueue_Queue is ...
func TestQueue_Queue(t *testing.T) {
	system := grafo.NewSystem("system/1")
	if system == nil {
		t.Errorf("Queue:NewSystem: system can not be <nil>")
	}
	root := grafo.NewServer("root/1", nil)
	if root == nil {
		t.Errorf("Network:NewServer: server can not be <nil>")

	}
	queue1 := grafo.NewServer("queue/1", nil)
	if queue1 == nil {
		t.Errorf("Network:NewServer: server can not be <nil>")

	}
	if err := system.AddQueue(nil, grafo.ServerToLeaf(root), 0); err != nil {
		t.Errorf("Queue:AddQueue: return error code: %#v\n", err)
	}
	if err := system.AddQueue(grafo.ServerToLeaf(root), grafo.ServerToLeaf(queue1), 10); err != nil {
		t.Errorf("Queue:AddQueue: return error code: %#v\n", err)
	}
	//tools.ToDisplay("%s\n", system.ToMermaid())
}

// TestQueue_Check is ...
func TestQueue_Check(t *testing.T) {
	system := grafo.NewSystem("system/1")
	root := grafo.NewServer("root/1", nil)
	queue1 := grafo.NewServer("queue/1", nil)
	system.AddQueue(nil, grafo.ServerToLeaf(root), 0)
	system.AddQueue(grafo.ServerToLeaf(root), grafo.ServerToLeaf(queue1), 10)
	b := root.Branches[0]
	b.Check(grafo.NewJob("job/1", 1))
	b.Check(grafo.NewJob("job/2", 2))
	if len(root.Branches[0].(*grafo.Queue).Jobs) != 2 {
		t.Errorf("Queue:Check: branches len mismatch: exp: %d got: %d\n", 2, len(root.Branches[0].(*grafo.Queue).Jobs))
	}
	//tools.ToDisplay("%#v\n", root.Branches[0].(*grafo.Queue).Jobs)
	//for _, j := range root.Branches[0].(*grafo.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
}

// TestQueue_Serve is ...
func TestQueue_Serve(t *testing.T) {
	system := grafo.NewSystem("system/1")
	root := grafo.NewServer("root/1", nil)
	queue1 := grafo.NewServer("queue/1", grafo.NewServerContent("queue/1", 5))
	system.AddQueue(nil, grafo.ServerToLeaf(root), 0)
	system.AddQueue(grafo.ServerToLeaf(root), grafo.ServerToLeaf(queue1), 10)
	b := root.Branches[0]
	b.Check(grafo.NewJob("job/1", 10))
	b.Check(grafo.NewJob("job/2", 2))
	//for _, j := range root.Branches[0].(*grafo.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
	jobs := &root.Branches[0].(*grafo.Queue).Jobs
	if job, ok := queue1.Content.(*grafo.ServerContent).Serve(jobs); ok {
		t.Errorf("Queue:Serve: job can not completed: %#v\n", job)
	}
	//for _, j := range root.Branches[0].(*grafo.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
	if job, ok := queue1.Content.(*grafo.ServerContent).Serve(jobs); !ok {
		t.Errorf("Queue:Serve: job should completed: %#v\n", job)
	}
	//for _, j := range root.Branches[0].(*grafo.Queue).Jobs {
	//    tools.ToDisplay("%#v\n", j)
	//}
}
