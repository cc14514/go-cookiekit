package graph

import (
	"github.com/cc14514/go-cookiekit/collections/bag"
	"github.com/cc14514/go-cookiekit/collections/stack"
	"github.com/cc14514/go-cookiekit/collections/queue"
)

// 深度优先 Depth First Search
type DFSearch struct {
	count  int
	s      int      // 起点 s
	marked *bag.Bag // 与 s 连通的顶点集合
	edgeTo []int    // 边的映射,用来寻找路径
}

func (self *DFSearch) PathTo(v int) []int {
	if !self.Marked(v) {
		return nil
	}
	sk := stack.New()
	for x := v; x != self.s; x = self.edgeTo[x] {
		sk.Push(x)
	}
	sk.Push(self.s)
	result := make([]int, 0)
	for !sk.Empty() {
		result = append(result, sk.Pop().(int))
	}
	return result
}

func (self *DFSearch) GenSearch(graph SimpleGraph, s int) Search {
	self.count, self.s, self.marked, self.edgeTo = 0, s, bag.New(), make([]int, graph.E())
	self.load(graph, s)
	return self
}

func (self *DFSearch) load(graph SimpleGraph, s int) {
	self.marked.Insert(s)
	self.count ++
	for _, a := range graph.Adj(s) {
		if self.marked.Count(a) < 1 {
			self.edgeTo[a] = s
			self.load(graph, a)
		}
	}
}

func (self *DFSearch) Marked(v int) bool {
	return self.marked.Count(v) > 0
}

func (self *DFSearch) Count() int {
	return self.count
}

// 广度优先 Breadth First Search
type BFSearch struct {
	DFSearch
}

func (self *BFSearch) GenSearch(graph SimpleGraph, s int) Search {
	self.count, self.s, self.marked, self.edgeTo = 0, s, bag.New(), make([]int, graph.E())
	self.load(graph, s)
	return self
}

func (self *BFSearch) load(graph SimpleGraph, s int) {
	q := queue.New()
	self.marked.Insert(s)
	self.count ++
	q.Push(s)
	for !q.Empty() {
		v := q.Pop()
		for _, a := range graph.Adj(v.(int)) {
			if !self.Marked(a) {
				self.marked.Insert(a)
				self.count ++
				self.edgeTo[a] = v.(int)
				q.Push(a)
			}
		}
	}
}

// 无环图 Cycle : 深度优先, 判断是否包含环
// 前提是没有平行边和自环
type CycleImpl struct {
	marked  *bag.Bag // 与 s 连通的顶点集合
	isCycle bool
}

func NewCycle(graph SimpleGraph) Cycle {
	c := &CycleImpl{bag.New(), false}
	// 因为图未必是全连同图，所以每条边都要深度遍历一次，
	// 因为有 marked 的存在会过滤掉重复的子图
	for k, _ := range graph.GetAdj() {
		if c.marked.Count(k) < 1 {
			c.load(graph, k, k)
		}
	}
	return c
}

//graph 是图对象
//v 要展开的顶点
//u 上次调用此方法的 v
func (self *CycleImpl) load(graph SimpleGraph, v, u int) {
	self.marked.Insert(v)
	for _, a := range graph.Adj(v) {
		if self.marked.Count(a) < 1 {
			self.load(graph, a, v)
		} else if a != u {
			// 顺着邻接表的一个顶点开始深度优先遍历
			// 如果存在一个顶点被非上一个顶点标记过，则一定存在环
			self.isCycle = true
		}
	}
}

func (self *CycleImpl) HasCycle() bool {
	return self.isCycle
}

// TODO 二分图 TowColor
