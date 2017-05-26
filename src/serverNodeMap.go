// node map manages map status and external interfaces addServerNode and getPath
package main

import (
	"fmt"

	"github.com/kellydunn/golang-geo"
)

const MAX_HOP_RANGE = 7000     // km
const SPEED_OF_LIGHT = 299.792 // km / ms
const CACHEING_ENABLED = false
const LATENCY_PRIORITY_HURISTIC_MODIFIER = 100

type ServerNodeList []*ServerNode

type ServerNodeMap struct {
	id        int
	list      ServerNodeList
	byName    map[string]*ServerNode
	nodeCount int
}

// external
func newServerNodeMap() ServerNodeMap {
	list := make([]*ServerNode, 0)
	byName := make(map[string]*ServerNode)
	return ServerNodeMap{1, list, byName, 0}
}

// external
func (this *ServerNodeMap) addServerNode(id int, name string, latitude, longitude float64) {
	location := geo.NewPoint(latitude, longitude)
	node := ServerNode{id, name, location, make([]*NearByServerNode, 0), make(PathCacheHash, 0)}
	distance := 0.0
	latency := 10
	minLatency := 0
	for i := range this.list {
		otherNode := this.list[i]
		distance = node.realDistance(otherNode)
		if distance < MAX_HOP_RANGE {
			minLatency = getMinLatency(distance)
			node.recordAdjacent(this.list[i], distance, latency, minLatency)
		}
	}
	this.nodeCount++
	this.list = append(this.list, &node)
	this.byName[node.name] = &node
}

// external
func (this *ServerNodeMap) removeServerNode(name string) {
	node := this.byName[name]
	for _, nearByNode := range node.nearByNodes {
		nearByNode.node.nearByNodes.removeNode(node)
	}
	this.list.removeNode(node)
}

// external
func (this *ServerNodeMap) setLatency(nodeAName, nodeBName string, latency int) {
	nodeA := this.byName[nodeAName]
	nodeB := this.byName[nodeBName]
	nodeA.setLatency(nodeB, latency) // assume symetrical latency
	nodeB.setLatency(nodeA, latency)
}

func getMinLatency(distance float64) int {
	return int(distance / SPEED_OF_LIGHT) // round down for min
}

// external
func (this *ServerNodeMap) getPath(fromName, toName string) *ServerNodePath {
	from, to := this.byName[fromName], this.byName[toName]
	cachePath := from.cache.getPathTo(to)
	if cachePath != nil && CACHEING_ENABLED {
		fmt.Println("cache hit")
		return cachePath
	}
	pathFinder := newPathFinder(from, to)
	path := pathFinder.findPath()
	from.cache.add(path, to)
	return path
}

// external
func (this *ServerNodeMap) printNode(name string) {
	node := this.byName[name]
	fmt.Println(node.name, len(node.nearByNodes))
	for i := range node.nearByNodes {
		nearByNode := node.nearByNodes[i]
		fmt.Println("->", nearByNode.node.name, nearByNode.distance, nearByNode.latency)
	}
}

// change to pointer reciever?
func (this ServerNodeList) removeNode(node *ServerNode) {
	i := this.indexOf(node)
	this[i] = this[len(this)-1]
	this[len(this)-1] = nil
	this = this[:len(this)-1]
}

func (this ServerNodeList) indexOf(node *ServerNode) int {
	for i, otherNode := range this {
		if otherNode == node {
			return i
		}
	}
	return -1
}
