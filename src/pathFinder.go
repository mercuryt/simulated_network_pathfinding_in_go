package main

import (
	"container/heap"
)

type PathFinder struct {
	destination *ServerNode
	openList    OpenList
	closedList  map[int]bool
}

func newPathFinder(origin, destination *ServerNode) *PathFinder {
	openList := make(OpenList, 1)
	openList[0] = &OpenListItem{origin, 0, nil, 0, 0}
	heap.Init(&openList)
	return &PathFinder{destination, openList, make(map[int]bool)}
}

func (this *PathFinder) findPath() *ServerNodePath {
	openListItem := heap.Pop(&this.openList).(*OpenListItem)
	if openListItem.node == this.destination {
		history := &PathHistory{this.destination, openListItem.history}
		return history.toPath()
	}
	var totalLatency int
	var nearByNode *NearByServerNode
	for i := range openListItem.node.nearByNodes {
		nearByNode = openListItem.node.nearByNodes[i]
		if !this.closedList[nearByNode.node.id] {
			this.closedList[nearByNode.node.id] = true
			totalLatency = nearByNode.latency + openListItem.totalLatency
			heap.Push(&this.openList, &OpenListItem{
				node:         nearByNode.node,
				totalLatency: totalLatency,
				history:      &PathHistory{openListItem.node, openListItem.history},
				priority:     (totalLatency * LATENCY_PRIORITY_HURISTIC_MODIFIER) + int(this.destination.realDistance(nearByNode.node)),
				index:        -1,
			})
		}
	}
	return this.findPath() // tail call just for fun
}
