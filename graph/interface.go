package graph

import (
	"github.com/cc14514/go-cookiekit/collections/bag"
)

// 简单图 接口
type SimpleGraph interface {
	V() int             //顶点数
	E() int             //边数
	AddEdge(v, w int)   //添加一条边
	GetAdj() []*bag.Bag //获取邻接表
	Adj(v int) []int    //和 v 相邻的顶点
	String() string     //对象的字符串表示
}

// 简单的 Search API
type Search interface {
	GenSearch(graph SimpleGraph, s int) Search
	Marked(v int) bool  // v 和 s 是连通的吗
	Count() int         // 与 s 连通的顶点总数
	PathTo(v int) []int //返回 s 到 v 的路径
}

// 判断一个图是否存在环
type Cycle interface {
	HasCycle() bool
}
