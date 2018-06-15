package graph_test

import (
	"reflect"
	"testing"

	"github.com/jrecuero/go-cli/graph"
)

// TestBlock_NewBlock ensures the block structure works properly
func TestBlock_NewBlock(t *testing.T) {
	b := graph.NewBlock(10, nil, nil, nil)
	if b.Start.BlockID != 10 || b.Start.IsStart == false {
		t.Errorf("new block operation failed: Start")
	} else if b.End.BlockID != 10 || b.End.IsEnd == false {
		t.Errorf("new block operation failed: End")
	} else if b.Loop.BlockID != 10 || b.Loop.IsLoop == false {
		t.Errorf("new block operation failed: Loop")
	} else if b.IsLoop == true {
		t.Errorf("new block operation failed: IsLoop")
	} else if b.IsSkip == true {
		t.Errorf("new block operation failed: IsSkip")
	} else if b.Terminated == true {
		t.Errorf("new block operation failed: Terminated")
	}
}

// TestBlock_CreateBlockNoLoopAndSkip ensures the block structure works properly
func TestBlock_CreateBlockNoLoopAndSkip(t *testing.T) {
	b := graph.NewBlock(1, nil, nil, nil)
	if b.CreateBlockNoLoopAndSkip() == false {
		t.Errorf("create block no loop and skip operation failed")
	} else if b.IsLoop != false || b.IsSkip != true {
		t.Errorf("create block no loop and skip operation failed: Is")
	} else if len(b.Start.Children) != 1 {
		t.Errorf("create block no loop and skip operation failed: len(Start)")
	} else if !reflect.DeepEqual(b.Start.Children[0], b.End) {
		t.Errorf("create block no loop and skip operation failed: match Start")
	} else if len(b.Loop.Children) != 1 {
		t.Errorf("create block no loop and skip operation failed: len(Loop)")
	} else if !reflect.DeepEqual(b.Loop.Children[0], b.End) {
		t.Errorf("create block no loop and skip operation failed: match Loop")
	}
}

// TestBlock_CreateBlockLoopAndSkip ensures the block structure works properly
func TestBlock_CreateBlockLoopAndSkip(t *testing.T) {
	b := graph.NewBlock(1, nil, nil, nil)
	if b.CreateBlockLoopAndSkip() == false {
		t.Errorf("create block loop and skip operation failed")
	} else if b.IsLoop != true || b.IsSkip != true {
		t.Errorf("create block loop and skip operation failed: Is")
	} else if len(b.Start.Children) != 1 {
		t.Errorf("create block loop and skip operation failed: len(Start)")
	} else if !reflect.DeepEqual(b.Start.Children[0], b.End) {
		t.Errorf("create block loop and skip operation failed: match Start")
	} else if len(b.Loop.Children) != 2 {
		t.Errorf("create block loop and skip operation failed: len(Loop)")
	} else if !reflect.DeepEqual(b.Loop.Children[0], b.Start) {
		t.Errorf("create block loop and skip operation failed: match Loop:Start")
	} else if !reflect.DeepEqual(b.Loop.Children[1], b.End) {
		t.Errorf("create block loop and skip operation failed: match Loop:End")
	}
}

// TestBlock_CreateBlockNoLoopAndNoSkip ensures the block structure works properly
func TestBlock_CreateBlockNoLoopAndNoSkip(t *testing.T) {
	b := graph.NewBlock(1, nil, nil, nil)
	if b.CreateBlockNoLoopAndNoSkip() == false {
		t.Errorf("create block no loop and no skip operation failed")
	} else if b.IsLoop != false || b.IsSkip != false {
		t.Errorf("create block no loop and no skip operation failed: Is")
	} else if len(b.Start.Children) != 0 {
		t.Errorf("create block no loop and no skip operation failed: len(Start)")
	} else if len(b.Loop.Children) != 1 {
		t.Errorf("create block no loop and no skip operation failed: len(Loop)")
	} else if !reflect.DeepEqual(b.Loop.Children[0], b.End) {
		t.Errorf("create block no loop and no skip operation failed: match Loop:End")
	}
}

// TestBlock_CreateBlockLoopAndNoSkip ensures the block structure works properly
func TestBlock_CreateBlockLoopAndNoSkip(t *testing.T) {
	b := graph.NewBlock(1, nil, nil, nil)
	if b.CreateBlockLoopAndNoSkip() == false {
		t.Errorf("create block loop and no skip operation failed")
	} else if b.IsLoop != true || b.IsSkip != false {
		t.Errorf("create block loop and no skip operation failed: Is")
	} else if len(b.Start.Children) != 0 {
		t.Errorf("create block loop and no skip operation failed: len(Start)")
	} else if len(b.Loop.Children) != 2 {
		t.Errorf("create block loop and no skip operation failed: len(Loop)")
	} else if !reflect.DeepEqual(b.Loop.Children[0], b.Start) {
		t.Errorf("create block loop and no skip operation failed: match Loop:Start")
	} else if !reflect.DeepEqual(b.Loop.Children[1], b.End) {
		t.Errorf("create block loop and no skip operation failed: match Loop:End")
	}
}

// TestBlock_Terminate ensures the block structure works properly
func TestBlock_Terminate(t *testing.T) {
	b := graph.NewBlock(1, nil, nil, nil)
	if b.CreateBlockNoLoopAndSkip() == false {
		t.Errorf("create block no loop and skip operation failed")
	} else if b.Terminate() == false {
		t.Errorf("terminate operation failed")
	} else if b.Terminated != true {
		t.Errorf("terminate operation failed: Terminated")
	} else if len(b.Start.Children) != 1 {
		t.Errorf("terminate operation failed: len(Start)")
	} else if !reflect.DeepEqual(b.Start.Children[0], b.End) {
		t.Errorf("terminate operation failed: match Start:End")
	}

	b = graph.NewBlock(2, nil, nil, nil)
	if b.CreateBlockLoopAndNoSkip() == false {
		t.Errorf("create block no loop and skip operation failed")
	} else if b.Terminate() == false {
		t.Errorf("terminate operation failed")
	} else if b.Terminated != true {
		t.Errorf("terminate operation failed: Terminated")
	} else if len(b.Start.Children) != 0 {
		t.Errorf("terminate operation failed: len(Start)")
	}
}
