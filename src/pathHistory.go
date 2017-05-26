// structurs for tracking branching path histories and extracting the winner
package main

import (
	"fmt"
)

type PathHistory struct {
	node     *ServerNode
	previous *PathHistory
}

type ServerNodePath []*ServerNode

func (this *PathHistory) toPath() *ServerNodePath {
	output := make(ServerNodePath, 0)
	output = append(output, this.node)
	pathPart := this
	for pathPart.previous != nil {
		pathPart = pathPart.previous
		output = append(output, pathPart.node)
	}
	output.reverse()
	return &output
}

func (this ServerNodePath) reverse() {
	for i, j := 0, len(this)-1; i < j; i, j = i+1, j-1 {
		this[i], this[j] = this[j], this[i]
	}
}

func (this *ServerNodePath) print() {
	for i := range *this {
		fmt.Println((*this)[i].name)
	}
}
