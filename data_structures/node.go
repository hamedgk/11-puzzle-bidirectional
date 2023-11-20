package data_structures

import (
	"eleven-puzzle/data_structures/puzzle"
	"fmt"
)

type Node struct {
	Parent    *Node
	Direction puzzle.Direction
	Puzzle    puzzle.Puzzle
}

func (node *Node) Expand(queue Queue, explored map[puzzle.PuzzleBuffer]bool) {
	possibleMoves := node.Puzzle.PossibleBlankMoves()
	for _, direction := range possibleMoves {
		copyNode := *node
		copyNode.Direction = direction
		copyNode.Parent = node
		copyNode.Puzzle.MoveBlank(direction)
		if explored[copyNode.Puzzle.Buffer] {
			continue
		}
		queue.Enqueue(copyNode)
		explored[copyNode.Puzzle.Buffer] = true
	}
}

func (node *Node) IsGoal(buffer puzzle.PuzzleBuffer) bool {
	return node.Puzzle.Buffer == buffer
}

func TraceBack(node Node) {
	if node.Parent == nil {
		fmt.Print("init")
		return
	}

	TraceBack(*node.Parent)
	switch node.Direction {
	case puzzle.Up:
		fmt.Print(" -> Up")
	case puzzle.Down:
		fmt.Print(" -> Down")
	case puzzle.Right:
		fmt.Print(" -> Right")
	case puzzle.Left:
		fmt.Print(" -> Left")
	default:
		return
	}
}

func TraceForward(node Node) {
	if node.Parent == nil {
		return
	}

	switch node.Direction {
	case puzzle.Up:
		fmt.Print(" -> Down")
	case puzzle.Down:
		fmt.Print(" -> Up")
	case puzzle.Right:
		fmt.Print(" -> Left")
	case puzzle.Left:
		fmt.Print(" -> Right")
	default:
		return
	}
	TraceForward(*node.Parent)
}
