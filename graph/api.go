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

type CCImpl struct {
	marked *bag.Bag
	count  int
	id     []int
}

func NewCC(graph SimpleGraph) CC {
	cc := new(CCImpl)
	cc.marked = bag.New()
	cc.count = 0
	cc.id = make([]int, graph.V())
	for v, _ := range graph.GetAdj() {
		if cc.marked.Count(v) < 1 {
			cc.dfs(graph, v)
			cc.count ++
		}
	}
	return cc
}

func (self *CCImpl) dfs(graph SimpleGraph, v int) {
	self.marked.Insert(v)
	self.id[v] = self.count
	for _, w := range graph.Adj(v) {
		if self.marked.Count(w) < 1 {
			self.dfs(graph, w)
		}
	}
}

func (self *CCImpl) Connected(v, w int) bool {
	return self.id[v] == self.id[w]
}

func (self *CCImpl) Count() int {
	return self.count
}

func (self *CCImpl) ID(v int) int {
	return self.id[v]
}

// 无向图 Cycle : 深度优先, 判断是否包含环
// 前提是没有平行边和自环
type CycleImpl struct {
	marked  *bag.Bag // 与 s 连通的顶点集合
	isCycle bool
	edgeTo  []int
	cycle   []int
}

func NewCycle(graph SimpleGraph) Cycle {
	c := &CycleImpl{bag.New(), false, make([]int, graph.V()), nil}
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
		if self.HasCycle() {
			return
		}
		if self.marked.Count(a) < 1 {
			self.edgeTo[a] = v // 记录遍历路径 a <- v
			self.dfs(graph, a, v)
		} else if a != u {
			// 顺着邻接表的一个顶点开始深度优先遍历
			// 如果存在一个顶点被标记过，但并非上一个顶点，则一定存在环
			self.isCycle = true
			self.cycle = make([]int, 0)
			for x := v; x != a; x = self.edgeTo[x] {
				self.cycle = append(self.cycle, x)
			}
			self.cycle = append(self.cycle, a)
			self.cycle = append(self.cycle, v)
		}
	}
}
func (self *CycleImpl) Cycle() []int {
	if !self.HasCycle() {
		return nil
	}
	return self.cycle
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
	s      int      // 起点 ss
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

func (self *DirectedSearchDFS) PathTo(v int) []int {
	if !self.Marked(v) {
		return nil
	}
	sk := stack.New()
	for x := v; x != self.s; x = self.edgeTo[x] {
		sk.Push(x)
	}
	sk.Push(self.s)
	r := make([]int, 0)
	for !sk.Empty() {
		r = append(r, sk.Pop().(int))
	}
	return r
}

func (self *DirectedSearchDFS) GenSearch(digraph SimpleDigraph, s int) DirectedSearch {
	self.count, self.s, self.marked, self.edgeTo = 0, s, bag.New(), make([]int, digraph.V())
	self.dfs(digraph, s)
	return self
}

func (self *DirectedSearchDFS) dfs(digraph SimpleDigraph, s int) {
	self.marked.Insert(s)
	self.count ++
	for _, v := range digraph.Adj(s) {
		if !self.Marked(v) {
			self.edgeTo[v] = s
			self.dfs(digraph, v)
		}
	}
}

// 广度优先 Breadth First Search
type DirectedSearchBFS struct {
	DirectedSearchDFS
}

func (self *DirectedSearchBFS) GenSearch(graph SimpleDigraph, s int) DirectedSearch {
	self.count, self.s, self.marked, self.edgeTo = 0, s, bag.New(), make([]int, graph.V())
	self.bfs(graph, s)
	return self
}

func (self *DirectedSearchBFS) bfs(graph SimpleDigraph, s int) {
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

// 有向图 Cycle : 深度优先, 判断是否包含环
// 前提是没有平行边和自环
type DirectedCycleImpl struct {
	marked  *bag.Bag
	isCycle bool
	edgeTo  []int
	cycle   []int
	onStack []bool
}

func NewDirectedCycle(digraph SimpleDigraph) Cycle {
	dc := new(DirectedCycleImpl)
	dc.edgeTo = make([]int, digraph.V())
	dc.onStack = make([]bool, digraph.V())
	dc.marked = bag.New()
	for v, _ := range digraph.GetAdj() {
		if dc.marked.Count(v) < 1 {
			dc.dfs(digraph, v)
		}
	}
	return dc
}

func (self *DirectedCycleImpl) dfs(digraph SimpleDigraph, v int) {
	self.onStack[v] = true
	defer func() {
		self.onStack[v] = false
	}()
	self.marked.Insert(v)
	for _, w := range digraph.Adj(v) {
		if self.HasCycle() {
			return
		}
		if self.marked.Count(w) < 1 {
			self.edgeTo[w] = v //记录路径 w <- v
			self.dfs(digraph, w)
		} else if self.onStack[w] {
			//如果当前节点在递归栈中，并且已经被标记过了，那这就是一个有向环了
			self.isCycle = true
			self.cycle = make([]int, 0)
			for x := v; x != w; x = self.edgeTo[x] {
				self.cycle = append(self.cycle, x)
			}
			self.cycle = append(self.cycle, w)
			self.cycle = append(self.cycle, v)
		}
	}
	self.marked.Insert(v)
}

func (self *DirectedCycleImpl) HasCycle() bool {
	return self.isCycle
}

func (self *DirectedCycleImpl) Cycle() []int {
	return self.cycle
}

// 深度优先 有向图 顶点排序
type DFOrder struct {
	per         []int
	post        []int
	reversePost []int
	marked      *bag.Bag
}

func NewDFOrder(dig SimpleDigraph) DigOrder {
	o := new(DFOrder)
	o.per = make([]int, 0)
	o.post = make([]int, 0)
	o.reversePost = make([]int, 0)
	o.marked = bag.New()
	for v, _ := range dig.GetAdj() {
		if o.marked.Count(v) < 1 {
			o.dfs(dig, v)
		}
	}
	return o
}

func (self *DFOrder) dfs(dig SimpleDigraph, v int) {
	self.marked.Insert(v)
	self.per = append(self.per, v)
	defer func() {
		self.post = append(self.post, v)
		self.reversePost = append([]int{v}, self.reversePost[:]...)
	}()
	for _, w := range dig.Adj(v) {
		if self.marked.Count(w) < 1 {
			self.dfs(dig, w)
		}

	}
}

func (self *DFOrder) Per() []int {
	return self.per
}

func (self *DFOrder) Post() []int {
	return self.post
}

func (self *DFOrder) ReversePost() []int {
	return self.reversePost
}

// DAG 拓扑排序
type DigTopological struct {
	order []int
	isDAG bool
}

func NewDigTopological(dig SimpleDigraph) Topological {
	digt := new(DigTopological)
	cycle := NewDirectedCycle(dig)
	digt.isDAG = !cycle.HasCycle()
	if digt.isDAG {
		o := NewDFOrder(dig)
		digt.order = o.ReversePost()
		fmt.Println("p", o.Post())
		fmt.Println("r", o.ReversePost())
	}
	return digt
}

func (self *DigTopological) IsDAG() bool {
	return self.isDAG
}

func (self *DigTopological) Order() []int {
	return self.order
}
