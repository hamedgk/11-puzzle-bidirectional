package main

import (
	ds "eleven-puzzle/data_structures"
	"eleven-puzzle/data_structures/puzzle"
	"eleven-puzzle/helpers"
	"fmt"
)

var examplePuzzle = puzzle.PuzzleBuffer{
	{6, 4, 255, 8},
	{5, 11, 1, 2},
	{7, 9, 3, 10},
	//{1, 1, 1, 1},
	//{1, 0, 1, 1},
	//{255, 1, 1, 1},
}

func main() {
	sortedArray := puzzle.SortPuzzle(examplePuzzle)

	sourceExplored := map[puzzle.PuzzleBuffer]bool{}
	sourceFrontier := ds.NewQueue()
	sourceFrontier.Enqueue(
		ds.Node{
			Parent:    nil,
			Direction: puzzle.None,
			Puzzle:    puzzle.FromBuffer(examplePuzzle),
		},
	)
	targetExplored := map[puzzle.PuzzleBuffer]bool{}
	targetFrontier := ds.NewQueue()
	targetFrontier.Enqueue(
		ds.Node{
			Parent:    nil,
			Direction: puzzle.None,
			Puzzle:    puzzle.FromBuffer(sortedArray),
		},
	)

	for {
		lastSourceLayer := ds.NewQueue()

		//copy the last layer
		for current := sourceFrontier.Front(); current != nil; current = current.Next() {
			lastSourceLayer.Enqueue(current.Value.(ds.Node))
		}

		if sourceFrontier.IsEmpty() || targetFrontier.IsEmpty() {
			fmt.Println("empty frontier")
			return
		}

		// a == b
		if snode, tnode, success := helpers.FrontierIntersect(sourceFrontier.Front(), targetFrontier.Front()); success {
			ds.TraceBack(snode)
			ds.TraceForward(tnode)
			return
		}

		for i := 0; i < sourceFrontier.Len(); i++ {
			if node, ok := sourceFrontier.Dequeue(); ok {
				node.Expand(sourceFrontier, sourceExplored)
			}
		}

		//a-> == b
		if snode, tnode, success := helpers.FrontierIntersect(sourceFrontier.Front(), targetFrontier.Front()); success {
			ds.TraceBack(snode)
			ds.TraceForward(tnode)
			return
		}

		for i := 0; i < targetFrontier.Len(); i++ {
			if node, ok := targetFrontier.Dequeue(); ok {
				node.Expand(targetFrontier, targetExplored)
			}
		}

		//a == <-b
		if snode, tnode, success := helpers.FrontierIntersect(lastSourceLayer.Front(), targetFrontier.Front()); success {
			ds.TraceBack(snode)
			ds.TraceForward(tnode)
			return
		}

		//a-> == <-b
		if snode, tnode, success := helpers.FrontierIntersect(sourceFrontier.Front(), targetFrontier.Front()); success {
			ds.TraceBack(snode)
			ds.TraceForward(tnode)
			return
		}
	}
}
