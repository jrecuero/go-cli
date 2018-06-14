package syntax

import (
	"fmt"
	"strings"

	"github.com/jrecuero/go-cli/graph"
	"github.com/jrecuero/go-cli/parser"
)

// CommandSyntax represents the command syntax.
type CommandSyntax struct {
	Syntax string
	Parsed *parser.Syntax
	Graph  *graph.Graph
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
	return &CommandSyntax{
		Syntax: st,
		Parsed: ps,
		Graph:  graph.NewGraph(),
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

// CreateGraph creates graph using parsed syntax.
func (cs *CommandSyntax) CreateGraph(c *Command) bool {
	commandLabel := cs.Parsed.Command
	cs.Graph.AddNode(graph.NewNode(commandLabel, c))
	var insideBlock bool
	var block graph.BlockType
	for i, tok := range cs.Parsed.Tokens {
		switch tok {
		case parser.IDENT:
			label := cs.Parsed.Arguments[i]
			var newContent IContent
			newContent, _ = c.LookForArgument(label)
			newNode := graph.NewNode(label, newContent)
			// Check if we are in a block, and use AddNodeToBlock in that case.
			if insideBlock == true {
				cs.Graph.AddNodeToBlock(newNode)
			} else {
				cs.Graph.AddNode(newNode)
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
			endTok, index := lookForCloseBracket(cs.Parsed.Tokens, i)
			block = mapTokenToBlock[endTok]
			fmt.Printf("index=%d token=%d block=%d\n", index, endTok, block)
			// Create the graph block, any node while in the block should be
			// added to this block.
			graph.MapBlockToGraphFunc[block](cs.Graph)
			break
		case parser.CLOSEBRACKET:
			if insideBlock == false {
				return false
			}
			insideBlock = false
			cs.Graph.EndLoop()
			break
		case parser.PIPE:
			if insideBlock == false {
				return false
			}
			break
		case parser.ASTERISK:
			if insideBlock == true || block != graph.LOOPandSKIP {
				return false
			}
			block = graph.NOBLOCK
			break
		case parser.PLUS:
			if insideBlock == true || block != graph.LOOPandNOSKIP {
				return false
			}
			block = graph.NOBLOCK
			break
		case parser.QUESTION:
			if insideBlock == true || block != graph.NOLOOPandSKIP {
				return false
			}
			block = graph.NOBLOCK
			break
		case parser.ADMIRATION:
			if insideBlock == true || block != graph.NOLOOPandNOSKIP {
				return false
			}
			block = graph.NOBLOCK
			break
		case parser.AT:
			if insideBlock == true {
				return false
			}
			break
		case parser.OPENMARK:
			if insideBlock == false {
				return false
			}
			break
		case parser.CLOSEMARK:
			if insideBlock == false {
				return false
			}
			break
		}
	}
	cs.Graph.Terminate()
	return true
}
