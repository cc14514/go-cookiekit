package graph

import (
	"github.com/cc14514/go-cookiekit/collections/bag"
	"github.com/cc14514/go-cookiekit/collections/queue"
	"github.com/cc14514/go-cookiekit/collections/stack"
	"fmt"
)

// =======================
// 无向图 API
// =======================

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
	self.count, self.s, self.marked, self.edgeTo = 0, s, bag.New(), make([]int, graph.V())
	self.dfs(graph, s)
	return self
}

func (self *DFSearch) dfs(graph SimpleGraph, v int) {
	self.marked.Insert(v)
	self.count ++
	for _, w := range graph.Adj(v) {
		if self.marked.Count(w) < 1 {
			self.edgeTo[w] = v
			self.dfs(graph, w)
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
	self.count, self.s, self.marked, self.edgeTo = 0, s, bag.New(), make([]int, graph.V())
	self.bfs(graph, s)
	return self
}

func (self *BFSearch) bfs(graph SimpleGraph, s int) {
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
			c.dfs(graph, k, k)
		}
	}
	return c
}

//graph 是图对象
//v 要展开的顶点
//u 上次调用此方法的 v
func (self *CycleImpl) dfs(graph SimpleGraph, v, u int) {
	self.marked.Insert(v)
	for _, a := range graph.Adj(v) {
		if self.marked.Count(a) < 1 {
			self.dfs(graph, a, v)
		} else if a != u {
			// 顺着邻接表的一个顶点开始深度优先遍历
			// 如果存在一个顶点被标记过，但并非上一个顶点，则一定存在环
			self.isCycle = true
		}
	}
}

func (self *CycleImpl) HasCycle() bool {
	return self.isCycle
}

// 二分图 TowColor
// 无向图G为二分图的充分必要条件是，G至少有两个顶点，且其所有回路的长度均为偶数
type TowColorImpl struct {
	marked      *bag.Bag // 与 s 连通的顶点集合
	color       []bool
	isBipartite bool
}

func NewTowColor(graph SimpleGraph) TowColor {
	tc := new(TowColorImpl)
	tc.isBipartite = true
	tc.marked = bag.New()
	tc.color = make([]bool, graph.V())
	for v, _ := range graph.GetAdj() {
		tc.dfs(graph, v)
	}
	return tc
}

//graph 是图对象
//v 要展开的顶点
func (self *TowColorImpl) dfs(graph SimpleGraph, v int) {
	self.marked.Insert(v)
	for _, a := range graph.Adj(v) {
		if self.marked.Count(a) < 1 {
			// 和 a 相邻的节点必须跟 a 是相反的颜色
			self.color[a] = !self.color[v]
			self.dfs(graph, a)
		} else if self.color[a] == self.color[v] {
			// 顺着邻接表的一个顶点开始深度优先遍历
			// 如果存在一个顶点被标记过，但是跟我颜色相同，则断言一定不是二分图
			self.isBipartite = false
		}
	}
}

func (self *TowColorImpl) IsBipartite() bool {
	return self.isBipartite
}

// =======================
// 有向图 API
// =======================

// 深度优先 Depth First Search
type DirectedSearchDFS struct {
	count int
	// 单点可达性、多点可达性
	ss     []int    // 起点 ss
	marked *bag.Bag // 与 ss 连通的顶点集合
	edgeTo []int    // 边的映射,用来寻找路径
}

func (self *DirectedSearchDFS) Marked(v int) bool {
	b := self.marked.Count(v) > 0
	return b
}

func (self *DirectedSearchDFS) Count() int {
	return self.count
}

func (self *DirectedSearchDFS) PathTo(v int) map[int][]int {
	if !self.Marked(v) {
		return nil
	}
	result := make(map[int][]int)
	for _, s := range self.ss {
		sk := stack.New()
		for x := v; x != s; x = self.edgeTo[x] {
			sk.Push(x)
		}
		sk.Push(s)
		r := make([]int, 0)
		for !sk.Empty() {
			r = append(r, sk.Pop().(int))
		}
		result[s] = r
	}
	return result
}

func (self *DirectedSearchDFS) GenSearch(digraph SimpleDigraph, ss ... int) DirectedSearch {
	self.count, self.ss, self.marked, self.edgeTo = 0, ss, bag.New(), make([]int, digraph.V())
	for _, s := range ss {
		fmt.Println("a ->", s)
		if self.marked.Count(s) < 1 {
			fmt.Println("a <-", s)
			self.dfs(digraph, s)
		}
	}
	return self
}

func (self *DirectedSearchDFS) dfs(digraph SimpleDigraph, s int) {
	self.marked.Insert(s)
	self.count ++
	for _, v := range digraph.Adj(s) {
		if self.marked.Count(v) < 1 {
			//if !self.Marked(v) {
			fmt.Println("b", v, s)
			self.edgeTo[v] = s
			self.dfs(digraph, v)
		}
	}
}
