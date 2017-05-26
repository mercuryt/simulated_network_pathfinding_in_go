package main

import (
	"time"
)

const PATH_CACHE_DURATION = time.Duration(time.Second * 5)

type PathCache struct {
	path      *ServerNodePath
	expiresAt time.Time
}

func (this *PathCache) isValid() bool {
	return this.expiresAt.After(time.Now())
}

type PathCacheHash map[int]*PathCache

func (this PathCacheHash) getPathTo(node *ServerNode) *ServerNodePath {
	path := this[node.id]
	if path == nil {
		return nil
	}
	if !path.isValid() {
		delete(this, node.id)
		return nil
	}
	return this[node.id].path
}

func (this PathCacheHash) add(path *ServerNodePath, to *ServerNode) {
	this[to.id] = &PathCache{path, time.Now().Add(PATH_CACHE_DURATION)}
}
