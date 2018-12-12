package graph

import (
	"github.com/cc14514/go-cookiekit/collections/bag"
)

// 无向图 接口
type SimpleGraph interface {
	V() int             //顶点数
	E() int             //边数
	AddEdge(v, w int)   //添加一条边
	GetAdj() []*bag.Bag //获取邻接表
	Adj(v int) []int    //和 v 相邻的顶点
	String() string     //对象的字符串表示
}

// 有向图 接口
type SimpleDigraph interface {
	SimpleGraph
	Reverse() SimpleDigraph
}

type search interface {
	Marked(v int) bool  // v 和 s 是连通的吗
	Count() int         // 与 s 连通的顶点总数
	PathTo(v int) []int //返回 s 到 v 的路径
}

// 简单的 Search API
type Search interface {
	search
	GenSearch(graph SimpleGraph, s int) Search
}

// 有向图的 Search API
type DirectedSearch interface {
	search
	GenSearch(graph SimpleDigraph, s int) DirectedSearch
}

// 无向图连通分量
type CC interface {
	Connected(v, w int) bool // v 和 w 是连通的吗
	Count() int // 连通分量数
	ID(v int) int // v 所在的连通分量
}

// 判断一个图是否存在环
type Cycle interface {
	HasCycle() bool
	Cycle() []int
}

// 判断一个图是否为二分图
// 无向图G为二分图的充分必要条件是，G至少有两个顶点，且其所有回路的长度均为偶数
type TowColor interface {
	IsBipartite() bool
}

// 有向图顶点排序
type DigOrder interface {
	Per() []int         // 前序
	Post() []int        // 后序
	ReversePost() []int // 逆后序
}

// 有向无环图的拓扑排序
type Topological interface {
	IsDAG() bool
	Order() []int
}
