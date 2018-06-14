package graph

// BlockType represents block type can be present in a graph.
type BlockType int

const (
	// ILLEGAL block.
	ILLEGAL BlockType = iota

	// NOBLOCK block
	NOBLOCK
	// LOOPandSKIP block.
	LOOPandSKIP
	// LOOPandNOSKIP block.
	LOOPandNOSKIP
	// NOLOOPandSKIP block.
	NOLOOPandSKIP
	// NOLOOPandNOSKIP block
	NOLOOPandNOSKIP
)

// Block represents a graph block.
type Block struct {
	Start      *Node
	End        *Node
	Loop       *Node
	IsSkip     bool
	IsLoop     bool
	Terminated bool
}

// CreateBlockNoLoopAndSkip creates a graph block without a loop
// but it can be skipped.
func (b *Block) CreateBlockNoLoopAndSkip() bool {
	b.IsLoop = false
	b.IsSkip = true
	// Next statement is required for loops that can be skipped.
	b.Start.AddChild(b.End)
	b.Loop.AddChild(b.End)
	return true
}

// CreateBlockNoLoopAndNoSkip creates a graph block without a loop
// and it can not be skipped.
func (b *Block) CreateBlockNoLoopAndNoSkip() bool {
	b.IsLoop = false
	b.IsSkip = false
	b.Loop.AddChild(b.End)
	return true
}

// CreateBlockLoopAndSkip creates a graph block with a loop
// and it can be skipped.
func (b *Block) CreateBlockLoopAndSkip() bool {
	b.IsLoop = true
	b.IsSkip = true
	// Next statement required for loops that can be skipped.
	b.Start.AddChild(b.End)
	// Next statement required for repeated loops.
	b.Loop.AddChild(b.Start)
	b.Loop.AddChild(b.End)
	return true
}

// CreateBlockLoopAndNoSkip creates a graph block with a loop
// and it can not be skipped.
func (b *Block) CreateBlockLoopAndNoSkip() bool {
	b.IsLoop = true
	b.IsSkip = false
	// Next statement required for repeated loops.
	b.Loop.AddChild(b.Start)
	b.Loop.AddChild(b.End)
	return true
}

// Terminate ends a graph loop.
func (b *Block) Terminate() bool {
	// Blocks with skip option are adding a child to END first, but it is
	// better to place that child at the end of the array.
	if b.IsSkip == true {
		childrenLen := len(b.Start.Children)
		b.Start.Children = append(b.Start.Children[1:childrenLen], b.Start.Children[0])
	}
	b.Terminated = true
	return true
}

// NewBlock creates a new graph block.
func NewBlock(id int) *Block {
	b := &Block{
		Start:      NewNodeStart(id),
		End:        NewNodeEnd(id),
		Loop:       NewNodeLoop(id),
		IsLoop:     false,
		IsSkip:     false,
		Terminated: false,
	}
	return b
}
