package main

// open list impliments a priority queue for potental path steps
type OpenListItem struct {
	node         *ServerNode
	totalLatency int
	history      *PathHistory
	priority     int
	index        int
}

type OpenList []*OpenListItem

func (this OpenList) Len() int {
	return len(this)
}

func (this OpenList) Less(i, j int) bool {
	return this[i].priority < this[j].priority
}

func (this OpenList) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
	this[i].index = i
	this[j].index = j
}

func (this *OpenList) Push(x interface{}) {
	openListItem := x.(*OpenListItem)
	openListItem.index = len(*this)
	*this = append(*this, openListItem)
}
func (this *OpenList) Pop() interface{} {
	old := *this
	n := len(old)
	openListItem := old[n-1]
	openListItem.index = -1 // for safety
	*this = old[0 : n-1]
	return openListItem
}
