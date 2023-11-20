package helpers

import (
	"container/list"
	ds "eleven-puzzle/data_structures"
)

func FrontierIntersect(first, second *list.Element) (snode, tnode ds.Node, successful bool) {
	for current1 := first; current1 != nil; current1 = current1.Next() {
		sourceNode := current1.Value.(ds.Node)
		for current2 := second; current2 != nil; current2 = current2.Next() {
			targetNode := current2.Value.(ds.Node)
			if sourceNode.Puzzle == targetNode.Puzzle {
				snode, tnode, successful := sourceNode, targetNode, true
				return snode, tnode, successful
			}
		}

	}
	return snode, tnode, successful
}
