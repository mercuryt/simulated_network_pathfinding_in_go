// structures for building the node map
package main

import (
	"github.com/kellydunn/golang-geo"
)

type ServerNode struct {
	id          int
	name        string
	location    *geo.Point
	nearByNodes NearByServerNodes
	cache       PathCacheHash
}

type NearByServerNode struct {
	node       *ServerNode
	distance   float64
	latency    int
	minLatency int
}

type NearByServerNodes []*NearByServerNode

func (this *ServerNode) realDistance(other *ServerNode) float64 { // km
	return this.location.GreatCircleDistance(other.location)
}

func (this *ServerNode) recordAdjacent(other *ServerNode, distance float64, latency, minLatency int) {
	this.nearByNodes = append(this.nearByNodes, &NearByServerNode{other, distance, latency, minLatency})
	other.nearByNodes = append(other.nearByNodes, &NearByServerNode{this, distance, latency, minLatency})
}

func (this *ServerNode) setLatency(to *ServerNode, latency int) {
	index := this.nearByNodes.indexOf(to)
	this.nearByNodes[index].latency = latency
}

func (this NearByServerNodes) removeNode(node *ServerNode) {
	i := this.indexOf(node)
	this[i] = this[len(this)-1]
	this[len(this)-1] = nil
	this = this[:len(this)-1]
}

func (this NearByServerNodes) indexOf(node *ServerNode) int {
	for i, nearByNode := range this {
		if nearByNode.node == node {
			return i
		}
	}
	return -1
}
