package syntax

import (
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/parser"
)

// CommandSyntax represents the command syntax.
type CommandSyntax struct {
	Syntax string         // command syntax as a string.
	Parsed *parser.Syntax // command parsed instance.
	Graph  *graph.Graph   // command parsing tree graph instance.
}

// mapTokenToBlock maps the parser token with required graph block to be created.
var mapTokenToBlock = map[parser.Token]graph.BlockType{
	parser.QUESTION:   graph.NOLOOPandSKIP,
	parser.ASTERISK:   graph.LOOPandSKIP,
	parser.PLUS:       graph.LOOPandNOSKIP,
	parser.ADMIRATION: graph.NOLOOPandNOSKIP,
}

// NewCommandSyntax returns a new instance of CommandSyntax.
func NewCommandSyntax(st string) *CommandSyntax {
	ps, _ := parser.NewParser(strings.NewReader(st)).Parse()
	setupG := &graph.SetupGraph{
		RootContent:  NewContentJoint("Root", "Root content", NewCompleterJoint("root")),
		SinkContent:  GetCR(),
		JointContent: NewContentJoint("Joint", "Joint content", NewCompleterJoint("joint")),
		StartContent: NewContentJoint("Start", "Start content", NewCompleterStart()),
		EndContent:   NewContentJoint("End", "End content", NewCompleterEnd()),
		LoopContent:  NewContentJoint("Loop", "Loop content", NewCompleterLoop()),
	}
	return &CommandSyntax{
		Syntax: st,
		Parsed: ps,
		Graph:  graph.NewGraph(setupG),
	}
}

func lookForCloseBracket(toks []parser.Token, index int) (parser.Token, int) {
	for i, tok := range toks {
		if i < index {
			continue
		}
		if tok == parser.CLOSEBRACKET {
			retIndex := i + 1
			return toks[retIndex], retIndex
		}
	}
	return parser.ILLEGAL, -1
}

// addNodeToGraph adds a content node to a graph with proper casting.
func (cs *CommandSyntax) addNodeToGraph(cn *ContentNode) bool {
	return cs.Graph.AddNode(ContentNodeToNode(cn))
}

// addNodeToBlockToGraph adds a content node to a block graph with proper
// casting.
func (cs *CommandSyntax) addNodeToBlockToGraph(cn *ContentNode) bool {
	return cs.Graph.AddNodeToBlock(ContentNodeToNode(cn))
}

// addNodeAndNodeToBlockToGraph adds a key-value pair to a block graph with
// proper casting.
func (cs *CommandSyntax) addNodeAndNodeToBlockToGraph(cnkey *ContentNode, cnval *ContentNode) bool {
	return cs.Graph.AddIdentAndAnyToBlock(ContentNodeToNode(cnkey), ContentNodeToNode(cnval))
}

// addNodeAndNodeToPathBlockToGraph adds a key-value pair to a block graph with
// proper casting.
func (cs *CommandSyntax) addNodeAndNodeToPathBlockToGraph(cnkey *ContentNode, cnval *ContentNode) bool {
	return cs.Graph.AddIdentAndAnyToPathBlock(ContentNodeToNode(cnkey), ContentNodeToNode(cnval))
}

// CreateGraph creates graph using parsed syntax.
func (cs *CommandSyntax) CreateGraph(c *Command) bool {
	if c.HasChildren && cs.Graph.Next == nil {
		cs.Graph.Next = graph.NewNodeNext(NewContentJoint("Next", "Next content", NewCompleterJoint("next")))
	}
	commandLabel := cs.Parsed.Command
	cs.addNodeToGraph(NewContentNode(commandLabel, c))
	var insideBlock bool
	var piped bool
	var inpath bool
	var block graph.BlockType
	for i, tok := range cs.Parsed.Tokens {
		switch tok {
		case parser.IDENT:
			label := cs.Parsed.Arguments[i]
			newContent, _ := c.LookForArgument(label)
			newNode := NewContentNode(label, newContent)
			// Check if we are in a block, and use AddNodeToBlock in that case.
			if insideBlock == true {
				//cs.addNodeToBlockToGraph(newNode)
				keyContent := newContent.CreateKeywordFromSelf()
				keyNode := NewContentNode(keyContent.GetLabel(), keyContent)
				if !inpath {
					// First token in a block should always be a key-pair.
					cs.addNodeAndNodeToPathBlockToGraph(keyNode, newNode)
				} else if piped {
					// Next tokens shoudl check if a pipied has been found, it
					// piped was present, the add a key-pair.
					cs.Graph.TerminatePathToBlock()
					cs.addNodeAndNodeToPathBlockToGraph(keyNode, newNode)
				} else {
					// If pipe as not been found, then add a simple node.
					cs.addNodeToGraph(newNode)
				}
				inpath = true
			} else {
				cs.addNodeToGraph(newNode)
			}
			break
		case parser.OPENBRACKET:
			if insideBlock == true {
				return false
			}
			insideBlock = true
			// Look forward in the parsed syntax in order to identify which
			// kind of block has to be created.
			// Look for the next entry after parser.CLOSEBRACKET.
			endTok, _ := lookForCloseBracket(cs.Parsed.Tokens, i)
			block = mapTokenToBlock[endTok]
			//tools.Tester("index=%d token=%d block=%d\n", index, endTok, block)
			// Create the graph block, any node while in the block should be
			// added to this block.
			graph.MapBlockToGraphFunc[block](cs.Graph)
			piped = false
			break
		case parser.CLOSEBRACKET:
			if insideBlock == false {
				return false
			}
			insideBlock = false
			if piped || inpath {
				cs.Graph.TerminatePathToBlock()
			}
			cs.Graph.EndLoop()
			piped = false
			inpath = false
			break
		case parser.PIPE:
			if insideBlock == false {
				return false
			}
			piped = true
			break
		case parser.ASTERISK:
			if insideBlock == true || block != graph.LOOPandSKIP {
				return false
			}
			block = graph.NOBLOCK
			piped = false
			break
		case parser.PLUS:
			if insideBlock == true || block != graph.LOOPandNOSKIP {
				return false
			}
			block = graph.NOBLOCK
			piped = false
			break
		case parser.QUESTION:
			if insideBlock == true || block != graph.NOLOOPandSKIP {
				return false
			}
			block = graph.NOBLOCK
			piped = false
			break
		case parser.ADMIRATION:
			if insideBlock == true || block != graph.NOLOOPandNOSKIP {
				return false
			}
			block = graph.NOBLOCK
			piped = false
			break
		case parser.AT:
			if insideBlock == true {
				return false
			}
			piped = false
			break
		case parser.OPENMARK:
			if insideBlock == false {
				return false
			}
			piped = false
			break
		case parser.CLOSEMARK:
			if insideBlock == false {
				return false
			}
			piped = false
			break
		}
	}
	cs.Graph.Terminate()
	return true
}
